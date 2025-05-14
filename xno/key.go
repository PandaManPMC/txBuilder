package xno

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/PandaManPMC/txBuilder/xno/bip32"
	"github.com/PandaManPMC/txBuilder/xno/ed25519"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

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
		address, err = PubKeyToXNOAddress(pubKey)
		if nil != err {
			return "", "", err
		}
		return keyString, address, nil
	} else if Banano == netWork {
		address, err = PubKeyToBanAddress(pubKey)
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
	address, err := PubKeyToXNOAddress(pubKey)
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
	address, err := PubKeyToBanAddress(pubKey)
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
