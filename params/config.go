// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	TestnetGenesisHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
	RinkebyGenesisHash = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
	GoerliGenesisHash  = common.HexToHash("0xbf7e331f7f7c1dd2e05159666b3bf8bc7a8a3a9eb1d518969eab529dd9b88c1a")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
	RinkebyGenesisHash: RinkebyTrustedCheckpoint,
	GoerliGenesisHash:  GoerliTrustedCheckpoint,
}

// CheckpointOracles associates each known checkpoint oracles with the genesis hash of
// the chain it belongs to.
var CheckpointOracles = map[common.Hash]*CheckpointOracleConfig{
	MainnetGenesisHash: MainnetCheckpointOracle,
	TestnetGenesisHash: TestnetCheckpointOracle,
	RinkebyGenesisHash: RinkebyCheckpointOracle,
	GoerliGenesisHash:  GoerliCheckpointOracle,
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1),
		HomesteadBlock:      big.NewInt(1150000),
		DAOForkBlock:        big.NewInt(1920000),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2463000),
		EIP150Hash:          common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:         big.NewInt(2675000),
		EIP158Block:         big.NewInt(2675000),
		ByzantiumBlock:      big.NewInt(4370000),
		ConstantinopleBlock: big.NewInt(7280000),
		PetersburgBlock:     big.NewInt(7280000),
		IstanbulBlock:       big.NewInt(9069000),
		Ethash:              new(EthashConfig),
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 270,
		SectionHead:  common.HexToHash("0xb67c33d838a60c282c2fb49b188fbbac1ef8565ffb4a1c4909b0a05885e72e40"),
		CHTRoot:      common.HexToHash("0x781daa4607782300da85d440df3813ba38a1262585231e35e9480726de81dbfc"),
		BloomRoot:    common.HexToHash("0xfd8951fa6d779cbc981df40dc31056ed1a549db529349d7dfae016f9d96cae72"),
	}

	// MainnetCheckpointOracle contains a set of configs for the main network oracle.
	MainnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"),
		Signers: []common.Address{
			common.HexToAddress("0x1b2C260efc720BE89101890E4Db589b44E950527"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(3),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:         big.NewInt(10),
		EIP158Block:         big.NewInt(10),
		ByzantiumBlock:      big.NewInt(1700000),
		ConstantinopleBlock: big.NewInt(4230000),
		PetersburgBlock:     big.NewInt(4939394),
		IstanbulBlock:       big.NewInt(6485846),
		Ethash:              new(EthashConfig),
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 204,
		SectionHead:  common.HexToHash("0xa39168b51c3205456f30ce6a91f3590a43295b15a1c8c2ab86bb8c06b8ad1808"),
		CHTRoot:      common.HexToHash("0x9a3654147b79882bfc4e16fbd3421512aa7e4dfadc6c511923980e0877bdf3b4"),
		BloomRoot:    common.HexToHash("0xe72b979522d94fa45c1331639316da234a9bb85062d64d72e13afe1d3f5c17d5"),
	}

	// TestnetCheckpointOracle contains a set of configs for the Ropsten test network oracle.
	TestnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xEF79475013f154E6A65b54cB2742867791bf0B84"),
		Signers: []common.Address{
			common.HexToAddress("0x32162F3581E88a5f62e8A61892B42C46E2c18f7b"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: big.NewInt(3660663),
		PetersburgBlock:     big.NewInt(4321234),
		IstanbulBlock:       big.NewInt(5435345),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 163,
		SectionHead:  common.HexToHash("0x36e5deaa46f258bece94b05d8e10f1ef68f422fb62ed47a2b6e616aa26e84997"),
		CHTRoot:      common.HexToHash("0x829b9feca1c2cdf5a4cf3efac554889e438ee4df8718c2ce3e02555a02d9e9e5"),
		BloomRoot:    common.HexToHash("0x58c01de24fdae7c082ebbe7665f189d0aa4d90ee10e72086bf56651c63269e54"),
	}

	// RinkebyCheckpointOracle contains a set of configs for the Rinkeby test network oracle.
	RinkebyCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xebe8eFA441B9302A0d7eaECc277c09d20D684540"),
		Signers: []common.Address{
			common.HexToAddress("0xd9c9cd5f6779558b6e0ed4e6acf6b1947e7fa1f3"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
		},
		Threshold: 2,
	}

	// GoerliChainConfig contains the chain parameters to run a node on the Görli test network.
	GoerliChainConfig = &ChainConfig{
		ChainID:             big.NewInt(5),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(1561651),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// GoerliTrustedCheckpoint contains the light client trusted checkpoint for the Görli test network.
	GoerliTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 47,
		SectionHead:  common.HexToHash("0x00c5b54c6c9a73660501fd9273ccdb4c5bbdbe5d7b8b650e28f881ec9d2337f6"),
		CHTRoot:      common.HexToHash("0xef35caa155fd659f57167e7d507de2f8132cbb31f771526481211d8a977d704c"),
		BloomRoot:    common.HexToHash("0xbda330402f66008d52e7adc748da28535b1212a7912a21244acd2ba77ff0ff06"),
	}

	// GoerliCheckpointOracle contains a set of configs for the Goerli test network oracle.
	GoerliCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x18CA0E045F0D772a851BC7e48357Bcaab0a0795D"),
		Signers: []common.Address{
			common.HexToAddress("0x4769bcaD07e3b938B7f43EB7D278Bc7Cb9efFb38"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil, false, 32, 35, big.NewInt(0), big.NewInt(0), nil, nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, &CliqueConfig{Period: 0, Epoch: 30000}, nil, false, 32, 32, big.NewInt(0), big.NewInt(0), nil, nil}

	TestChainConfig = &ChainConfig{big.NewInt(10), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil, false, 32, 32, big.NewInt(0), big.NewInt(0), nil, nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))

	QuorumTestChainConfig = &ChainConfig{big.NewInt(10), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, new(EthashConfig), nil, nil, true, 64, 32, big.NewInt(0), big.NewInt(0), nil, big.NewInt(0)}
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// HashEqual returns an indicator comparing the itself hash with given one.
func (c *TrustedCheckpoint) HashEqual(hash common.Hash) bool {
	if c.Empty() {
		return hash == common.Hash{}
	}
	return c.Hash() == hash
}

// Hash returns the hash of checkpoint's four key fields(index, sectionHead, chtRoot and bloomTrieRoot).
func (c *TrustedCheckpoint) Hash() common.Hash {
	buf := make([]byte, 8+3*common.HashLength)
	binary.BigEndian.PutUint64(buf, c.SectionIndex)
	copy(buf[8:], c.SectionHead.Bytes())
	copy(buf[8+common.HashLength:], c.CHTRoot.Bytes())
	copy(buf[8+2*common.HashLength:], c.BloomRoot.Bytes())
	return crypto.Keccak256Hash(buf)
}

// Empty returns an indicator whether the checkpoint is regarded as empty.
func (c *TrustedCheckpoint) Empty() bool {
	return c.SectionHead == (common.Hash{}) || c.CHTRoot == (common.Hash{}) || c.BloomRoot == (common.Hash{})
}

// CheckpointOracleConfig represents a set of checkpoint contract(which acts as an oracle)
// config which used for light client checkpoint syncing.
type CheckpointOracleConfig struct {
	Address   common.Address   `json:"address"`
	Signers   []common.Address `json:"signers"`
	Threshold uint64           `json:"threshold"`
}

type MaxCodeConfigStruct struct {
	Block *big.Int `json:"block,omitempty"`
	Size  uint64   `json:"size,omitempty"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash   *EthashConfig   `json:"ethash,omitempty"`
	Clique   *CliqueConfig   `json:"clique,omitempty"`
	Istanbul *IstanbulConfig `json:"istanbul,omitempty"`

	IsQuorum             bool   `json:"isQuorum"`     // Quorum flag
	TransactionSizeLimit uint64 `json:"txnSizeLimit"` // Quorum - transaction size limit
	MaxCodeSize          uint64 `json:"maxCodeSize"`  // Quorum -  maximum CodeSize of contract
	// Quorum
	//
	// QIP714Block implements the permissions related changes
	QIP714Block            *big.Int `json:"qip714Block,omitempty"`
	MaxCodeSizeChangeBlock *big.Int `json:"maxCodeSizeChangeBlock,omitempty"`
	// to track multiple changes to maxCodeSize
	MaxCodeSizeConfig []MaxCodeConfigStruct `json:"maxCodeSizeConfig,omitempty"`
	// Quorum

	PrivacyEnhancementsBlock *big.Int `json:"privacyEnhancementsBlock,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period                 uint64 `json:"period"`                 // Number of seconds between blocks to enforce
	Epoch                  uint64 `json:"epoch"`                  // Epoch length to reset votes and checkpoint
	AllowedFutureBlockTime uint64 `json:"allowedFutureBlockTime"` // Max time (in seconds) from current time allowed for blocks, before they're considered future blocks
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// IstanbulConfig is the consensus engine configs for Istanbul based sealing.
type IstanbulConfig struct {
	Epoch          uint64   `json:"epoch"`                    // Epoch length to reset votes and checkpoint
	ProposerPolicy uint64   `json:"policy"`                   // The policy for proposer selection
	Ceil2Nby3Block *big.Int `json:"ceil2Nby3Block,omitempty"` // Number of confirmations required to move from one state to next [2F + 1 to Ceil(2N/3)]
}

// String implements the stringer interface, returning the consensus engine details.
func (c *IstanbulConfig) String() string {
	return "istanbul"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Istanbul != nil:
		engine = c.Istanbul
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v IsQuorum: %v Constantinople: %v TransactionSizeLimit: %v MaxCodeSize: %v Petersburg: %v Istanbul: %v PrivacyEnhancements: %v Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.IsQuorum,
		c.ConstantinopleBlock,
		c.TransactionSizeLimit,
		c.MaxCodeSize,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.PrivacyEnhancementsBlock,
		engine,
	)
}

// Quorum - validate code size and transaction size limit
func (c *ChainConfig) IsValid() error {

	if c.TransactionSizeLimit < 32 || c.TransactionSizeLimit > 128 {
		return errors.New("Genesis transaction size limit must be between 32 and 128")
	}

	if c.MaxCodeSize != 0 && (c.MaxCodeSize < 24 || c.MaxCodeSize > 128) {
		return errors.New("Genesis max code size must be between 24 and 128")
	}

	return nil
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// Quorum
//
// IsQIP714 returns whether num represents a block number where permissions is enabled
func (c *ChainConfig) IsQIP714(num *big.Int) bool {
	return isForked(c.QIP714Block, num)
}

// IsMaxCodeSizeChangeBlock returns whether num represents a block number
// where maxCodeSize change was done
func (c *ChainConfig) IsMaxCodeSizeChangeBlock(num *big.Int) bool {
	return isForked(c.MaxCodeSizeChangeBlock, num)
}

// Quorum
//
// GetMaxCodeSize returns maxCodeSize for the given block number
func (c *ChainConfig) GetMaxCodeSize(num *big.Int) int {
	maxCodeSize := MaxCodeSize

	if len(c.MaxCodeSizeConfig) > 0 {
		for _, data := range c.MaxCodeSizeConfig {
			if data.Block.Cmp(num) > 0 {
				break
			}
			maxCodeSize = int(data.Size) * 1024
		}
	} else if c.MaxCodeSize > 0 {
		if c.MaxCodeSizeChangeBlock != nil && c.MaxCodeSizeChangeBlock.Cmp(big.NewInt(0)) >= 0 {
			if c.IsMaxCodeSizeChangeBlock(num) {
				maxCodeSize = int(c.MaxCodeSize) * 1024
			}
		} else {
			maxCodeSize = int(c.MaxCodeSize) * 1024
		}
	}
	return maxCodeSize
}

// validates the maxCodeSizeConfig data passed in config
func (c *ChainConfig) CheckMaxCodeConfigData() error {
	if c.MaxCodeSize != 0 || (c.MaxCodeSizeChangeBlock != nil && c.MaxCodeSizeChangeBlock.Cmp(big.NewInt(0)) >= 0) {
		return errors.New("maxCodeSize & maxCodeSizeChangeBlock deprecated. Consider using maxCodeSizeConfig")
	}
	// validate max code size data
	// 1. Code size should not be less than 24 and greater than 128
	// 2. block entries are in ascending order
	prevBlock := big.NewInt(0)
	for _, data := range c.MaxCodeSizeConfig {
		if data.Size < 24 || data.Size > 128 {
			return errors.New("Genesis max code size must be between 24 and 128")
		}
		if data.Block == nil {
			return errors.New("Block number not given in maxCodeSizeConfig data")
		}
		if data.Block.Cmp(prevBlock) < 0 {
			return errors.New("invalid maxCodeSize detail, block order has to be ascending")
		}
		prevBlock = data.Block
	}

	return nil
}

// checks if changes to maxCodeSizeConfig proposed are compatible
// with already existing genesis data
func isMaxCodeSizeConfigCompatible(c1, c2 *ChainConfig, head *big.Int) (error, *big.Int, *big.Int) {
	if len(c1.MaxCodeSizeConfig) == 0 && len(c2.MaxCodeSizeConfig) == 0 {
		// maxCodeSizeConfig not used. return
		return nil, big.NewInt(0), big.NewInt(0)
	}

	// existing config had maxCodeSizeConfig and new one does not have the same return error
	if len(c1.MaxCodeSizeConfig) > 0 && len(c2.MaxCodeSizeConfig) == 0 {
		return fmt.Errorf("genesis file missing max code size information"), head, head
	}

	if len(c2.MaxCodeSizeConfig) > 0 && len(c1.MaxCodeSizeConfig) == 0 {
		return nil, big.NewInt(0), big.NewInt(0)
	}

	// check the number of records below current head in both configs
	// if they do not match throw an error
	c1RecsBelowHead := 0
	for _, data := range c1.MaxCodeSizeConfig {
		if data.Block.Cmp(head) <= 0 {
			c1RecsBelowHead++
		} else {
			break
		}
	}

	c2RecsBelowHead := 0
	for _, data := range c2.MaxCodeSizeConfig {
		if data.Block.Cmp(head) <= 0 {
			c2RecsBelowHead++
		} else {
			break
		}
	}

	// if the count of past records is not matching return error
	if c1RecsBelowHead != c2RecsBelowHead {
		return errors.New("maxCodeSizeConfig data incompatible. updating maxCodeSize for past"), head, head
	}

	// validate that each past record is matching exactly. if not return error
	for i := 0; i < c1RecsBelowHead; i++ {
		if c1.MaxCodeSizeConfig[i].Block.Cmp(c2.MaxCodeSizeConfig[i].Block) != 0 ||
			c1.MaxCodeSizeConfig[i].Size != c2.MaxCodeSizeConfig[i].Size {
			return errors.New("maxCodeSizeConfig data incompatible. maxCodeSize historical data does not match"), head, head
		}
	}

	return nil, big.NewInt(0), big.NewInt(0)
}

// IsPrivacyEnhancementsEnabled returns whether num represents a block number after the PrivacyEnhancementsEnabled fork
func (c *ChainConfig) IsPrivacyEnhancementsEnabled(num *big.Int) bool {
	return isForked(c.PrivacyEnhancementsBlock, num)
}

// /Quorum

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64, isQuorumEIP155Activated bool) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// check if the maxCodesize data passed is compatible 1st
	// this is being handled separately as it can have breaks
	// at multiple block heights and cannot be handled with in
	// checkCompatible

	// compare the maxCodeSize data between the old and new config
	err, cBlock, newCfgBlock := isMaxCodeSizeConfigCompatible(c, newcfg, bhead)
	if err != nil {
		return newCompatError(err.Error(), cBlock, newCfgBlock)
	}

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead, isQuorumEIP155Activated)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, geth isn't pluggable enough
// to guarantee that forks
func (c *ChainConfig) CheckConfigForkOrder() error {
	type fork struct {
		name  string
		block *big.Int
	}
	var lastFork fork
	for _, cur := range []fork{
		{"homesteadBlock", c.HomesteadBlock},
		{"eip150Block", c.EIP150Block},
		{"eip155Block", c.EIP155Block},
		{"eip158Block", c.EIP158Block},
		{"byzantiumBlock", c.ByzantiumBlock},
		{"constantinopleBlock", c.ConstantinopleBlock},
		{"petersburgBlock", c.PetersburgBlock},
		{"istanbulBlock", c.IstanbulBlock},
	} {
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		lastFork = cur
	}
	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int, isQuorumEIP155Activated bool) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isQuorumEIP155Activated && c.ChainID != nil && isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isQuorumEIP155Activated && c.ChainID != nil && c.IsEIP155(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP155 chain ID", c.ChainID, newcfg.ChainID)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	if c.Istanbul != nil && newcfg.Istanbul != nil && isForkIncompatible(c.Istanbul.Ceil2Nby3Block, newcfg.Istanbul.Ceil2Nby3Block, head) {
		return newCompatError("Ceil 2N/3 fork block", c.Istanbul.Ceil2Nby3Block, newcfg.Istanbul.Ceil2Nby3Block)
	}
	if isForkIncompatible(c.QIP714Block, newcfg.QIP714Block, head) {
		return newCompatError("permissions fork block", c.QIP714Block, newcfg.QIP714Block)
	}
	if newcfg.MaxCodeSizeChangeBlock != nil && isForkIncompatible(c.MaxCodeSizeChangeBlock, newcfg.MaxCodeSizeChangeBlock, head) {
		return newCompatError("max code size change fork block", c.MaxCodeSizeChangeBlock, newcfg.MaxCodeSizeChangeBlock)
	}
	if isForkIncompatible(c.PrivacyEnhancementsBlock, newcfg.PrivacyEnhancementsBlock, head) {
		return newCompatError("Privacy Enhancements fork block", c.PrivacyEnhancementsBlock, newcfg.PrivacyEnhancementsBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsPrivacyEnhancementsEnabled                            bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:                      new(big.Int).Set(chainID),
		IsHomestead:                  c.IsHomestead(num),
		IsEIP150:                     c.IsEIP150(num),
		IsEIP155:                     c.IsEIP155(num),
		IsEIP158:                     c.IsEIP158(num),
		IsByzantium:                  c.IsByzantium(num),
		IsConstantinople:             c.IsConstantinople(num),
		IsPetersburg:                 c.IsPetersburg(num),
		IsIstanbul:                   c.IsIstanbul(num),
		IsPrivacyEnhancementsEnabled: c.IsPrivacyEnhancementsEnabled(num),
	}
}
