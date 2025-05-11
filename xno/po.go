package xno

type WorkGenerateRsp struct {
	Difficulty string `json:"difficulty"`
	Multiplier string `json:"multiplier"`
	Work       string `json:"work"`
}

type ProcessRsp struct {
	Hash string `json:"hash"`
}

// ActiveDifficultyRsp
// {"deprecated":"1","network_minimum":"fffffff800000000",
// "network_receive_minimum":"fffffe0000000000","network_current":"fffffff800000000",
// "network_receive_current":"fffffe0000000000","multiplier":"1"}
type ActiveDifficultyRsp struct {
	Deprecated            string `json:"deprecated"`
	NetworkMinimum        string `json:"network_minimum"`
	NetworkReceiveMinimum string `json:"network_receive_minimum"`
	NetworkCurrent        string `json:"network_current"`
	NetworkReceiveCurrent string `json:"network_receive_current"`
	Multiplier            string `json:"multiplier"`
}

type AccountInfoRsp struct {
	Frontier                   string `json:"frontier"`
	OpenBlock                  string `json:"open_block"`
	RepresentativeBlock        string `json:"representative_block"`
	Balance                    string `json:"balance"`
	ModifiedTimestamp          string `json:"modified_timestamp"`
	BlockCount                 string `json:"block_count"`
	AccountVersion             string `json:"account_version"`
	ConfirmationHeight         string `json:"confirmation_height"`
	ConfirmationHeightFrontier string `json:"confirmation_height_frontier"`
}

type AccountBalanceRsp struct {
	Balance    string `json:"balance"`
	Pending    string `json:"pending"`
	Receivable string `json:"receivable"`
}

type AccountsPendingRsp struct {
	Deprecated        string              `json:"deprecated"`
	Blocks            map[string][]string `json:"blocks"`
	RequestsLimit     string              `json:"requestsLimit"`
	RequestsRemaining string              `json:"requestsRemaining"`
	RequestLimitReset string              `json:"requestLimitReset"`
}

type BlockInfoRsp struct {
	BlockAccount   string `json:"block_account"`   // 该区块所属账户地址
	Amount         string `json:"amount"`          // 本区块变动的金额（单位为 raw）
	Balance        string `json:"balance"`         // 本区块后的账户余额（单位为 raw）
	Height         string `json:"height"`          // 区块在账户链中的高度（从 1 开始）
	LocalTimestamp string `json:"local_timestamp"` // 区块生成时的本地时间戳（可能为 "0"）
	Successor      string `json:"successor"`       // 紧随其后的区块哈希（若为最新区块则为空）
	Confirmed      string `json:"confirmed"`       // 区块是否已被确认（"true"/"false"）
	Contents       struct {
		Type           string `json:"type"`            // 区块类型，通常为 "state"
		Account        string `json:"account"`         // 账户地址
		Previous       string `json:"previous"`        // 上一个区块的哈希（open 区块时为空）
		Representative string `json:"representative"`  // 当前账户代表地址
		Balance        string `json:"balance"`         // 当前账户余额（raw）
		Link           string `json:"link"`            // 链接字段，send 为对方公钥；receive/open 为来源区块哈希
		LinkAsAccount  string `json:"link_as_account"` // 链接字段解析为账户地址
		Signature      string `json:"signature"`       // 区块签名，使用账户私钥签署
		Work           string `json:"work"`            // 工作量证明，满足难度要求
	} `json:"contents"`
	Subtype string `json:"subtype"` // 区块子类型（send, receive, open, change, epoch）
}
