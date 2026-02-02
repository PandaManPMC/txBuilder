package ethWal

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/txBuilder/hdWallet"
	"golang.org/x/crypto/sha3"
	"testing"
)

func Test1(t *testing.T) {
	privateKey, err := hdWallet.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey("1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd")
	if nil != err {
		t.Fatal(err)
	}

	// 2. 获取公钥
	publicKey := privateKey.PublicKey

	// 3. 将公钥序列化为字节数组 (只需要 X 和 Y 坐标，去掉前缀 0x04)
	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)

	// 4. 对公钥进行 Keccak-256 哈希运算
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	hashed := hash.Sum(nil)

	// 5. 获取哈希的后 20 字节
	address := hashed[len(hashed)-20:]

	// 6. 转换为十六进制字符串
	fmt.Printf("Ethereum Address: 0x%s\n", hex.EncodeToString(address))
}

func TestPrivateKeyToAddressETH(t *testing.T) {
	privateKey, err := hdWallet.GetInstanceByHDWalletUtil().LoadWalletByPrivateKey("1ea107cf1e8cbca5a1e9ee2661505b1836db495c574eace64ecdbc20b29b83fd")
	if nil != err {
		t.Fatal(err)
	}
	t.Log(PrivateKeyToAddressETH(privateKey))
}

func TestIValidAddress(t *testing.T) {
	t.Log(ValidAddress("0xa07880f94796250e9b37F4aFbcbAeb1e55A385c1"))
	t.Log(ValidAddress("0xa07880f94796250e9b37F4aFbcbAeb1e55A385cp"))
	t.Log(ValidAddress("TVTV9aEDdszTNYayNBdjpQ7xfXH3DMyzXq"))
	t.Log(ValidAddress("0xa07880f94796250e9b37F4aFbcbAeb1e55A385c"))
	t.Log(ValidAddress(""))
}
