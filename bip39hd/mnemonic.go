package bip39hd

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonicBits
//
//	熵位数 (bits)	校验位数 (bits)	总位数 (bits)	助记词单词数
//	128					4				132				12
//	160					5				165				15
//	192					6				198				18
//	224					7				231				21
//	256					8				264				24
func GenerateMnemonicBits(bits int) (string, error) {
	entropy, err := bip39.NewEntropy(bits)
	if nil != err {
		return "", fmt.Errorf("failed to generate entropy: %v", err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if nil != err {
		return "", fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return mnemonic, nil
}

// GenerateMnemonicBy12 生成助记词
func GenerateMnemonicBy12() (string, error) {
	// 128 位熵（生成 12 个单词的助记词）
	return GenerateMnemonicBits(128)
}

// GenerateMnemonicBy15 生成助记词
func GenerateMnemonicBy15() (string, error) {
	return GenerateMnemonicBits(160)
}

// GenerateMnemonicBy18 生成助记词
func GenerateMnemonicBy18() (string, error) {
	return GenerateMnemonicBits(192)
}

// GenerateMnemonicBy21 生成助记词
func GenerateMnemonicBy21() (string, error) {
	return GenerateMnemonicBits(224)
}

// GenerateMnemonicBy24 生成助记词
func GenerateMnemonicBy24() (string, error) {
	return GenerateMnemonicBits(256)
}
