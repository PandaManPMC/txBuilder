package tronWal

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/hdWallet"
	"testing"
)

func TestHexToTronAddress(t *testing.T) {
	//tronWeb.address.fromHex("418840E6C55B9ADA326D211D818C34A994AECED808")
	//> "TNPeeaaFB7K9cmo4uQpcU32zGK8G1NYqeL"
	s := "418840E6C55B9ADA326D211D818C34A994AECED808"

	a, e := HexToTronAddress(s)
	if nil != e {
		t.Fatal(e)
	}
	t.Log(a)
}

func TestValid(t *testing.T) {
	t.Log(ValidAddress("TKjPqUwwT9Q4VPKH9BMKvgafxEyDWT9t8Z"))
}

func TestMnemonic(t *testing.T) {

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy12()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	mnemonic, err = hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	hd, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	pri, err := hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKey(hd, 0)
	if nil != err {
		t.Fatal(err)
	}

	pri_s := hdWallet.GetInstanceByHDWalletUtil().PriKeyToHexString(pri)
	t.Log(pri_s)

	addr := TronAddressByPrivateKey(pri)
	t.Log(addr)

	l_pri, err := hdWallet.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(pri_s)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(l_pri)

	addr = PriKeyToAddressTron(l_pri)
	t.Log(addr)
}

func TestMnemonic2(t *testing.T) {

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
	t.Log(len(mnemonic))

	pri, addr, err := ImportWallet(mnemonic, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr)

	pri_s := hdWallet.GetInstanceByHDWalletUtil().PriKeyToHexString(pri)
	t.Log(pri_s)

	addr = TronAddressByPrivateKey(pri)
	t.Log(addr)

	l_pri, err := hdWallet.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(pri_s)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(l_pri)

	addr = PriKeyToAddressTron(l_pri)
	t.Log(addr)
}

func TestTronAddressToHex(t *testing.T) {
	// Tron 地址
	address := "TVTV9aEDdszTNYayNBdjpQ7xfXH3DMyzXq"
	d, err := DecodeCheck(address)
	if nil != err {
		t.Fatal(err)
	}
	h := hex.EncodeToString(d)
	t.Log(h)

	// 转为 Hex
	hexAddress, err := HexAddressPadded64(address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hex 地址:", hexAddress)
}

func TestValidAddress(t *testing.T) {
	t.Log(ValidAddress("TVTV9aEDdszTNYayNBdjpQ7xfXH3DMyzXq"))
	t.Log(ValidAddress("TVTV9aEDdszTNYayNBdjpQ7xfXH3DMyzX1"))
	t.Log(ValidAddress("0xa07880f94796250e9b37F4aFbcbAeb1e55A385c1"))
	t.Log(ValidAddress("TVTV9aEDdszTNYayNBdjpQ7xfXH3DMyzX"))
	t.Log(ValidAddress(""))

}
