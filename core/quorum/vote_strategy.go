package quorum

import (
	crand "crypto/rand"
	"encoding/json"
	"math"
	"math/rand"
	"sync"
	"time"

	"math/big"

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
	mux                        *event.TypeMux
	minBlockTime, maxBlockTime int // min and max block creation deadline
	minVoteTime, maxVoteTime   int // min and max block voting deadline
	activeMu                   sync.Mutex
	active                     bool
	voteTimer                  *time.Timer
	deadlineTimer              *time.Timer
	rand                       *rand.Rand
}

// NewRandomDeadelineStrategy returns a block maker strategy that
// generated blocks randomly between the given min and max seconds.
func NewRandomDeadelineStrategy(mux *event.TypeMux, minBlockTime, maxBlockTime, minVoteTime, maxVoteTime uint) *randomDeadlineStrategy {
	if minBlockTime > maxBlockTime {
		minBlockTime, maxBlockTime = maxBlockTime, minBlockTime
	}
	if minBlockTime == 0 {
		glog.Info("Set minimum block deadeline interval to 1 second")
		minBlockTime += 1
	}
	if minBlockTime == maxBlockTime {
		maxBlockTime += 1
	}

	if minVoteTime > maxVoteTime {
		minVoteTime, maxVoteTime = maxVoteTime, minVoteTime
	}
	if minVoteTime == 0 {
		glog.Info("Set minimum block deadeline interval to 1 second")
		minVoteTime += 1
	}
	if minVoteTime == maxVoteTime {
		maxVoteTime += 1
	}

	seed, err := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		glog.Fatalf("Unable to seed (psuedo) random deadline strategy")
	}

	s := &randomDeadlineStrategy{
		mux:          mux,
		minBlockTime: int(minBlockTime),
		maxBlockTime: int(maxBlockTime),
		minVoteTime:  int(minVoteTime),
		maxVoteTime:  int(maxVoteTime),
		active:       true,
		rand:         rand.New(rand.NewSource(seed.Int64())),
	}

	return s
}

func resetTimer(t *time.Timer, duration time.Duration) {
	t.Stop()
	select {
	case <-t.C:
	default:
	}
	t.Reset(duration)
}

// Start generating block create request events.
func (s *randomDeadlineStrategy) Start() error {
	if glog.V(logger.Debug) {
		glog.Infof("Random deadline strategy configured with minBlockTime=%d, maxBlockTime=%d, minVoteTime=%d, maxVoteTime=%d",
			s.minBlockTime, s.maxBlockTime, s.minVoteTime, s.maxVoteTime)
	}

	s.voteTimer = time.NewTimer(time.Duration(s.minBlockTime+rand.Intn(s.maxVoteTime-s.minVoteTime)) * time.Second)
	s.deadlineTimer = time.NewTimer(time.Duration(s.minBlockTime+rand.Intn(s.maxBlockTime-s.minBlockTime)) * time.Second)

	go func() {
		sub := s.mux.Subscribe(core.ChainHeadEvent{})
		for {
			select {
			case <-s.voteTimer.C:
				s.activeMu.Lock()
				if s.active {
					s.mux.Post(Vote{})
				}
				s.activeMu.Unlock()

				resetTimer(s.voteTimer, time.Duration(s.minVoteTime+s.rand.Intn(s.maxVoteTime-s.minVoteTime))*time.Second)
			case <-s.deadlineTimer.C:
				s.activeMu.Lock()
				if s.active {
					s.mux.Post(CreateBlock{})
				}
				s.activeMu.Unlock()
				resetTimer(s.deadlineTimer, time.Duration(s.minBlockTime+s.rand.Intn(s.maxBlockTime-s.minBlockTime))*time.Second)
			case <-sub.Chan():
				resetTimer(s.deadlineTimer, time.Duration(s.minBlockTime+s.rand.Intn(s.maxBlockTime-s.minBlockTime))*time.Second)
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
		"minblocktime": s.minBlockTime,
		"maxblocktime": s.maxBlockTime,
		"minvotetime":  s.minVoteTime,
		"maxvotetime":  s.maxVoteTime,
		"status":       status,
	})
}
