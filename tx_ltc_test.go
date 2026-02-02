package txBuilder

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PandaManPMC/txBuilder/ltcNetParams"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"io"
	"net/http"
	"testing"
)

// 该测试已完成
func TestSignTxLTCLegacyCompressed_OK(t *testing.T) {
	// transfer doge
	txBuild := NewTxBuild(1, &ltcNetParams.MainNetParams)
	txBuild.AddInput("e4cb23ca102de2fc041eb253871a9db4ae60e18dcc2d779299e197a2707f9aa2", 0,
		"76a9146e06119ee4cef69f0d514aedad93fefda2de3e2288ac", "", "", 500000)
	txBuild.AddInput("e4cb23ca102de2fc041eb253871a9db4ae60e18dcc2d779299e197a2707f9aa2", 1,
		"76a9146e06119ee4cef69f0d514aedad93fefda2de3e2288ac", "", "", 796523)
	txBuild.AddInput("11cac6b7b21aa072d6d037c342cd734b6235662355148408287f86f30b1bf952", 1,
		"76a9146e06119ee4cef69f0d514aedad93fefda2de3e2288ac", "", "", 696523)

	var amount int64 = 0
	for _, in := range txBuild.inputs {
		amount += in.amount
	}

	// 手续费
	var OneInputByteSize int64 = 148 // 一个输入占用字节
	var OneOutPutByteSize int64 = 34 // 一个输出占用字节
	var ltcBaseFee int64 = 20        // 每字节费用
	ltcInputFee := (int64(len(txBuild.inputs))*OneInputByteSize + 1) * ltcBaseFee
	ltcOutFee := 1 * OneOutPutByteSize * ltcBaseFee
	ltcFee := ltcInputFee + ltcOutFee

	transferAmount := amount - ltcFee
	t.Log("输入数量=", len(txBuild.inputs), "amount=", amount, " fee=", ltcFee, " transferAmount=", transferAmount)

	txBuild.AddOutput("ltc1qvc7uzxrsnzeyknyw5hkglsthe0pxm5ljathpj4", transferAmount)

	// 私钥
	privateBytes, _ := hex.DecodeString("e930ac0e8685333d2488241cfc7aac447b798325bdce633a426058dd68722962")
	prvKey, pubKey := btcec.PrivKeyFromBytes(privateBytes)

	fmt.Println("压缩公钥：", hex.EncodeToString(pubKey.SerializeCompressed()))
	fmt.Println("非压缩公钥：", hex.EncodeToString(pubKey.SerializeUncompressed()))

	compressedPubKey := pubKey.SerializeCompressed()
	hash160 := btcutil.Hash160(compressedPubKey)
	fmt.Printf("PubKeyHash: %x\n", hash160)                                         // ec56ca01cb5ad7fe845c55b85e1b2cbeb4641c95
	fmt.Printf("PubKeyHash: %x\n", btcutil.Hash160(pubKey.SerializeUncompressed())) // b02720b0c294de4a4cb5fc7cbf24bdd916e0f773

	pubKeyMap := make(map[int]string)
	for i := 0; i < len(txBuild.inputs); i++ {
		pubKeyMap[i] = hex.EncodeToString(pubKey.SerializeCompressed())
	}
	txHex, hashes, err := txBuild.UnSignedTx(pubKeyMap)

	if nil != err {
		t.Fatal(err)
	}

	if size, err := Size(txHex); nil != err {
		t.Fatal(err)
	} else {
		fmt.Println(size)
	}

	t.Log(hashes)
	signatureMap := make(map[int]string)
	for i, h := range hashes {
		sign := ecdsa.Sign(prvKey, RemoveZeroHex(h))
		signatureMap[i] = hex.EncodeToString(sign.Serialize())
	}

	txHex, err = SignTx(txHex, pubKeyMap, signatureMap)
	if err != nil {
		// todo
		fmt.Println(err)
	}
	t.Log(txHex)

	t.Log(CalcTxID(txHex))

	// https://go.getblock.io/de8e3dde043f430ea37e96e6199a146c
	type ReqData struct {
		JsonRpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  any    `json:"params"`
		Id      string `json:"id"`
	}
	reqData := ReqData{
		JsonRpc: "2.0",
		Method:  "sendrawtransaction",
		Params:  []string{txHex},
		Id:      "getblock.io",
	}

	inData, err := json.Marshal(reqData)
	if nil != err {
		t.Fatal(err)
	}

	endpoint := "https://go.getblock.io/xxx"
	res, err := post(endpoint, inData, nil)

	if nil != err {
		t.Log(string(res))
		t.Fatal(err)
	}

	t.Log(string(res))

	//	tx_ltc_test.go:42: 输入数量= 3 amount= 1993046  fee= 9580  transferAmount= 1983466
	//	压缩公钥： 03b9930e79b27c3e517323b915ec103058ac62c7c14099b2375311e72a7e90ef03
	//	非压缩公钥： 04b9930e79b27c3e517323b915ec103058ac62c7c14099b2375311e72a7e90ef0358bd97048f6f466b20c6c8428cd06ee712c7ee1b98d541bc62526fc8deba864b
	//	PubKeyHash: 6e06119ee4cef69f0d514aedad93fefda2de3e22
	//	PubKeyHash: 290c348cae6eaa228b77c39743ffafe8f767ac1b
	//	458
	//	tx_ltc_test.go:75: map[0:37e415dd2450b56dcbb829beb72d6726043f2edb7e109ed5c6af4c85279c9763 1:ef90d7e2191ebe5e81cc1eb5cb0715830eea86caa690e6f0f9569566fde3baba 2:c782b4eb7e046eddd5ea233072ed8418e1aac9fedcdaaeba3debad3b0aa0989d]
	//	tx_ltc_test.go:87: 0100000003a29a7f70a297e19992772dcc8de160aeb49d1a8753b21e04fce22d10ca23cbe4000000006b483045022100e9742946b11daae6dea1ad930a4dc4c1be22dcab39a3322525f9be0d9599334c022055ed93f96cce94afc77dc4ed8efb21c22fa97e4b91b506fc303868ccd5ab8125012103b9930e79b27c3e517323b915ec103058ac62c7c14099b2375311e72a7e90ef03ffffffffa29a7f70a297e19992772dcc8de160aeb49d1a8753b21e04fce22d10ca23cbe4010000006b483045022100aa68ef71abaa21eaaf2f2e82d061eb17185759fb691e6386fc0a4476bb70f38402201628974bf1c2897ad82b9886e991429bf5ab5e5208ea6aaedac0896d64a31149012103b9930e79b27c3e517323b915ec103058ac62c7c14099b2375311e72a7e90ef03ffffffff52f91b0bf3867f2808841455236635624b73cd42c337d0d672a01ab2b7c6ca11010000006b483045022100e4866623a288abd427e14a9484a4c7a345cf45c21dc0a5340d972143c2197f8402202a901260fd0b103d48fd95a758e0c90de4e2adb6a0f7c2e171851df7ee994e0d012103b9930e79b27c3e517323b915ec103058ac62c7c14099b2375311e72a7e90ef03ffffffff01ea431e0000000000160014663dc1187098b24b4c8ea5ec8fc177cbc26dd3f200000000
	//	tx_ltc_test.go:89: 65dd78f1c38691341be6a6da500967ebe2f2caf4c7d74354c0ce6a8a935cee2e <nil>
	//	tx_ltc_test.go:117: {"result":"65dd78f1c38691341be6a6da500967ebe2f2caf4c7d74354c0ce6a8a935cee2e","error":null,"id":"getblock.io"}
	//
	//	--- PASS: TestSignTxLTCLegacyCompressed (1.95s)
}

func post(url string, data []byte, header map[string]string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != err {
		return nil, err
	}

	if nil == header || 0 == len(header) {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if _, isOk := header["Content-Type"]; !isOk {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return body, errors.New(fmt.Sprintf("StatusCode=%d", resp.StatusCode))
	}

	return body, nil
}
