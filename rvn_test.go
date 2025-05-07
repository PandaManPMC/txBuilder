package txBuilder

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/ravencoin/btcec"
	"github.com/PandaManPMC/txBuilder/ravencoin/chaincfg"
	"github.com/PandaManPMC/txBuilder/ravencoin/chaincfg/chainhash"
	"github.com/PandaManPMC/txBuilder/ravencoin/txscript"
	"github.com/PandaManPMC/txBuilder/ravencoin/wire"
	"github.com/PandaManPMC/txBuilder/ravenutil"
	"github.com/PandaManPMC/txBuilder/rvnNetParams"
	"testing"
)

func TestRVNSignTxLegacyCompressedPKStr(t *testing.T) {
	pk := "7e53ae5d15dd6af9601fb0cbc6ce0ecda62fa8e56a4620b402a8a1061e648b87"

	tx := NewTxBuild(wire.TxVersion, &rvnNetParams.MainNetParams)
	tx.AddInput("760544b0faf975c4c0a2d908f5ff6109e9dd2b42dfe3f9704b265760937b32bc", 0,
		"76a914b5c2e59a0ce4f5c8090fc3a5c835f2bea331e2a388ac", "", "", 60000000)

	tx.AddOutput("RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a", 59000000)
	txHex, txId, err := RVNSignTxLegacyCompressedPKStr(tx, pk)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(txHex)
	t.Log(txId)
}

func TestSimplifiedSignTx(t *testing.T) {
	// 私钥和对应地址
	privKeyBytes, _ := hex.DecodeString("7e53ae5d15dd6af9601fb0cbc6ce0ecda62fa8e56a4620b402a8a1061e648b87")
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)

	// 已知 UTXO 信息
	txidStr := "7ce072ae8c059a6431210a11aeb2d2da08141d9d7c1fed5389a5935a398c6f9c"
	txid, _ := chainhash.NewHashFromStr(txidStr)
	vout := uint32(0)
	amount := int64(60000000) // 0.6 RVN
	prevPkScript, _ := hex.DecodeString("76a914b5c2e59a0ce4f5c8090fc3a5c835f2bea331e2a388ac")

	// 构造交易
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	redeemTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txid, vout), nil))

	// 输出地址与金额
	toAddr, _ := ravenutil.DecodeAddress("RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a", &chaincfg.MainNetParams)
	toPkScript, _ := txscript.PayToAddrScript(toAddr)
	redeemTx.AddTxOut(wire.NewTxOut(59000000, toPkScript)) // 0.59 RVN

	// 签名
	lookupKey := func(a ravenutil.Address) (*btcec.PrivateKey, bool, error) {
		return privKey, true, nil
	}
	sigScript, err := txscript.SignTxOutput(&chaincfg.MainNetParams,
		redeemTx, 0, prevPkScript, txscript.SigHashAll,
		txscript.KeyClosure(lookupKey), nil, nil)
	if err != nil {
		t.Fatal("SignTxOutput error:", err)
	}
	redeemTx.TxIn[0].SignatureScript = sigScript

	// 可选：验证签名是否正确
	vm, err := txscript.NewEngine(prevPkScript, redeemTx, 0, txscript.StandardVerifyFlags, nil, nil, amount)
	if err != nil {
		t.Fatal("Script engine error:", err)
	}
	if err := vm.Execute(); err != nil {
		t.Fatal("Script verification failed:", err)
	}

	// 输出结果
	var buf bytes.Buffer
	_ = redeemTx.Serialize(&buf)
	fmt.Printf("Raw transaction hex: %x\n", buf.Bytes())
	fmt.Printf("Transaction ID (txid): %s\n", redeemTx.TxHash().String())
}

func TestExampleSignTxStand(t *testing.T) {
	privKeyBytes, err := hex.DecodeString("7e53ae5d15dd6af9601fb0cbc6ce0ecda62fa8e56a4620b402a8a1061e648b87")
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	pubKeyHash := ravenutil.Hash160(pubKey.SerializeCompressed())
	addr, err := ravenutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(addr)

	// 构造 originTx，只为了提供签名脚本
	originTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&chainhash.Hash{}, ^uint32(0))
	txIn := wire.NewTxIn(prevOut, []byte{txscript.OP_0, txscript.OP_0})
	originTx.AddTxIn(txIn)

	// 这个是你提供的 UTXO 的 pkScript（hex）
	prevPkScriptHex := "76a914b5c2e59a0ce4f5c8090fc3a5c835f2bea331e2a388ac"
	pkScript, _ := hex.DecodeString(prevPkScriptHex)
	txOut := wire.NewTxOut(60000000, pkScript) // 0.6 RVN
	originTx.AddTxOut(txOut)
	originTxHash, _ := chainhash.NewHashFromStr("7ce072ae8c059a6431210a11aeb2d2da08141d9d7c1fed5389a5935a398c6f9c")

	// 创建实际要发送的交易
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut = wire.NewOutPoint(originTxHash, 0)
	txIn = wire.NewTxIn(prevOut, nil)
	redeemTx.AddTxIn(txIn)

	// 接收地址
	toAddrStr := "RPVW6ifbuCr4BQqAAjZ4APaT3sqL61tU8a"
	toAddr, err := ravenutil.DecodeAddress(toAddrStr, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Invalid to address:", err)
		return
	}
	toPkScript, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		fmt.Println("PayToAddrScript error:", err)
		return
	}
	txOut = wire.NewTxOut(59000000, toPkScript)
	redeemTx.AddTxOut(txOut)

	// 签名
	lookupKey := func(a ravenutil.Address) (*btcec.PrivateKey, bool, error) {
		return privKey, true, nil
	}
	sigScript, err := txscript.SignTxOutput(&chaincfg.MainNetParams,
		redeemTx, 0, pkScript, txscript.SigHashAll,
		txscript.KeyClosure(lookupKey), nil, nil)
	if err != nil {
		fmt.Println("SignTxOutput error:", err)
		return
	}
	redeemTx.TxIn[0].SignatureScript = sigScript

	// 可选：执行验证脚本
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(pkScript, redeemTx, 0, flags, nil, nil, 60000000)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := vm.Execute(); err != nil {
		fmt.Println("Script execution failed:", err)
		return
	}
	fmt.Println("Transaction successfully signed")

	// 输出交易 hex 和 txid
	var buf bytes.Buffer
	err = redeemTx.Serialize(&buf)
	if err != nil {
		fmt.Println("Serialize error:", err)
		return
	}
	fmt.Printf("Raw transaction hex: %x\n", buf.Bytes())
	fmt.Printf("Transaction ID (txid): %s\n", redeemTx.TxHash().String())

	//=== RUN   TestExampleSignTxStand
	//rvn_test.go:29: RRrFrujuR52sCMffc32nkLWmj362GwddJ3
	//Transaction successfully signed
	//Raw transaction hex: 01000000019c6f8c395a93a58953ed1f7c9d1d1408dad2b2ae110a2131649a058cae72e07c000000006a47304402207b84f0c356baefd14caad293f0a26ea77d6a18977e9ac5f711823262c651bc6902202ad5a507bfd4e44f0f0298fb48e136be03c39ae491fd45d85eb63a89d2f27146012102f8ddfcb601196c46c44cddce54eefb8e883fe50efcc5a6c54d6e584acd7a26a0ffffffff01c0448403000000001976a9149be5cba426fcde664640d936050b4892b9a1326788ac00000000
	//Transaction ID (txid): 210cc60f4a54b6867fd56bcc528ba1d93c00669d91c33e51aeabf65de9e94601
}
