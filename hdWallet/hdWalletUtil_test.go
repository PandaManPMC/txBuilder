package hdWallet

import (
	"testing"
)

func TestMnemonic(t *testing.T) {

	mnemonic, err := GetInstanceByHDWalletUtil().GenerateMnemonicBy12()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	mnemonic, err = GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	hd, err := GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	pri, err := GetInstanceByHDWalletUtil().WalletPrivateKey(hd, 0)
	if nil != err {
		t.Fatal(err)
	}

	pri_s := GetInstanceByHDWalletUtil().PriKeyToHexString(pri)
	t.Log(pri_s)

	l_pri, err := GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(pri_s)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(l_pri)

}
