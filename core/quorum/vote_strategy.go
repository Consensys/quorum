package quorum

import (
	"math/rand"
	"sync"
	"time"

	"encoding/json"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

type BlockMakerStatus int

const (
	_                       = iota
	Active BlockMakerStatus = iota
	Paused                  = iota
)

type BlockMakerStrategy interface {
	// Start generating blocks
	Start() error
	// (temporary) stop generating blocks
	Pause() error
	// Resume after a pause
	Resume() error
	// Status returns indication if this implementation
	// is generation CreateBlock events.
	Status() BlockMakerStatus
}

// randomDeadlineStrategy asks the block voter to generate blocks
// after a deadline is passed without importing a new head. This
// deadline is chosen random between 2 limits.
type randomDeadlineStrategy struct {
	mux           *event.TypeMux
	min, max      int // min and max deadline
	activeMu      sync.Mutex
	active        bool
	deadlineTimer *time.Timer
}

// NewRandomDeadelineStrategy returns a block maker strategy that
// generated blocks randomly between the given min and max seconds.
func NewRandomDeadelineStrategy(mux *event.TypeMux, min, max uint) *randomDeadlineStrategy {
	if min > max {
		min, max = max, min
	}
	if min == 0 {
		glog.Info("Set minimum block deadeline interval to 1 second")
		min += 1
	}
	if min == max {
		max += 1
	}
	s := &randomDeadlineStrategy{
		mux:    mux,
		min:    int(min),
		max:    int(max),
		active: true,
	}
	return s
}

func resetBlockMakerTimer(t *time.Timer, min, max int) {
	t.Stop()
	select {
	case <-t.C:
	default:
	}
	t.Reset(time.Duration(min+rand.Intn(max-min)) * time.Second)
}

// Start generating block create request events.
func (s *randomDeadlineStrategy) Start() error {
	if glog.V(logger.Debug) {
		glog.Infof("Random deadline strategy configured with min=%d, max=%d", s.min, s.max)
	}
	s.deadlineTimer = time.NewTimer(time.Duration(s.min+rand.Intn(s.max-s.min)) * time.Second)
	go func() {
		sub := s.mux.Subscribe(core.ChainHeadEvent{})
		for {
			select {
			case <-s.deadlineTimer.C:
				s.activeMu.Lock()
				if s.active {
					s.mux.Post(CreateBlock{})
				}
				s.activeMu.Unlock()
				resetBlockMakerTimer(s.deadlineTimer, s.min, s.max)
			case <-sub.Chan():
				resetBlockMakerTimer(s.deadlineTimer, s.min, s.max)
			}
		}
	}()

	return nil
}

// Pause stops generating block create requests.
// Can be resumed with Resume.
func (s *randomDeadlineStrategy) Pause() error {
	s.activeMu.Lock()
	s.active = false
	s.activeMu.Unlock()
	return nil
}

// Resume if paused.
func (s *randomDeadlineStrategy) Resume() error {
	s.activeMu.Lock()
	s.active = true
	s.activeMu.Unlock()
	return nil
}

// Status returns an indication if this strategy is currently
// generating block create request.
func (s *randomDeadlineStrategy) Status() BlockMakerStatus {
	s.activeMu.Lock()
	defer s.activeMu.Unlock()

	if s.active {
		return Active
	}
	return Paused
}

func (s *randomDeadlineStrategy) MarshalJSON() ([]byte, error) {
	s.activeMu.Lock()
	defer s.activeMu.Unlock()

	status := "active"
	if !s.active {
		status = "paused"
	}

	return json.Marshal(map[string]interface{}{
		"type":         "deadline",
		"minblocktime": s.min,
		"maxblocktime": s.max,
		"status":       status,
	})
}
