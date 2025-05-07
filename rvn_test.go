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
	"testing"
)

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
}
