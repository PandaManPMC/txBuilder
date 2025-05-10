package xno

import (
	"encoding/base32"
	"errors"

	"golang.org/x/crypto/blake2b"
)

// AddressToPubkey converts address to a pubkey.
func AddressToPubkey(address string) (pubkey []byte, err error) {
	err = errors.New("invalid address")
	switch len(address) {
	case 64:
		if address[:4] != "xrb_" && address[:4] != "ban_" {
			return
		}
		address = address[4:]
	case 65:
		if address[:5] != "nano_" {
			return
		}
		address = address[5:]
	default:
		return
	}
	b32 := base32.NewEncoding("13456789abcdefghijkmnopqrstuwxyz")
	if pubkey, err = b32.DecodeString("1111" + address[:52]); err != nil {
		return
	}
	pubkey = pubkey[3:]
	checksum, err := checksum(pubkey)
	if err != nil {
		return
	}
	if b32.EncodeToString(checksum) != address[52:] {
		err = errors.New("checksum mismatch")
	}
	return
}

func PubKeyToXNOAddress(pubKey []byte) (address string, err error) {
	return PubKeyToAddress(pubKey, "nano")
}

func PubKeyToBanAddress(pubKey []byte) (address string, err error) {
	return PubKeyToAddress(pubKey, "ban")
}

// PubKeyToAddress converts pubkey to an address.
func PubKeyToAddress(pubKey []byte, prefix string) (address string, err error) {
	if len(pubKey) != 32 {
		return "", errors.New("invalid pubkey length")
	}
	checksum, err := checksum(pubKey)
	if err != nil {
		return
	}
	pubKey = append([]byte{0, 0, 0}, pubKey...)
	b32 := base32.NewEncoding("13456789abcdefghijkmnopqrstuwxyz")
	return prefix + "_" + b32.EncodeToString(pubKey)[4:] + b32.EncodeToString(checksum), nil
}

func checksum(pubkey []byte) (checksum []byte, err error) {
	hash, err := blake2b.New(5, nil)
	if err != nil {
		return
	}
	hash.Write(pubkey)
	for _, b := range hash.Sum(nil) {
		checksum = append([]byte{b}, checksum...)
	}
	return
}
