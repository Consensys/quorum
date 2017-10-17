package quorum

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

type Status int

const (
	_             = iota
	Active Status = iota
	Paused        = iota
)

func (s Status) MarshalJSON() ([]byte, error) {
	switch s {
	case Active:
		return []byte(`"active"`), nil
	case Paused:
		return []byte(`"inactive"`), nil
	}
	return nil, fmt.Errorf("Could not determine status")
}

type BlockVoteMakerStrategy interface {
	// Start generating blocks/voting
	Start() error
	// (temporary) stop generating blocks
	PauseBlockMaking() error
	// Resume block maker after a pause
	ResumeBlockMaking() error
	// (temporary) stop generating voting events
	PauseVoting() error
	// Resume voting after a pause
	ResumeVoting() error
	// Status returns indication if this implementation
	// is generation CreateBlock and/or Voting events.
	Status() (Status, Status)
}

// randomDeadlineStrategy asks the block voter to generate blocks
// after a deadline is passed without importing a new head. This
// deadline is chosen at random between 2 limits.
type randomDeadlineStrategy struct {
	mux                        *event.TypeMux
	minBlockTime, maxBlockTime int // min and max block creation deadline
	minVoteTime, maxVoteTime   int // min and max block voting deadline
	activeMu                   sync.Mutex
	blockCreateActive          bool
	votingActive               bool
	voteTimer                  *time.Timer
	deadlineTimer              *time.Timer
	rand                       *rand.Rand
}

// NewRandomDeadelineStrategy returns a block maker strategy that
// generated blocks randomly between the given min and max seconds.
func NewRandomDeadelineStrategy(mux *event.TypeMux, minBlockTime, maxBlockTime, minVoteTime, maxVoteTime uint, activateVoting, activateBlockCreation bool) *randomDeadlineStrategy {
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
		mux:               mux,
		minBlockTime:      int(minBlockTime),
		maxBlockTime:      int(maxBlockTime),
		minVoteTime:       int(minVoteTime),
		maxVoteTime:       int(maxVoteTime),
		blockCreateActive: activateBlockCreation,
		votingActive:      activateVoting,
		rand:              rand.New(rand.NewSource(seed.Int64())),
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

	lastVotedHeight := uint64(0)

	go func() {
		sub := s.mux.Subscribe(core.ChainHeadEvent{})
		for {
			select {
			case <-s.voteTimer.C:
				s.activeMu.Lock()
				if s.votingActive {
					if glog.V(logger.Debug) {
						glog.Infoln("issue vote request event")
					}
					go s.mux.Post(Vote{})
				}
				s.activeMu.Unlock()

				resetTimer(s.voteTimer, time.Duration(s.minVoteTime+s.rand.Intn(s.maxVoteTime-s.minVoteTime))*time.Second)
			case <-s.deadlineTimer.C:
				s.activeMu.Lock()
				if s.blockCreateActive {
					if glog.V(logger.Debug) {
						glog.Infoln("issue create block request event")
					}
					go s.mux.Post(CreateBlock{})
				}
				s.activeMu.Unlock()
				resetTimer(s.deadlineTimer, time.Duration(s.minBlockTime+s.rand.Intn(s.maxBlockTime-s.minBlockTime))*time.Second)
			case e := <-sub.Chan():
				if s.votingActive {
					// don't wait for the timer and vote immediately when a new block is imported
					che := e.Data.(core.ChainHeadEvent)
					if che.Block.NumberU64() > lastVotedHeight {
						lastVotedHeight = che.Block.NumberU64()
						go func() {
							if glog.V(logger.Debug) {
								glog.Infoln("Generate vote event for chain head event")
							}
							// post in different go-routine to prevent a deadlock when a
							// new ChainHeadEvent is posted before the Vote event.
							s.mux.Post(Vote{
								Hash:   che.Block.Hash(),
								Number: new(big.Int).Add(big.NewInt(1), che.Block.Number()),
							})
						}()
					}
				}
				resetTimer(s.voteTimer, time.Duration(s.minVoteTime+s.rand.Intn(s.maxVoteTime-s.minVoteTime))*time.Second)
				resetTimer(s.deadlineTimer, time.Duration(s.minBlockTime+s.rand.Intn(s.maxBlockTime-s.minBlockTime))*time.Second)
			}
		}
	}()

	return nil
}

// Pause stops generating block create requests.
// Can be resumed with Resume.
func (s *randomDeadlineStrategy) PauseBlockMaking() error {
	glog.Infoln("Pause block creation")
	s.activeMu.Lock()
	s.blockCreateActive = false
	s.activeMu.Unlock()
	return nil
}

// Resume if paused.
func (s *randomDeadlineStrategy) ResumeBlockMaking() error {
	glog.Infoln("Resume block creation")
	s.activeMu.Lock()
	s.blockCreateActive = true
	s.activeMu.Unlock()
	return nil
}

func (s *randomDeadlineStrategy) PauseVoting() error {
	glog.Infoln("Pause voting")
	s.activeMu.Lock()
	s.votingActive = false
	s.activeMu.Unlock()
	return nil
}

func (s *randomDeadlineStrategy) ResumeVoting() error {
	glog.Infoln("Resume voting")
	s.activeMu.Lock()
	s.votingActive = true
	s.activeMu.Unlock()
	return nil
}

// Status returns an indication if this strategy is currently
// generating block create request.
func (s *randomDeadlineStrategy) Status() (blockMaking, voting Status) {
	s.activeMu.Lock()
	defer s.activeMu.Unlock()

	blockMaking, voting = Paused, Paused
	if s.blockCreateActive {
		blockMaking = Active
	}
	if s.votingActive {
		voting = Active
	}
	return
}

func (s *randomDeadlineStrategy) MarshalJSON() ([]byte, error) {
	block, vote := s.Status()
	s.activeMu.Lock()
	defer s.activeMu.Unlock()

	return json.Marshal(map[string]interface{}{
		"type":          "deadline",
		"minblocktime":  s.minBlockTime,
		"maxblocktime":  s.maxBlockTime,
		"minvotetime":   s.minVoteTime,
		"maxvotetime":   s.maxVoteTime,
		"blockCreation": block,
		"voting":        vote,
	})
}
