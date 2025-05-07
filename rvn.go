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
)

func RVNSignTxLegacyCompressedPKStr(txBuild *TransactionBuilder, privateKeyStr string) (txHex, txId string, err error) {
	privateBytes, err := hex.DecodeString(privateKeyStr)
	if nil != err {
		return "", "", err
	}
	return RVNSignTxLegacyCompressed(txBuild, privateBytes)
}

// RVNSignTxLegacyCompressed 基于压缩的公钥地址，签名交易获得 hex
func RVNSignTxLegacyCompressed(txBuild *TransactionBuilder, privateBytes []byte) (txHex, txId string, err error) {
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateBytes)

	// 构造交易
	redeemTx := wire.NewMsgTx(wire.TxVersion)

	// 已知 UTXO 信息
	for _, v := range txBuild.inputs {
		txidStr := v.txId
		txid, err := chainhash.NewHashFromStr(txidStr)
		if nil != err {
			return "", "", err
		}
		vout := uint32(0)
		redeemTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(txid, vout), nil))
	}

	// 输出地址与金额
	for _, v := range txBuild.outputs {
		toAddr, _ := ravenutil.DecodeAddress(v.address, &chaincfg.MainNetParams)
		toPkScript, _ := txscript.PayToAddrScript(toAddr)
		redeemTx.AddTxOut(wire.NewTxOut(v.amount, toPkScript))
	}

	// 签名
	lookupKey := func(a ravenutil.Address) (*btcec.PrivateKey, bool, error) {
		return privateKey, true, nil
	}

	for inx, v := range txBuild.inputs {
		prevPkScript, _ := hex.DecodeString(v.privateKeyHex)
		sigScript, err := txscript.SignTxOutput(&chaincfg.MainNetParams,
			redeemTx, inx, prevPkScript, txscript.SigHashAll,
			txscript.KeyClosure(lookupKey), nil, nil)
		if nil != err {
			return "", "", err
		}
		redeemTx.TxIn[inx].SignatureScript = sigScript

		// 可选：验证签名是否正确
		vm, err := txscript.NewEngine(prevPkScript, redeemTx, inx, txscript.StandardVerifyFlags, nil, nil, v.amount)
		if nil != err {
			return "", "", err
		}
		if e := vm.Execute(); e != nil {
			return "", "", e
		}
	}

	// 输出结果
	var buf bytes.Buffer
	if e := redeemTx.Serialize(&buf); nil != e {
		return "", "", e
	}

	txHex = fmt.Sprintf("%x", buf.Bytes())
	txId = redeemTx.TxHash().String()
	return txHex, txId, nil
}
