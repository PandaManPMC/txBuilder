package btcWal

import (
	"github.com/PandaManPMC/txBuilder/hdWallet"
	"testing"
)

func TestHDWallet(t *testing.T) {
	mne := "gown super smile wing hunt keep carpet stereo nurse umbrella case gun list fun valve stock job debate drip angry dumb tree finish lend"
	hd, err := hdWallet.GetInstanceByHDWalletUtil().ImportWalletFromMnemonic(mne)
	if nil != err {
		t.Fatal(err)
	}

	wallet, err := hdWallet.GetInstanceByHDWalletUtil().WalletPrivateKeyByCoinType(hd, 3, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(wallet)

	addr, err := GenerateLegacyAddress(wallet, 0x1E)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr) // DJo1dayr1qmRkkmVMHomY1uf3a2neU9fTc
}

func TestImportWallet(t *testing.T) {
	mne, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mne)
	//mne := "gown super smile wing hunt keep carpet stereo nurse umbrella case gun list fun valve stock job debate drip angry dumb tree finish lend"
	pk, addr, err := ImportWallet(mne, 3, 0)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(pk)
	t.Log(addr)
	t.Log(IsValidDOGEAddress(addr))
}

func TestHDWallet2(t *testing.T) {
	coinType := hdWallet.DOGEHDCoinType
	addressType := DogeAddress

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
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

	addr, err := GenerateLegacyAddress(privateKey, addressType)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))

	//wallet_test.go:115: flat globe fuel spend brother tornado pistol remember survey bless spread kitchen hill current deny rail crisp witness siren elder office leg aware seed
	//wallet_test.go:128: 1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd
	//wallet_test.go:135: DNyqb4rfhb3niqVGJm3tUjcZzZoybNNoxf
	//wallet_test.go:136: true

}

func TestGenerateAddressCompressed(t *testing.T) {
	coinType := hdWallet.DOGEHDCoinType
	addressType := DogeAddress

	mnemonic, err := hdWallet.GetInstanceByHDWalletUtil().GenerateMnemonicBy24()
	if nil != err {
		t.Fatal(err)
	}
	t.Log(mnemonic)
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

	addr, err := GenerateAddressCompressed(privateKey, addressType)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log("isBTCAddress=", IsValidBTCAddress(addr))
	t.Log("isLTCAddress=", IsValidLTCAddress(addr))
	t.Log("isDogeAddress=", IsValidDOGEAddress(addr))
}

func TestPrivateKetToAddr(t *testing.T) {
	privateKey := "1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd"

	pk, err := hdWallet.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey(privateKey)
	if nil != err {
		t.Fatal(err)
	}

	addr, err := GenerateLegacyAddress(pk, LTCAddress)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)
	t.Log(IsValidDOGEAddress(addr))

	addr, err = GenerateAddressCompressed(pk, RVNAddress)
	if nil != err {
		t.Fatal(err)
	}
	t.Log(addr)

}

func TestValidAddress(t *testing.T) {
	t.Log(IsValidLTCAddress("LVFhrz2p7qdNcRTEtqsLKpihmzFsZwaYjN"))
	t.Log(IsValidLTCAddress("ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklfz"))

	t.Log(IsValidLTCAddress("LLoRk2grZxuRkxaRvxVP7eSGWLwZGU1yx81"))
	t.Log(IsValidLTCAddress("LLoRk2grZxuRkxaRvxVP7eSGWLwZGU1yx0"))

	//格式看上去没问题，但 Bech32 地址的 最后6位是校验和，它通过 polymod 算法校验前面部分是否正确。
	t.Log(IsValidLTCAddress("ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklf1"))

	t.Log(IsValidRVNAddress("RBLfC7xp6PowS1a6qt9XwPp5ZqZouCvAFo"))

	t.Log("--------------- 符合规范的")
	t.Log(IsValidLTCAddress("LVFhrz2p7qdNcRTEtqsLKpihmzFsZwaYjN"))
	t.Log(IsValidLTCAddress("ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklfz"))
	t.Log("--------------- 不符合规范的")
	t.Log(IsValidLTCAddress("LVFhrz2p7qdNcRTEtqsLKpihmzFsZwaYj1"))
	t.Log(IsValidLTCAddress("ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklf1"))
	t.Log(IsValidLTCAddress("LVFhrz2p7qdNcRTEtqsLKpihmzFsZwaYjN1"))
	t.Log(IsValidLTCAddress("ltc1q4jamujuysk7mxm3qpzl387qghktcx07vmnklfz1"))
}
