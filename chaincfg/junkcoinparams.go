// Copyright (c) 2025 Junkcoin Developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// JunkcoinMainNetParams defines the network parameters for the Junkcoin main network.
var JunkcoinMainNetParams = Params{
	Name:        "junkcoin-mainnet",
	Net:         0x6a756e6b, // "junk" in ASCII as a 4-byte uint32
	DefaultPort: "9771",     // Junkcoin mainnet port
	DNSSeeds: []DNSSeed{
		{"mainnet.junk-coin.com", true},
		{"junk-seed.s3na.xyz", true},
		{"jkc-seed.junkiewally.xyz", true},
	},

	// Chain parameters
	GenesisBlock:             &junkcoinGenesisBlock,
	GenesisHash:              &junkcoinGenesisHash,
	PowLimit:                 junkcoinMainPowLimit,
	PowLimitBits:             0x1e0ffff0, // Junkcoin difficulty
	BIP0034Height:            0,          // Always active on Junkcoin
	BIP0065Height:            0,          // Always active on Junkcoin
	BIP0066Height:            0,          // Always active on Junkcoin
	CoinbaseMaturity:         70,         // 70 blocks maturity
	SubsidyReductionInterval: 518400,     // Halving every 2 years (518400 blocks)
	TargetTimespan:           time.Hour * 24, // 24 hours
	TargetTimePerBlock:       time.Minute * 1, // 1 minute blocks
	RetargetAdjustmentFactor: 4,               // 25% less, 400% more
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{0, newHashFromStr("a2effa738145e377e08a61d76179c21703e13e48910b30a2a87f0dfe794b64c6")},
		{1, newHashFromStr("ca55073a54775a1ef78294f53f38a3e02d0654d7417f3cbbe4d28d17d50e07d0")},
		{53, newHashFromStr("b623a39a5a0534990a59916d5803fa2bd6a6d52d8e594546936a42a2cc9b0441")},
		{117, newHashFromStr("6cab49bd69fcce2bb48793cc064bb49e75f068e7029b5173db83654fbcb5953d")},
		{200, newHashFromStr("45257b0f2ee6d5c55ac16a76817d7151b776d6452ae6f21426eaa42345b831f8")},
		{6452, newHashFromStr("506562c2172d9f10e86d2b467ed3bb7b9eba40148d18d1e660c1ff692604f3fc")},
		{10978, newHashFromStr("1c9f7f7a4702f8225df430b259ac58c387de99439be8a8789841a1c011ead7fc")},
		{17954, newHashFromStr("6036051659e92a17cb7488040e05a94483b7a7f88b184156c136d51ff0390a7d")},
		{23978, newHashFromStr("7924154aa896363ec9be3ca5f939602f72cf4a5396e6e1cd9139335dd1819487")},
		{33212, newHashFromStr("448040ac454da8654d9c58ad79386aa1a88fd113be0fcc5ca39ecd3eae8c8618")},
		{45527, newHashFromStr("f2420d964001d4d2c8bc0d9283f3f684d4d91a509a50985888458a68e08e1c82")},
		{57484, newHashFromStr("c3e95c6fb35f4b39006c89538415b4f50a253a3ac1cad0e583fb287f6bd91be1")},
		{69240, newHashFromStr("c34f5d113fe92f3206ef8855caf51cd6252286e3381b253bbc1237211198c22b")},
		{73892, newHashFromStr("d05129c2d9f3e99565bf84fbceabbc61728e4d644173e194823b639f7c406b04")},
		{168312, newHashFromStr("deea2bcecb1146ae9cd74d67b29b4d0161e9bb63beb9022ca10f3625dda6c0e6")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 9576, // 95% of MinerConfirmationWindow
	MinerConfirmationWindow:       10080, // 24 hours worth of blocks (1440 * 7)

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "jc", // always jc for mainnet

	// Address encoding magics (Junkcoin uses different prefixes)
	PubKeyHashAddrID:        0x10, // starts with 7 (legacy addresses)
	ScriptHashAddrID:        0x05, // starts with 3
	PrivateKeyID:            0x90, // starts with N (WIF)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 2013, // BIP44 coin type for Junkcoin
}

// JunkcoinTestNetParams defines the network parameters for the Junkcoin test network.
var JunkcoinTestNetParams = Params{
	Name:        "junkcoin-testnet",
	Net:         0x6a756e6b + 1, // "junk" + 1 in ASCII as a 4-byte uint32
	DefaultPort: "19771",        // Junkcoin testnet port
	DNSSeeds: []DNSSeed{
		{"testnet.junk-coin.com", true},
		{"junk-testnet.s3na.xyz", true},
	},

	// Chain parameters
	GenesisBlock:             &junkcoinTestNetGenesisBlock,
	GenesisHash:              &junkcoinTestNetGenesisHash,
	PowLimit:                 junkcoinTestNetPowLimit,
	PowLimitBits:             0x1e0ffff0, // Junkcoin testnet difficulty
	BIP0034Height:            0,          // Always active on Junkcoin testnet
	BIP0065Height:            0,          // Always active on Junkcoin testnet
	BIP0066Height:            0,          // Always active on Junkcoin testnet
	CoinbaseMaturity:         30,         // 30 blocks maturity for testnet
	SubsidyReductionInterval: 518400,     // Halving every 2 years (518400 blocks)
	TargetTimespan:           time.Hour * 4, // 4 hours for testnet
	TargetTimePerBlock:       time.Minute * 1, // 1 minute blocks
	RetargetAdjustmentFactor: 4,               // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 2, // TargetTimePerBlock * 2
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{0, newHashFromStr("a2effa738145e377e08a61d76179c21703e13e48910b30a2a87f0dfe794b64c6")},
	},

	// Consensus rule change deployments.
	RuleChangeActivationThreshold: 1512, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016, // 4 hours worth of blocks (240 * 8.4)

	// Mempool parameters
	RelayNonStdTxs: true,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tj", // always tj for testnet

	// Address encoding magics (Junkcoin testnet uses different prefixes)
	PubKeyHashAddrID:        0x6f, // starts with m or n
	ScriptHashAddrID:        0xc4, // starts with 2
	PrivateKeyID:            0xef, // starts with 9 (uncompressed) or c (compressed)
	WitnessPubKeyHashAddrID: 0x03, // starts with QW
	WitnessScriptHashAddrID: 0x28, // starts with T7n

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 11337, // Custom coin type for Junkcoin testnet
}

// junkcoinMainPowLimit is the highest proof of work value a Junkcoin block can
// have for the main network. It is the value 2^224 - 1.
var junkcoinMainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)

// junkcoinTestNetPowLimit is the highest proof of work value a Junkcoin block
// can have for the test network. It is the value 2^224 - 1.
var junkcoinTestNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)

// junkcoinGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the main network.
var junkcoinGenesisMerkleRoot = mustParseHash("4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b")

// junkcoinGenesisBlock defines the genesis block of the block chain which serves as
// the public transaction ledger for the main network.
var junkcoinGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{}, // All zeroes
		MerkleRoot: junkcoinGenesisMerkleRoot,
		Timestamp:  time.Unix(1231006505, 0), // TODO: Update with actual timestamp
		Bits:       0x1d00ffff,
		Nonce:      2083236893, // TODO: Update with actual nonce
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// junkcoinTestNetGenesisMerkleRoot is the hash of the first transaction in the
// genesis block for the test network.
var junkcoinTestNetGenesisMerkleRoot = mustParseHash("4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b")

// junkcoinTestNetGenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network.
var junkcoinTestNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{}, // All zeroes
		MerkleRoot: junkcoinTestNetGenesisMerkleRoot,
		Timestamp:  time.Unix(1296688602, 0), // TODO: Update with actual timestamp
		Bits:       0x1d00ffff,
		Nonce:      414098458, // TODO: Update with actual nonce
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// mustParseHash converts the passed big-endian hex string into a
// chainhash.Hash and will panic if there is an error. It only differs from the
// one available in chainhash in the fact that it panics on an error so it is
// only safe to use with hard coded values.
func mustParseHash(s string) chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(s)
	if err != nil {
		panic(err)
	}
	return *hash
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network.
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec| */
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for | */
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.|
			},
		},
	},
	LockTime: 0,
}

// Initialize Junkcoin genesis hashes
var (
	junkcoinGenesisHash     = junkcoinGenesisBlock.BlockHash()
	junkcoinTestNetGenesisHash = junkcoinTestNetGenesisBlock.BlockHash()
)

func init() {
	// Register Junkcoin networks
	err := Register(&JunkcoinMainNetParams)
	if err != nil {
		panic(err)
	}

	err = Register(&JunkcoinTestNetParams)
	if err != nil {
		panic(err)
	}
}
