package txBuilder

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/rvnNetParams"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func Size(txHex string) (int, error) {
	// 计算交易字节数
	txBytes, _ := hex.DecodeString(txHex)
	txReader := bytes.NewReader(txBytes)
	tx := &wire.MsgTx{}
	if err := tx.Deserialize(txReader); nil != err {
		return 0, err
	}
	// 获取交易大小
	txSize := tx.SerializeSize()
	return txSize, nil
}

func TXCompressedSizePkStr(txBuild *TransactionBuilder, privateKeyStr string) (size int, err error) {
	privateBytes, err := hex.DecodeString(privateKeyStr)
	if nil != err {
		return 0, err
	}
	return TXCompressedSize(txBuild, privateBytes)
}

// TXCompressedSize 压缩公钥组成的交易的 size，用于计算 fee
func TXCompressedSize(txBuild *TransactionBuilder, privateBytes []byte) (size int, err error) {
	_, pubKey := btcec.PrivKeyFromBytes(privateBytes)

	pubKeyMap := make(map[int]string)
	pubKeyMap[0] = hex.EncodeToString(pubKey.SerializeCompressed())
	txHex, _, err := txBuild.UnSignedTx(pubKeyMap)
	if nil != err {
		return 0, err
	}

	size, err = Size(txHex)
	return size, nil
}

func SignTxLegacyCompressedPKStr(txBuild *TransactionBuilder, privateKeyStr string) (txHex, txId string, err error) {
	if txBuild.netParams.Net == rvnNetParams.MainNetParams.Net {
		// rvn
		return RVNSignTxLegacyCompressedPKStr(txBuild, privateKeyStr)
	}
	privateBytes, err := hex.DecodeString(privateKeyStr)
	if nil != err {
		return "", "", err
	}
	return SignTxLegacyCompressed(txBuild, privateBytes)
}

// SignTxLegacyCompressed 基于压缩的公钥地址，签名交易获得 hex
func SignTxLegacyCompressed(txBuild *TransactionBuilder, privateBytes []byte) (txHex, txId string, err error) {
	prvKey, pubKey := btcec.PrivKeyFromBytes(privateBytes)

	pubKeyMap := make(map[int]string)
	for i := 0; i < len(txBuild.inputs); i++ {
		pubKeyMap[i] = hex.EncodeToString(pubKey.SerializeCompressed())
	}
	txHex, hashes, err := txBuild.UnSignedTx(pubKeyMap)
	if nil != err {
		return "", "", err
	}

	signatureMap := make(map[int]string)
	for i, h := range hashes {
		sign := ecdsa.Sign(prvKey, RemoveZeroHex(h))
		signatureMap[i] = hex.EncodeToString(sign.Serialize())
	}

	txHex, err = SignTx(txHex, pubKeyMap, signatureMap)
	if nil != err {
		return "", "", err
	}

	txId, err = CalcTxID(txHex)
	return txHex, txId, err
}

// SignTx 签名交易，地址默认是由压缩公钥生成
func SignTx(raw string, pubKeyMap map[int]string, signatureMap map[int]string) (string, error) {
	txBytes, err := hex.DecodeString(raw)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(txBytes)
	tx := &wire.MsgTx{}
	err = tx.Deserialize(reader)
	if err != nil {
		return "", err
	}

	if len(tx.TxIn) != len(signatureMap) {
		return "", fmt.Errorf("signature miss")
	}

	for i := 0; i < len(tx.TxIn); i++ {
		builder := txscript.NewScriptBuilder()
		publicKey, err := btcec.ParsePubKey(RemoveZeroHex(pubKeyMap[i]))
		if err != nil {
			return "", err
		}
		redeemScript := publicKey.SerializeCompressed()
		sig1 := append(RemoveZeroHex(signatureMap[i]), byte(txscript.SigHashAll))
		scriptBuilder, err := builder.AddData(sig1).AddData(redeemScript).Script()
		if err != nil {
			return "", err
		}
		tx.TxIn[i].SignatureScript = scriptBuilder
	}
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

// SignTxByUncompressed 签名交易，地址是非压缩公钥生成
func SignTxByUncompressed(raw string, pubKeyMap map[int]string, signatureMap map[int]string) (string, error) {
	txBytes, err := hex.DecodeString(raw)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(txBytes)
	tx := &wire.MsgTx{}
	err = tx.Deserialize(reader)
	if err != nil {
		return "", err
	}

	if len(tx.TxIn) != len(signatureMap) {
		return "", fmt.Errorf("signature miss")
	}

	for i := 0; i < len(tx.TxIn); i++ {
		builder := txscript.NewScriptBuilder()
		publicKey, err := btcec.ParsePubKey(RemoveZeroHex(pubKeyMap[i]))
		if err != nil {
			return "", err
		}
		redeemScript := publicKey.SerializeUncompressed()
		sig1 := append(RemoveZeroHex(signatureMap[i]), byte(txscript.SigHashAll))
		scriptBuilder, err := builder.AddData(sig1).AddData(redeemScript).Script()
		if err != nil {
			return "", err
		}
		tx.TxIn[i].SignatureScript = scriptBuilder
	}
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

// CalcTxID 获取交易hash
func CalcTxID(txHex string) (string, error) {
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return "", err
	}
	first := sha256.Sum256(txBytes)
	second := sha256.Sum256(first[:])
	// 注意：比特币/狗狗/莱特币使用 txid 的小端表示
	return hex.EncodeToString(reverseBytes(second[:])), nil
}

// 工具函数：字节反转（大端转小端）
func reverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
