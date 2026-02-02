package txBuilder

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// Base58 alphabet (Bitcoin-style)
const b58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// randomBytes returns n cryptographically-random bytes (ignores error for brevity)
func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

// ----------------- Base58 encode -----------------
func base58Encode(input []byte) string {
	// count leading zeros
	zero := byte(0)
	leadingZeros := 0
	for _, b := range input {
		if b == zero {
			leadingZeros++
		} else {
			break
		}
	}

	// convert to big.Int
	x := new(big.Int).SetBytes(input)
	mod := new(big.Int)
	base := big.NewInt(58)
	var chars []byte
	for x.Cmp(big.NewInt(0)) > 0 {
		x.DivMod(x, base, mod)
		chars = append(chars, b58Alphabet[mod.Int64()])
	}
	for i := 0; i < leadingZeros; i++ {
		chars = append(chars, b58Alphabet[0])
	}
	// reverse
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// ----------------- Base58Check (double SHA256 checksum) -----------------
func base58CheckEncode(version byte, payload []byte) string {
	data := append([]byte{version}, payload...)
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	checksum := h2[:4]
	full := append(data, checksum...)
	return base58Encode(full)
}

// ----------------- Address generators -----------------

// ETH: 0x + 40 hex chars (random)
func RandomETH() string {
	b := randomBytes(20) // 20 bytes = 40 hex chars
	return "0x" + hex.EncodeToString(b)
}

// SOL: Base58 of 32 bytes
func RandomSOL() string {
	b := randomBytes(32)
	return base58Encode(b)
}

// DOGE/LTC/RVN: Base58Check P2PKH-like with version byte and 20-byte payload
func RandomDOGE() string { return base58CheckEncode(0x1e, randomBytes(20)) } // typically starts with 'D'
func RandomLTC() string  { return base58CheckEncode(0x30, randomBytes(20)) } // typically starts with 'L' or 'M'
func RandomRVN() string  { return base58CheckEncode(0x3c, randomBytes(20)) } // Ravencoin typical ver

// TRX: 0x41 + 20 bytes, then Base58Check
func RandomTRX() string {
	payload20 := randomBytes(20)
	data := append([]byte{0x41}, payload20...)
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	checksum := h2[:4]
	full := append(data, checksum...)
	return base58Encode(full)
}
