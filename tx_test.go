package txBuilder

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/dogeNetParams"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"testing"
)

func TestSignTxDogeLegacy(t *testing.T) {
	// transfer doge
	txBuild := NewTxBuild(1, &dogeNetParams.DogeMainNetParams)
	txBuild.AddInput("d912375ffefd88afd1e8cdf60eab02be7ca3847c18fec555ba740919ac99aca4", 0,
		"76a914b02720b0c294de4a4cb5fc7cbf24bdd916e0f77388ac", "", "", 100000000)
	txBuild.AddOutput("D747b4mYvNAejeHrp6jKdecGhRC5DLSTu3", 80000000)

	privateBytes, _ := hex.DecodeString("70d4946e5b6c5746c4b03f729f09a843df0b0da74335307202199be49982ae5d")
	prvKey, pubKey := btcec.PrivKeyFromBytes(privateBytes)
	fmt.Println("压缩公钥：", hex.EncodeToString(pubKey.SerializeCompressed()))
	fmt.Println("非压缩公钥：", hex.EncodeToString(pubKey.SerializeUncompressed()))

	compressedPubKey := pubKey.SerializeCompressed()
	hash160 := btcutil.Hash160(compressedPubKey)
	fmt.Printf("PubKeyHash: %x\n", hash160)                                         // ec56ca01cb5ad7fe845c55b85e1b2cbeb4641c95
	fmt.Printf("PubKeyHash: %x\n", btcutil.Hash160(pubKey.SerializeUncompressed())) // b02720b0c294de4a4cb5fc7cbf24bdd916e0f773

	pubKeyMap := make(map[int]string)
	pubKeyMap[0] = hex.EncodeToString(pubKey.SerializeUncompressed())
	txHex, hashes, err := txBuild.UnSignedTx(pubKeyMap)
	if size, err := Size(txHex); nil != err {

	} else {
		fmt.Println(size)
	}

	t.Log(hashes)
	signatureMap := make(map[int]string)
	for i, h := range hashes {
		sign := ecdsa.Sign(prvKey, RemoveZeroHex(h))
		signatureMap[i] = hex.EncodeToString(sign.Serialize())
	}

	txHex, err = SignTxByUncompressed(txHex, pubKeyMap, signatureMap)
	if err != nil {
		// todo
		fmt.Println(err)
	}
	fmt.Println(txHex)
	fmt.Println(CalcTxID(txHex))
}

func TestSignTxDogeLegacyCompressed(t *testing.T) {
	// transfer doge
	txBuild := NewTxBuild(1, &dogeNetParams.DogeMainNetParams)
	txBuild.AddInput("db7d96072a952774246be2fb18c1b0d5fd028ef1414807fa9f5a2a1ae35e9bdb", 0,
		"76a914cd60d88a6a618035555d869ca2ddf605b01522c588ac", "", "", 100000000)
	txBuild.AddOutput("D747b4mYvNAejeHrp6jKdecGhRC5DLSTu3", 92680000)

	privateBytes, _ := hex.DecodeString("639ce62c7bc26b27191b50acd1d5fbd8c732b04d377ce71229447b482a7067d0")
	prvKey, pubKey := btcec.PrivKeyFromBytes(privateBytes)
	fmt.Println("压缩公钥：", hex.EncodeToString(pubKey.SerializeCompressed()))
	fmt.Println("非压缩公钥：", hex.EncodeToString(pubKey.SerializeUncompressed()))

	compressedPubKey := pubKey.SerializeCompressed()
	hash160 := btcutil.Hash160(compressedPubKey)
	fmt.Printf("PubKeyHash: %x\n", hash160)                                         // ec56ca01cb5ad7fe845c55b85e1b2cbeb4641c95
	fmt.Printf("PubKeyHash: %x\n", btcutil.Hash160(pubKey.SerializeUncompressed())) // b02720b0c294de4a4cb5fc7cbf24bdd916e0f773

	pubKeyMap := make(map[int]string)
	pubKeyMap[0] = hex.EncodeToString(pubKey.SerializeCompressed())
	txHex, hashes, err := txBuild.UnSignedTx(pubKeyMap)

	if size, err := Size(txHex); nil != err {

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
	fmt.Println(txHex)
	fmt.Println(CalcTxID(txHex))
}
