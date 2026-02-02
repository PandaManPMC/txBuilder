package btcWal

import (
	"github.com/PandaManPMC/txBuilder/hdWallet"
	"testing"
)

func TestHDWalletLTCGenerateLegacyAddress(t *testing.T) {
	coinType := hdWallet.LTCHDCoinType
	addressType := LTCAddress

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)

	mnemonic = "strike charge window soul vacuum midnight amused ill usage core claim innocent sure shop roast fat rebuild crazy battle result gather language inform promote"

	wallet, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	privateKey, err := hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(wallet, coinType, 0)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyStr := hdWallet.GetInstanceByHDWalletUtil().PriKeyToHexString(privateKey)

	t.Log(privateKeyStr)

	// LRwVGxbXyMSbLTiaDaWiHQjMRXuEC6b3rJ

	addr, err := GenerateLegacyAddress(privateKey, addressType)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))
}

func TestHDWalletLTCGenerateAddressCompressed(t *testing.T) {
	coinType := hdWallet.LTCHDCoinType
	addressType := LTCAddress

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)

	mnemonic = "upgrade report cotton echo side ahead ball april add north rice butter crouch filter author guitar case buddy merge festival chapter cushion shallow awful"

	wallet, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	privateKey, err := hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(wallet, coinType, 0)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyStr := hdWallet.GetInstanceByHDWalletUtil().PriKeyToHexString(privateKey)

	t.Log(privateKeyStr)

	// LdCb9yuteSFRMAPMA2xazgrSTvwh8wo2JV

	addr, err := GenerateAddressCompressed(privateKey, addressType)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))
}

func TestLTCImportWalletSegWitP2WPKH(t *testing.T) {
	mnemonic := "upgrade report cotton echo side ahead ball april add north rice butter crouch filter author guitar case buddy merge festival chapter cushion shallow awful"

	pk, addr, err := ImportWalletSegWitP2WPKH(mnemonic, hdWallet.LTCHDCoinType, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(pk)
	t.Log(addr)
	t.Log(IsValidDOGEAddress(addr))
}

func TestHDWalletLTCSegWitP2WPKH(t *testing.T) {
	coinType := hdWallet.LTCHDCoinType

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}

	mnemonic = "upgrade report cotton echo side ahead ball april add north rice butter crouch filter author guitar case buddy merge festival chapter cushion shallow awful"
	t.Log(mnemonic)

	wallet, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mnemonic)
	if nil != err {
		t.Fatal(err)
	}

	privateKey, err := hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKeySegWitP2WPKHByCoinType(wallet, coinType, 0)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyStr := hdWallet.GetInstanceByHDWalletUtil().PriKeyToHexString(privateKey)

	t.Log(privateKeyStr)

	//addr, err := GenerateSegWitP2WPKHAddress(privateKey, "ltc")
	addr, err := GenerateSegWitP2WPKHAddressByLTC(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))
}
