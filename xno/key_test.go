package xno

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/xno/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeriveKey(t *testing.T) {
	seed, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001")
	key, err := DeriveKey(seed, 1)
	require.Nil(t, err)
	t.Log(hex.EncodeToString(key))
	assert.Equal(t, "1495f2d49159cc2eaaaa97ebb42346418e1268aff16d7fca90e6bad6d0965520", hex.EncodeToString(key))
}

func TestMnemonic(t *testing.T) {
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
		address, err := util.PubKeyToXNOAddress(pubKey)
		require.Nil(t, err)
		t.Log(fmt.Sprintf("%d %s %s %s", i, address, keyString, hex.EncodeToString(pubKey)))
	}
}

func TestBip39(t *testing.T) {
	//seed, err := newBip39Seed("edge defense waste choose enrich upon flee junk siren film clown finish "+
	//	"luggage leader kid quick brick print evidence swap drill paddle truly occur", "some password")
	seed, err := NewBip39Seed("edge defense waste choose enrich upon flee junk siren film clown finish "+
		"luggage leader kid quick brick print evidence swap drill paddle truly occur", "")
	require.Nil(t, err)
	key, err := DeriveBip39Key(seed, 0)
	t.Log(string(key))

	keyString := hex.EncodeToString(key)
	t.Log(keyString)

	key, err = hex.DecodeString(keyString)

	require.Nil(t, err)
	pubKey, _, err := DeriveKeypair(key)
	t.Log(string(pubKey))
	require.Nil(t, err)
	address, err := util.PubKeyToXNOAddress(pubKey)
	t.Log(address)
	require.Nil(t, err)
	//assert.Equal(t, "nano_1pu7p5n3ghq1i1p4rhmek41f5add1uh34xpb94nkbxe8g4a6x1p69emk8y1d", address)
}
