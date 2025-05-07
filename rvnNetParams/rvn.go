package rvnNetParams

import (
	"github.com/btcsuite/btcd/chaincfg"
	"math/big"
	"time"
)

var (
	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)

	// mainPowLimit is the highest proof of work value a Ravencoin block can
	// have for the main network.  It is the value 2^224 - 1.
	mainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)
)

// MainNetParams defines the network parameters for the main Ravencoin network.
var MainNetParams = chaincfg.Params{
	Name:        "mainnet",
	Net:         0x5241564e,
	DefaultPort: "8767",
	DNSSeeds: []chaincfg.DNSSeed{
		{"seed-raven.biractivate.com", false},
		{"seed-raven.ravencoin.com", false},
		{"seed-raven.ravencoin.org", false},
	},

	// Chain parameters
	GenesisBlock:             &genesisBlock,
	GenesisHash:              genesisHash,
	PowLimit:                 mainPowLimit,
	PowLimitBits:             0x1d00ffff,
	BIP0034Height:            227931, // 000000000000024b89b42a942fe0d9fea3bb44ab7bd1b19115dd6a759c0808b8
	BIP0065Height:            388381, // 000000000000000004c2b624ed5d7756c508d90fd0da2c7c679febfa6c4735f0
	BIP0066Height:            363725, // 00000000000000000379eaa19dce8c9b722d46ae6a57c2f1a988119488b50931
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 2100000,
	TargetTimespan:           2016 * 60,       // 1.4 days
	TargetTimePerBlock:       time.Minute * 1, // 10 minutes
	RetargetAdjustmentFactor: 4,               // 25% less, 400% more
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{
		{535721, newHashFromStr("000000000001217f58a594ca742c8635ecaaaf695d1a63f6ab06979f1c159e04")},
		{697376, newHashFromStr("000000000000499bf4ebbe61541b02e4692b33defc7109d8f12d2825d4d2dfa0")},
		{740000, newHashFromStr("00000000000027d11bf1e7a3b57d3c89acc1722f39d6e08f23ac3a07e16e3172")},
		{909251, newHashFromStr("000000000000694c9a363eff06518aa7399f00014ce667b9762f9a4e7a49f485")},
		{1040000, newHashFromStr("000000000000138e2690b06b1ddd8cf158c3a5cf540ee5278debdcdffcf75839")},
		{1186833, newHashFromStr("0000000000000d4840d4de1f7d943542c2aed532bd5d6527274fc0142fa1a410")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1613, // 95% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016, //
	Deployments: [chaincfg.DefinedDeployments]chaincfg.ConsensusDeployment{
		chaincfg.DeploymentTestDummy: {
			BitNumber: 28,
			//StartTime:  1199145601, // January 1, 2008 UTC
			//ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		//chaincfg.DeploymentAssets: {
		//	BitNumber:  6,
		//	StartTime:  1540944000, // Oct 31, 2018
		//	ExpireTime: 1572480000, // Oct 31, 2019
		//},
		//chaincfg.DeploymentMsgRestAssets: {
		//	BitNumber:  7,
		//	StartTime:  1578920400, // UTC: Mon Jan 13 2020 13:00:00
		//	ExpireTime: 1610542800, // UTC: Wed Jan 13 2021 13:00:00
		//},
		//chaincfg.DeploymentTransferScriptSize: {
		//	BitNumber:  8,
		//	StartTime:  1588788000, // UTC: Wed May 06 2020 18:00:00
		//	ExpireTime: 1620324000, // UTC: Thu May 06 2021 18:00:00
		//},
		//chaincfg.DeploymentEnforceValue: {
		//	BitNumber:  9,
		//	StartTime:  1593453600, // UTC: Mon Jun 29 2020 18:00:00
		//	ExpireTime: 1624989600, // UTC: Mon Jun 29 2021 18:00:00
		//},
		//chaincfg.DeploymentCoinbaseAssets: {
		//	BitNumber:  10,
		//	StartTime:  1597341600, // UTC: Thu Aug 13 2020 18:00:00
		//	ExpireTime: 1628877600, // UTC: Fri Aug 13 2021 18:00:00
		//},
	},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "rc", // always bc for main net

	// Address encoding magics
	PubKeyHashAddrID:        0x3c, // starts with R
	ScriptHashAddrID:        0x7a, // starts with 3
	PrivateKeyID:            0x80, // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashAddrID: 0x06, // starts with p2
	WitnessScriptHashAddrID: 0x0A, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 175,
}
