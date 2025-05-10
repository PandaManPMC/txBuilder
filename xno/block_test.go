package xno

import (
	"encoding/hex"
	"encoding/json"
	"github.com/PandaManPMC/txBuilder/xno/util"
	"math/big"
	"strings"
	"testing"
)

//proud dice elbow voyage cheap table trumpet evidence follow betray drink spray off begin manual often best crawl deal miss green electric sign black
//0 nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t 0953c1345c16fe3a7b8a46d7e063f6366118609f98e2ac85a987324f6dc36120 07369b6cb1cfe25f1600edef7707099df27a4dc0aadd9c65a53ea2329258110f
//1 nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f c444ff41d97498c78ba45709208be9b846565474133784f403101330e2e75281 32de710327a464e793404746f817b059fadb54f4394126f54798e4556d43368a
//2 nano_1mt73adt7u6r8a7c9o31x7catd1dbhu7qqkbwxhgys5re5gthg4t3socympf a164f407de55c6587805fccb2d762375cd09f7df875b5174fe78ab50484aad17 4f450a17a2ec98320aa3d420e9548d2c0b4bf65bde49e75eef647860dda7b85a
//3 nano_3dct1xutr1hhf3bj8wf7rpjiryjoocoqadfmkqaa5741zmwgem6eqgistrg6 ec3e6a7bc4c2351526222666737beac8e1d3ede9bc300064f589036ef07fdfda ad5a0777ac01ef68531371a5c5a30c7a35aaab742db395d0819440fcf8e64c8c
//4 nano_3dj8efspmh86dmgyw6du5kbiqe1oe7uxehibr38b7i4at6hwtmwpjohyzff9 2a07ba5c3fea1c9f6cda8f9ec7993ff28e4c510fb02ad2bc0799b97f5384484c ae26637369bcc45cddee117b1c930bb0156177d63e09c04c92c048d11fcd4f96

func TestSendBlock(t *testing.T) {
	balance := RawAmount{
		Int: big.Int{},
	}
	balance.SetString("900000000000000000000000000000", 10)

	previousHash := "CF8A877EB271293BF78E805C8C27DE408D009E01754CB211F7A48B1A3001F41E"
	previous, err := hex.DecodeString(previousHash)
	if nil != err {
		t.Fatal(err)
	}
	link, err := util.AddressToPubkey("nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f")
	if nil != err {
		t.Fatal(err)
	}

	block := Block{
		Type:           "state",
		Account:        "nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t",
		Previous:       previous,
		Representative: "nano_3patrick68y5btibaujyu7zokw7ctu4onikarddphra6qt688xzrszcg4yuo",
		Balance:        &balance,
		Link:           link,
		LinkAsAccount:  "nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f",
		Signature:      nil,
		Work:           nil,
	}

	privateKey := "0953c1345c16fe3a7b8a46d7e063f6366118609f98e2ac85a987324f6dc36120"

	if e := block.SignBlockStrPK(privateKey); nil != e {
		t.Fatal(e)
	}

	t.Log(hex.EncodeToString(block.Signature))

	client := Client{
		//URL:        "https://node.somenano.com/proxy",
		URL:        "http://127.0.0.1:7076",
		AuthHeader: "",
		Ctx:        nil,
	}

	wg, err := client.WorkGenerateSend(previousHash)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wg)

	work, err := hex.DecodeString(wg.Work)
	if nil != err {
		t.Fatal(err)
	}
	block.Work = work

	client = Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	rsp, err := client.ProcessSend(block)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestOpenBlock(t *testing.T) {
	balance := RawAmount{
		Int: big.Int{},
	}
	balance.SetString("100000000000000000000000000000", 10)
	previousHash := "0000000000000000000000000000000000000000000000000000000000000000"
	previous, err := hex.DecodeString(previousHash)
	if nil != err {
		t.Fatal(err)
	}

	linkHash := "D0AAA71AF24C3A872A6ED981F51BD847E669305BEA4B0721E23EDDF5969F4407"
	link, err := hex.DecodeString(linkHash)
	if nil != err {
		t.Fatal(err)
	}

	//nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t
	//D0AAA71AF24C3A872A6ED981F51BD847E669305BEA4B0721E23EDDF5969F4407
	//1 nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f c444ff41d97498c78ba45709208be9b846565474133784f403101330e2e75281 32de710327a464e793404746f817b059fadb54f4394126f54798e4556d43368a
	block := Block{
		Type:           "state",
		Account:        "nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f",
		Previous:       previous,
		Representative: "nano_3patrick68y5btibaujyu7zokw7ctu4onikarddphra6qt688xzrszcg4yuo",
		Balance:        &balance,
		Link:           link,
		LinkAsAccount:  "nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t",
		Signature:      nil,
		Work:           nil,
	}

	privateKey := "c444ff41d97498c78ba45709208be9b846565474133784f403101330e2e75281"

	if e := block.SignBlockStrPK(privateKey); nil != e {
		t.Fatal(e)
	}

	t.Log(hex.EncodeToString(block.Signature))

	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}

	pub, err := util.AddressToPubkey("nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f")
	if nil != err {
		t.Fatal(err)
	}

	//wg, err := client.WorkGenerate(linkHash)
	wg, err := client.WorkGenerateReceive(hex.EncodeToString(pub))
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wg)
	t.Log(wg.Work)

	work, err := hex.DecodeString(wg.Work)
	if nil != err {
		t.Fatal(err)
	}
	block.Work = work

	client = Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	rsp, err := client.ProcessReceive(block)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestActiveDifficulty(t *testing.T) {
	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	rsp, err := client.ActiveDifficulty()
	if nil != err {
		t.Fatal(err)
	}
	buf, err := json.Marshal(rsp)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(rsp)
	t.Log(string(buf))
}

func TestSendBlock2(t *testing.T) {
	privateKey := "c444ff41d97498c78ba45709208be9b846565474133784f403101330e2e75281"
	fromAddress := "nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f"
	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	accountInfo, err := client.AccountInfo(fromAddress)
	if nil != err {
		t.Fatal(err)
	}

	balance := RawAmount{
		Int: big.Int{},
	}
	balance.SetString("0", 10)

	previousHash := accountInfo.Frontier
	previous, err := hex.DecodeString(previousHash)
	if nil != err {
		t.Fatal(err)
	}
	link, err := util.AddressToPubkey("nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t")
	if nil != err {
		t.Fatal(err)
	}

	block := Block{
		Type:           "state",
		Account:        "nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f",
		Previous:       previous,
		Representative: "nano_3patrick68y5btibaujyu7zokw7ctu4onikarddphra6qt688xzrszcg4yuo",
		Balance:        &balance,
		Link:           link,
		LinkAsAccount:  "nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t",
		Signature:      nil,
		Work:           nil,
	}

	if e := block.SignBlockStrPK(privateKey); nil != e {
		t.Fatal(e)
	}

	t.Log(hex.EncodeToString(block.Signature))

	client = Client{
		//URL:        "https://node.somenano.com/proxy",
		URL:        "http://127.0.0.1:7076",
		AuthHeader: "",
		Ctx:        nil,
	}

	wg, err := client.WorkGenerateSend(previousHash)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wg)

	work, err := hex.DecodeString(wg.Work)
	if nil != err {
		t.Fatal(err)
	}
	block.Work = work

	client = Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	rsp, err := client.ProcessSend(block)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestReceiveBlock(t *testing.T) {
	privateKey := "0953c1345c16fe3a7b8a46d7e063f6366118609f98e2ac85a987324f6dc36120"

	//address := "nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t"
	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	//accountInfo, err := client.AccountInfo(address)
	//if nil != err {
	//	t.Fatal(err)
	//}

	balance := RawAmount{
		Int: big.Int{},
	}
	balance.SetString("1000000000000000000000000000000", 10)
	//previousHash := accountInfo.Frontier
	previousHash := "D0AAA71AF24C3A872A6ED981F51BD847E669305BEA4B0721E23EDDF5969F4407"
	previous, err := hex.DecodeString(previousHash)
	if nil != err {
		t.Fatal(err)
	}

	linkHash := "2C7B8068DD22CBC7A16F01C6114D4924DBEEC13DEF5AA4540423F5B6D37D2303"
	link, err := hex.DecodeString(linkHash)
	if nil != err {
		t.Fatal(err)
	}

	//nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t
	//D0AAA71AF24C3A872A6ED981F51BD847E669305BEA4B0721E23EDDF5969F4407
	//1 nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f c444ff41d97498c78ba45709208be9b846565474133784f403101330e2e75281 32de710327a464e793404746f817b059fadb54f4394126f54798e4556d43368a
	block := Block{
		Type:           "state",
		Account:        "nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t",
		Previous:       previous,
		Representative: "nano_3patrick68y5btibaujyu7zokw7ctu4onikarddphra6qt688xzrszcg4yuo",
		Balance:        &balance,
		Link:           link,
		LinkAsAccount:  "nano_1epyg63khb56wybn1jt8z1du1phtufchagc36utnh896copn8fncirmfe41f",
		Signature:      nil,
		Work:           nil,
	}

	if e := block.SignBlockStrPK(privateKey); nil != e {
		t.Fatal(e)
	}

	t.Log(hex.EncodeToString(block.Signature))

	client = Client{
		URL:        "http://127.0.0.1:7076",
		AuthHeader: "",
		Ctx:        nil,
	}

	//wg, err := client.WorkGenerate(linkHash)
	wg, err := client.WorkGenerateReceive(previousHash)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wg)
	t.Log(wg.Work)

	work, err := hex.DecodeString(wg.Work)
	if nil != err {
		t.Fatal(err)
	}
	block.Work = work

	ha, err := block.Hash()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(string(ha))
	t.Log(hex.EncodeToString(ha)) // 6f7cdcbb470df3c2609d838486c43f6b0e4f229904dbfb767ac754d616bf978e

	if "33D7A66F90766B090F4179E2A95320BBE1340571D2BA538DFC0B0ADACCFABB0C" == strings.ToUpper(hex.EncodeToString(ha)) {
		t.Log("hash 正确")
	}

	client = Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}

	// 33D7A66F90766B090F4179E2A95320BBE1340571D2BA538DFC0B0ADACCFABB0C
	rsp, err := client.ProcessReceive(block)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(rsp)

}

func TestAccountBalance(t *testing.T) {
	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	bal, err := client.AccountBalance("nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(bal)
}

func TestAccountsPending(t *testing.T) {
	client := Client{
		URL:        "https://node.somenano.com/proxy",
		AuthHeader: "",
		Ctx:        nil,
	}
	bal, err := client.AccountsPending("nano_13spmfpd5mz4dwd13uhhgw5im9hkhb8w3cpxmjktcho48cb7i6ahcri15q7t")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(bal)
}
