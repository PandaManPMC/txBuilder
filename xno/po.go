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
