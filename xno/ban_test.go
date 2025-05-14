package xno

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/bip39hd"
	"github.com/stretchr/testify/require"
	"testing"
)

// audit zero run country thunder fog ocean jealous settle yard lava lend party myth vacant green palm spot sugar ready body main offer share
// 0 ban_1wjpjiewc7czooe8de1xs93bknc733psan17pfzb3x56w7gwmoake6q4689p a1732fbc19257980f8dc19bf8fd6005a0c1b65a3532fec700c0a6eb99d42e07c 72368c19c5155fad5865b01dc9c2995145086d945005b37e90f464e15dc9d512
// 1 ban_379r48swnhiecac3pkx89n6h4hoioeiu7nke3pfzhypec65jn7mjoqsxfnyb 6a9134cf73c8bcbcf0dbc54cd6911fc77d738bc25a3546ca3bf8aa5faddc31e7 94f811b3ca3e0c52141b4ba63d08f13eb0ab21b2d24c0d9bf7facc51071a1671
// 2 ban_38ihyf4d9haasebnkt31ukhfcnqgqcxpkem8keo3g5furrtqini34hqqkiew b993d56716807270b0e63cef0688a8f539db46aa9f5cca69143938165e835ec2 9a0ff344b3bd08cb13496820dc9ed552eebabb693266932a170dbbc635785201
// 3 ban_3ncqrp5kbd8bcxermnkteuytbk4crjcsoifzyz85sgsf1mu7d55uo79dgt6u 232c1ee81710dca22a91738128743a16c245f398b97309d67b7a5fd82f06bb6a d157c58724acc9575989d25a66fda4c84ac4559ac1bff7cc3cbb2d04f6558c7b
// 4 ban_3573zrkfak9nsod7kneu57nafei49a98pckuj37ais5jkjh9oszsjyzto4y4 107c88a61e63b8536aa32c3bcf0b07e420e9ded46e00a4b717e7ce2b50309107 8ca1fe24d448f4cd5659519b196886b2023a0e6b2a5b884a886471945e7ae7f9
func TestBanMnemonic(t *testing.T) {
	mnemonic, err := bip39hd.GenerateMnemonicBy24()
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
	mnemonic := "audit zero run country thunder fog ocean jealous settle yard lava lend party myth vacant green palm spot sugar ready body main offer share"
	seed, _ := NewBip39Seed(mnemonic, "")

	t.Log(hex.EncodeToString(seed))
	// b6b2785ef52d4e970f99eb0b0c825d6fb1644fe823307e0a34bb991ab8a00040
	t.Log(mnemonic)
	for i := 0; i < 5; i++ {
		pk, address, err := ImportWallet(mnemonic, Banano, i)
		if nil != err {
			t.Fatal(err)
		}
		t.Log(fmt.Sprintf("%d %s %s", i, address, pk))
	}

	for i := 0; i < 5; i++ {
		pk, address, err := ImportWallet(mnemonic, Nano, i)
		if nil != err {
			t.Fatal(err)
		}
		t.Log(fmt.Sprintf("%d %s %s", i, address, pk))
	}
}

func TestBan(t *testing.T) {
	pk := "6a25d9b837a4a12b1ac77912d6bbe61be7fad8b67e33170dfc2cbb33110b3c6f"
	addr, err := PrivateKeyToBanAddressStr(pk)
	t.Log(err)
	t.Log(addr)
}
