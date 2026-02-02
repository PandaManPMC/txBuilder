package hdWallet

type HDCoinType int

// HDWallet coinType BIP-44
const (
	BTCHDCoinType  HDCoinType = 0
	LTCHDCoinType  HDCoinType = 2
	DOGEHDCoinType HDCoinType = 3
	ETHHDCoinType  HDCoinType = 60
	RVNHDCoinType  HDCoinType = 175
	TRONHDCoinType HDCoinType = 195
	SOLHDCoinType  HDCoinType = 501
)
