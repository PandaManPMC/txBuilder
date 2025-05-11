package xno

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Client is used for connecting to http rpc endpoints.
type Client struct {
	URL        string
	AuthHeader string
	Ctx        context.Context
}

func (c *Client) Send(body interface{}) (result []byte, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(body); err != nil {
		return
	}
	if c.Ctx == nil {
		c.Ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(c.Ctx, http.MethodPost, c.URL, &buf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if c.AuthHeader != "" {
		req.Header.Set("Authorization", c.AuthHeader)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	buf.Reset()
	if _, err = io.Copy(&buf, resp.Body); err != nil {
		return
	}
	if err = resp.Body.Close(); err != nil {
		return
	}
	var v struct{ Error, Message string }
	if err = json.Unmarshal(buf.Bytes(), &v); err != nil {
		return
	}
	if v.Error != "" {
		err = errors.New(v.Error)
	} else if v.Message != "" {
		err = errors.New(v.Message)
	}
	return buf.Bytes(), err
}

func (c *Client) ActiveDifficulty() (*ActiveDifficultyRsp, error) {
	params := make(map[string]any)
	params["action"] = "active_difficulty"

	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}

	rsp := new(ActiveDifficultyRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

// WorkGenerate xno 的 pow，不同交易有不同级别的难度
// Epoch version	Block Subtype	Difficulty Threshold
// 1	All	ffffffc000000000
// 2	Send or change	fffffff800000000
// 2	Receive, open or epoch	fffffe0000000000
// 不同 pow 使用的 hash 不同，发送交易hash=上一个块 hash，接收交易hash=收款 publicKey 的 []byte
func (c *Client) WorkGenerate(hash string, cmd string) (*WorkGenerateRsp, error) {
	difficulty := "ffffffc000000000"
	switch cmd {
	case "send":
		fallthrough
	case "change":
		difficulty = "fffffff800000000"
	case "receive":
		fallthrough
	case "open":
		fallthrough
	case "epoch":
		difficulty = "fffffe0000000000"
	}
	workGenerate := make(map[string]any)
	workGenerate["action"] = "work_generate"
	workGenerate["hash"] = hash
	workGenerate["difficulty"] = difficulty

	res, err := c.Send(workGenerate)
	if nil != err {
		return nil, err
	}

	rsp := new(WorkGenerateRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) WorkGenerateSend(hash string) (*WorkGenerateRsp, error) {
	return c.WorkGenerate(hash, "send")
}

func (c *Client) WorkGenerateReceive(hash string) (*WorkGenerateRsp, error) {
	return c.WorkGenerate(hash, "receive")
}

func (c *Client) AccountInfo(address string) (*AccountInfoRsp, error) {
	params := make(map[string]any)
	params["action"] = "account_info"
	params["account"] = address
	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(AccountInfoRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) AccountBalance(address string) (*AccountBalanceRsp, error) {
	params := make(map[string]any)
	params["action"] = "account_balance"
	params["account"] = address
	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(AccountBalanceRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) AccountsPending(address string) (*AccountsPendingRsp, error) {
	params := make(map[string]any)
	params["action"] = "accounts_pending"
	params["accounts"] = []string{address}
	params["count"] = "10"
	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(AccountsPendingRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) BlockInfo(hash string) (*BlockInfoRsp, error) {
	params := make(map[string]any)
	params["action"] = "block_info"
	params["json_block"] = "true"
	params["hash"] = hash
	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(BlockInfoRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) ProcessSend(block Block) (*ProcessRsp, error) {
	params := make(map[string]any)
	params["action"] = "process"
	params["json_block"] = "true"
	params["subtype"] = "send"
	params["block"] = block

	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(ProcessRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}

func (c *Client) ProcessReceive(block Block) (*ProcessRsp, error) {
	params := make(map[string]any)
	params["action"] = "process"
	params["json_block"] = "true"
	params["subtype"] = "receive"
	params["block"] = block

	res, err := c.Send(params)
	if nil != err {
		return nil, err
	}
	println(string(res))
	rsp := new(ProcessRsp)
	if e := json.Unmarshal(res, rsp); nil != e {
		return nil, e
	}
	return rsp, nil
}
