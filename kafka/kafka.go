package kafka

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/statediff"
	"github.com/shopify/sarama"
)

// Kafka structure to be used by other modules
// when reading or writing from kafka streams
type Kafka struct {
	eventMux    *event.TypeMux
	events      *event.TypeMuxSubscription
	server      *p2p.Server
	blockchain  *core.BlockChain
	stateDiffBuilder *statediff.StateDiffBuilder
	lastBlock   *types.Block
	Producer    sarama.AsyncProducer
	Client      sarama.Client
	quit        chan struct{}
	chainDb     ethdb.Database
	chainConfig *params.ChainConfig
}

type outputBlock struct {
	Origin              string          `json:"obOrigin"`
	TotalDifficulty     *big.Int        `json:"obTotalDifficulty"     gencodec:"required"`
	BlockData           *types.Header   `json:"obBlockData"           gencodec:"required"`
	ReceiptTransactions []outputTx      `json:"obReceiptTransactions" gencodec:"required"`
	BlockUncles         []*types.Header `json:"obBlockUncles"         gencodec:"required"`

	encoded []byte
	err     error
}

type outputTx struct {
	Origin string             `json:"otOrigin" gencodec:"required"`
	Hash   common.Hash        `json:"otHash"   gencodec:"required"`
	Signer common.Address     `json:"otSigner" gencodec:"required"`
	Tx     *types.Transaction `json:"otBaseTx" gencodec:"required"`
}

func (ob *outputBlock) ensureEncoded() {
	if ob.encoded == nil && ob.err == nil {
		ob.encoded, ob.err = json.Marshal(ob)
	}
}

// Implement Encoder interface for outputBlock
func (ob *outputBlock) Length() int {
	ob.ensureEncoded()
	return len(ob.encoded)
}

// Implement Encoder interface for outputBlock
func (ob *outputBlock) Encode() ([]byte, error) {
	ob.ensureEncoded()
	return ob.encoded, ob.err
}

// New runs the event loop.
func New(emux *event.TypeMux, db ethdb.Database, config *params.ChainConfig, bc *core.BlockChain) (*Kafka, error) {
	if sdBuilder, err := statediff.NewStateDiffBuilder(db); err != nil {
		log.Error("Error while creating StateDiffBuilder in kafka", "err", err)
		return nil, err
	} else {
	  return &Kafka{
	  	eventMux:    emux,
	  	chainDb:     db,
			blockchain: bc,
	  	chainConfig: config,
	  	stateDiffBuilder: sdBuilder,
	  }, nil
	}
}

// Protocols
func (k *Kafka) Protocols() []p2p.Protocol { return nil }

// APIs
func (k *Kafka) APIs() []rpc.API { return nil }

// Start
func (k *Kafka) Start(server *p2p.Server) error {
	k.server = server
	// TODO: Move to config file
	brokerList := []string{"localhost:9092"}

	log.Info("Starting Kafka")
	if client, err := newClient(brokerList); err != nil {
		return err
	} else {
		k.Client = client
	}

	if producer, err := newProducer(k.Client); err != nil {
		return err
	} else {
		k.Producer = producer
		if err := k.initializeTopicWithGenesis(); err != nil {
			return err
		}
		go k.loop()
		return nil
	}
}

// Stop
func (k *Kafka) Stop() error {
	// Close the producer before closing the underlying client
	if err := k.Producer.Close(); err != nil {
		log.Error("Error while closing kafka producer", "err", err)
		return err
	}

	if err := k.Client.Close(); err != nil {
		log.Error("Error while closing kafka consumer", "err", err)
		return err
	}

	log.Info("Kafka stopped")
	return nil
}

func (k *Kafka) loop() {
	k.events = k.eventMux.Subscribe(core.ChainEvent{}, core.TxPreEvent{})
	newChainEvent := make(chan core.ChainEvent, 1024)
	chainEventSub := k.blockchain.SubscribeChainEvent(newChainEvent)
	defer k.events.Unsubscribe()
	defer chainEventSub.Unsubscribe()
	log.Debug("starting kafka loop===============================")

	for {
		select {
		case chainEvent := <- newChainEvent:
				transactions := chainEvent.Block.Transactions()
				var receiptTransactions []outputTx
				signer := types.MakeSigner(k.chainConfig, chainEvent.Block.Number())
				for i := 0; i < len(transactions); i++ {
					if message, err := transactions[i].AsMessage(signer); err != nil {

						// TODO: refactor this as a function to get transactions in the shape
						// we need them
					} else {
						receiptTransactions = append(receiptTransactions, outputTx{
							Origin: "Unknown",
							Hash:   transactions[i].Hash(),
							Signer: message.From(),
							Tx:     transactions[i],
						})
					}

				}
				// block := core.GetBlock(k.chainDb, chainEvent.Hash, chainEvent.Block.Number().Uint64())
				opBlock := &outputBlock{
					Origin:              "Unknown",
					TotalDifficulty:     chainEvent.Block.Difficulty(),
					BlockData:           chainEvent.Block.Header(),
					ReceiptTransactions: receiptTransactions,
					BlockUncles:         chainEvent.Block.Uncles(),
				}
				k.Producer.Input() <- &sarama.ProducerMessage{
					// TODO: move to config file or generate on startup
					Topic: "indexevents",
					Key:   nil,
					Value: opBlock,
				}

				if stateDiff, err := k.stateDiffBuilder.CreateStateDiff(k.lastBlock.Root(), chainEvent.Block.Root(), *chainEvent.Block.Number(), chainEvent.Block.Hash()); err!= nil {
					log.Error("Failed to create StateDiff for blocks", "old Block", k.lastBlock.Number(), "new block", chainEvent.Block.Number(), "err", err)
				} else {
					log.Info("StateDiff is:", "statediff", stateDiff)
					k.Producer.Input() <- &sarama.ProducerMessage{
						Topic: "statediff",
						Key:   nil,
						Value: stateDiff,
					}
						k.lastBlock = chainEvent.Block
					}
			break
		}

	}
}

func (k * Kafka)  initializeTopicWithGenesis () error {
	bHash := core.GetCanonicalHash(k.chainDb,0)
	block := core.GetBlock(k.chainDb, bHash, 0)
	k.lastBlock = block
  opBlock := &outputBlock{
		Origin:              "Unknown",
		TotalDifficulty:     block.Difficulty(),
		BlockData:           block.Header(),
		ReceiptTransactions: nil,
		BlockUncles:         block.Uncles(),
	}
	k.Producer.Input() <- &sarama.ProducerMessage{
		// TODO: move to config file or generate on startup
		Topic: "indexevents",
		Key:   nil,
		Value: opBlock,
	}
	return nil
}

func newClient(brokerList []string) (sarama.Client, error) {
	log.Info("Creating Kafka client")
	// By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionNone   // Don't compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	if client, err := sarama.NewClient(brokerList, config); err != nil {
		log.Error("Failed to initialize Kafka client", "err", err)
		return nil, err
	} else {
		log.Info("Created Kafka client")
		return client, nil
	}
}

func newProducer(client sarama.Client) (sarama.AsyncProducer, error) {
	log.Info("Creating Kafka producer")
	producer, err := sarama.NewAsyncProducerFromClient(client)

	if err != nil {
		log.Error("Failed to initialize Kafka producer", "err", err)
		return nil, err
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Error("Failed to write kafka log entry", "err", err)
		}
	}()
	log.Info("Created Kafka producer")
	return producer, nil
}
