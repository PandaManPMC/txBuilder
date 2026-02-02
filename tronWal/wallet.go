package tronWal

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/hdWallet"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
	"strings"
)

func PubKeyToAddressTron(pubKey ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(pubKey)
	addressTron := append([]byte{0x41}, address.Bytes()...)
	return EncodeCheck(addressTron)
}

func PriKeyToAddressTron(privateKey *ecdsa.PrivateKey) string {
	return PubKeyToAddressTron(privateKey.PublicKey)
}

func EncodeCheck(input []byte) string {
	h256h0 := sha256.New()
	h256h0.Write(input)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	inputCheck := input
	inputCheck = append(inputCheck, h1[:4]...)

	return base58.Encode(inputCheck)
}

// ValidAddress 地址验证
func ValidAddress(addr string) bool {
	if len(addr) != 34 {
		return false
	}
	if string(addr[0:1]) != "T" {
		return false
	}
	_, err := DecodeCheck(addr)
	return err == nil
}

func DecodeCheck(input string) ([]byte, error) {
	decodeCheck := base58.Decode(input)

	if len(decodeCheck) < 4 {
		return nil, fmt.Errorf("addres base58 not check ok")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {
		return decodeData, nil
	}

	return nil, fmt.Errorf("addres hash not check ok")
}

// HexToTronAddress tronWal 的 hex 地址转为 tron 地址
func HexToTronAddress(hexAddress string) (string, error) {
	b, e := hex.DecodeString(hexAddress)
	if nil != e {
		return "", e
	}
	return EncodeCheck(b), nil
}

func TronAddressToHex(address string) (string, error) {
	decoded, err := DecodeCheck(address)
	if nil != err {
		return "", err
	}
	hexString := hex.EncodeToString(decoded)
	return hexString, nil
}

func HexAddressPadded64(address string) (string, error) {
	hexString, err := TronAddressToHex(address)
	if nil != err {
		return "", err
	}
	// 左侧补零到 64 位长度
	paddedHex := strings.Repeat("0", 64-len(hexString)) + hexString
	return paddedHex, nil
}

func IntToHexPadded64(params int64) string {
	hexStr := strconv.FormatInt(params, 16)
	paddedHex := strings.Repeat("0", 64-len(hexStr)) + hexStr
	return paddedHex
}

func TronAddressByPrivateKey(privateKey *ecdsa.PrivateKey) string {
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	addressTron := append([]byte{0x41}, address.Bytes()...)
	return EncodeCheck(addressTron)
}

func ImportWallet(mnemonic string, index int) (privateKey *ecdsa.PrivateKey, address string, err error) {
	wallet, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		return nil, "", fmt.Errorf("failed to import mnemonic: %v", err)
	}

	privateKey, err = hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(wallet, hdWallet.TRONHDCoinType, index)
	if nil != err {
		return nil, "", fmt.Errorf("failed to WalletPrivateKey: %v", err)
	}

	address = TronAddressByPrivateKey(privateKey)
	return
}
