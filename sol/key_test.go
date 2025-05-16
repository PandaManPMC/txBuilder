package sol

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/PandaManPMC/base58"
	"testing"
)

// fire mixture square town record horror save angle sad stage charge pencil puppy rich draw write pistol curve upper lunch outside quality fish grow
// 0   3cmDFuaUDycjQAeGAHeNbo8KBbbFg3smgik65zRj3256   044293a67345ec60183ce431a615eb64b74dbbd9c6c38936706f15925869f62226e0b12c059ee743020f5abcb2879e610512d5dbb8cd1004788f2435c1829b41
// 1   D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn7Q   3da7df9f04628052017f6653909d67bab63eae098fb7fc7ed9a04d6371aa3587b29166ed467339a9f49c2a5b121bc92ea5e065d122ebb295d0ac3128dfef99ef
// 2   2eRJS1WHDEZjFK4333moyMW8H2d3k9D3k8ZaSsCnnAh1   f84ea91576ef6d80ce4c330cbdaa1be3e6910b9333f13e486c03da99b88cf63a1871a75035dc39441f8a387dbe66ef44da2ee444274baf65db972cecb7a0a60c
// 3   2qy1811C9mgVquHUrPuuzRrxFrp3Y7wh5QCZzVvQbZda   a1cc43533542b9259849b12834b53351a7436107204277cb65a7da99b402b1371b66e3344ae301cc032071acfbb3f5fe6a77c4566309ffbc8666a6953fc9ac29
// 4   2XhAmxTuzV63r3DXv68ZNaJK9dqjhvbQiQQJfNjEf3Tc   341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f
func TestImportWallet(t *testing.T) {
	//mnemonic, err := bip39hd.GenerateMnemonicBy24()
	//if nil != err {
	//	t.Fatal(err)
	//}

	mnemonic := "fire mixture square town record horror save angle sad stage charge pencil puppy rich draw write pistol curve upper lunch outside quality fish grow"
	t.Log(mnemonic)

	for i := 0; i < 5; i++ {
		pk, addr, err := ImportWallet(mnemonic, i)
		if nil != err {
			t.Fatal(err)
		}
		t.Log(i, " ", addr, " ", pk)
		//pkByte, err := hex.DecodeString(pk)
		//if nil != err {
		//	t.Fatal(err)
		//}
		//t.Log(pkByte)
	}
}

func TestPrivateKeyStrToAddress(t *testing.T) {
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	addr, err := PrivateKeyStrToAddress(privateKey)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr)
}

func TestImportPrivateKeyFromHex(t *testing.T) {
	privateKey := "341f58e5e65307ea627763a2d730deb6da5b95aca75ef5de2a4d916accd2dc3316b887a803bff07bcf4cd46281ba1cb6ada25fbf0419b3a9496cc20631ff058f"
	pk, err := ImportPrivateKeyFromHex(privateKey)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(pk)

	pub := pk.Public().(ed25519.PublicKey)
	address := base58.Encode(pub)
	t.Log(address)

	arr, err := PrivateKeyHexToByteArray(privateKey)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(arr)
	//[52,31,88,229,230,83,7,234,98,119,99,162,215,48,222,182,218,91,149,172,167,94,245,222,42,77,145,106,204,210,220,51,22,184,135,168,3,191,240,123,207,76,212,98,129,186,28,182,173,162,95,191,4,25,179,169,73,108,194,6,49,255,5,143]
	t.Log()

	sign := ed25519.Sign(pk, []byte{1, 2, 3})
	signS := hex.EncodeToString(sign)
	t.Log(signS)
}

func TestValidSolanaAddress(t *testing.T) {
	t.Log(ValidSolanaAddress("D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn7Q"))
	t.Log(ValidSolanaAddress("D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn71"))
	t.Log(ValidSolanaAddress("D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn7"))
	t.Log(ValidSolanaAddress("D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn776"))
	t.Log(ValidSolanaAddress("D24AYSaV2Kw27aHQuKjrFJMjmYyzPy6xb7Q2sXqsWn7_"))
}
