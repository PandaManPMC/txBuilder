package xno

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/PandaManPMC/txBuilder/xno/bip32"
	"github.com/PandaManPMC/txBuilder/xno/ed25519"
	"github.com/PandaManPMC/txBuilder/xno/util"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

// GenerateMnemonicBits
//
//	熵位数 (bits)	校验位数 (bits)	总位数 (bits)	助记词单词数
//	128					4				132				12
//	160					5				165				15
//	192					6				198				18
//	224					7				231				21
//	256					8				264				24
func GenerateMnemonicBits(bits int) (string, error) {
	entropy, err := bip39.NewEntropy(bits)
	if nil != err {
		return "", fmt.Errorf("failed to generate entropy: %v", err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if nil != err {
		return "", fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return mnemonic, nil
}

// GenerateMnemonicBy12 生成助记词
func GenerateMnemonicBy12() (string, error) {
	// 128 位熵（生成 12 个单词的助记词）
	return GenerateMnemonicBits(128)
}

// GenerateMnemonicBy15 生成助记词
func GenerateMnemonicBy15() (string, error) {
	return GenerateMnemonicBits(160)
}

// GenerateMnemonicBy18 生成助记词
func GenerateMnemonicBy18() (string, error) {
	return GenerateMnemonicBits(192)
}

// GenerateMnemonicBy21 生成助记词
func GenerateMnemonicBy21() (string, error) {
	return GenerateMnemonicBits(224)
}

// GenerateMnemonicBy24 生成助记词
func GenerateMnemonicBy24() (string, error) {
	return GenerateMnemonicBits(256)
}

const (
	Nano   = "Nano"
	Banano = "Banano"
)

func ImportWallet(mnemonic, netWork string, offsetPath int) (pk, address string, err error) {
	seed, err := NewBip39Seed(mnemonic, "")
	key, err := DeriveBip39Key(seed, uint32(offsetPath))
	keyString := hex.EncodeToString(key)
	pubKey, _, err := DeriveKeypair(key)

	if Nano == netWork {
		address, err = util.PubKeyToXNOAddress(pubKey)
		if nil != err {
			return "", "", err
		}
		return keyString, address, nil
	} else if Banano == netWork {
		address, err = util.PubKeyToBanAddress(pubKey)
		if nil != err {
			return "", "", err
		}
		return keyString, address, nil
	}

	return "", "", errors.New("invalid network")
}

func PrivateKeyToXNOAddressStr(privateKey string) (string, error) {
	key, err := hex.DecodeString(privateKey)
	if nil != err {
		return "", err
	}
	pubKey, _, err := DeriveKeypair(key)
	if nil != err {
		return "", err
	}
	address, err := util.PubKeyToXNOAddress(pubKey)
	if nil != err {
		return "", err
	}
	return address, nil
}

func PrivateKeyToBanAddressStr(privateKey string) (string, error) {
	key, err := hex.DecodeString(privateKey)
	if nil != err {
		return "", err
	}
	pubKey, _, err := DeriveKeypair(key)
	if nil != err {
		return "", err
	}
	address, err := util.PubKeyToBanAddress(pubKey)
	if nil != err {
		return "", err
	}
	return address, nil
}

func DeriveKey(seed []byte, index uint32) (key []byte, err error) {
	if len(seed) != 32 {
		err = errors.New("seed must be 32 bytes")
		return
	}
	hash, err := blake2b.New256(nil)
	if err != nil {
		return
	}
	hash.Write(seed)
	if err = binary.Write(hash, binary.BigEndian, index); err != nil {
		return
	}
	return hash.Sum(nil), nil
}

func DeriveKeypair(key []byte) (pubkey, privkey []byte, err error) {
	return ed25519.GenerateKey(bytes.NewReader(key))
}

func NewBip39Seed(mnemonic, password string) (seed []byte, err error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}

func DeriveBip39Key(seed []byte, index uint32) (key []byte, err error) {
	key2, err := bip32.NewMasterKey(seed)
	if err != nil {
		return
	}
	for _, i := range []uint32{44, 165, index} {
		if key2, err = key2.NewChildKey(0x80000000 | i); err != nil {
			return
		}
	}
	return key2.Key, nil
}
