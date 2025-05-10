package xno

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// repeat need bless tumble stay entire dinosaur dice arch goat nuclear term bind child party gravity lawn faculty confirm cram process exotic absurd aisle
// 0 ban_3cahum8tsrgt79w9zc6dzwstdqz3saoyqkdbiax5n53y4y7tc9paserrug6k 4a6208511c80ef8696c015dc6b16a07e1ae13ac2d7271f7c23e9c1e5c4725c92 a90fdccdace1da29f87fa88bff33a5dfe1ca2bebc969823a3a0c3e178ba51ec8
// 1 ban_398o4crk4gbexm1fuk1sfkkrgjow8xm93qykzwdi58r7ekbpbsxnokuy5b3x 79720a327fe7bb0c4db390fc0809248c7c04b97627bb56cb9812b4a5431e2c0b 9cd512b121392cecc0ddc8196ca58746bc376670dfd2ff17019b05649364e7b4
// 2 ban_3awwwpf3zkxm94d7135fiefjzztaacgu5j3de145quethdu8qt6xbffsjx4r 819741a42fcda2171392e261191f7fd4e8aac260d37552011b8bc3ade6648081 a39ce59a1fcbb3389650046d831b1fff48429db1c42b60043bed9a7af66be89d
// 3 ban_3unz6aca3aqx968td5a1gnfhodsm9k5i8qc6skoe68fqrucgxisc485r33ga fe07437c45817431bd46a603ca2d78546cbfd65c3a436d13005f931f5a43b523 ee9f221480a2fd390da58d00751afaaf333c87035d44ccaac219b7c6d4eec32a
// 4 ban_1xexkdwestugqipmf94mgrojmnmp7its3ujwzkdhzu5k6fcab945hqh81uxm 5ce87e721ab7d0606330db0b93008db1eaf20e7d1a06ac233c88374b1a32e236 759d92f8cceb6ebc2d369c53762b19d2762c3590ee3cfc96ffec722354849c43
func TestBanMnemonic(t *testing.T) {
	mnemonic, err := GenerateMnemonicBy24()
	require.Nil(t, err)
	t.Log(mnemonic)

	seed, err := NewBip39Seed(mnemonic, "")
	require.Nil(t, err)

	for i := 0; i < 5; i++ {
		key, err := DeriveBip39Key(seed, uint32(i))
		require.Nil(t, err)

		keyString := hex.EncodeToString(key)
		key, err = hex.DecodeString(keyString)
		require.Nil(t, err)

		pubKey, _, err := DeriveKeypair(key)
		require.Nil(t, err)
		address, err := PubKeyToBanAddress(pubKey)
		require.Nil(t, err)
		t.Log(fmt.Sprintf("%d %s %s %s", i, address, keyString, hex.EncodeToString(pubKey)))
	}
}

func TestImportWallet(t *testing.T) {
	mnemonic := "repeat need bless tumble stay entire dinosaur dice arch goat nuclear term bind child party gravity lawn faculty confirm cram process exotic absurd aisle"
	for i := 0; i < 5; i++ {
		pk, address, err := ImportWallet(mnemonic, Banano, i)
		if nil != err {
			t.Fatal(err)
		}
		t.Log(fmt.Sprintf("%d %s %s", i, address, pk))
	}
}
