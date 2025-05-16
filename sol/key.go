package sol

import (
	"encoding/hex"
	"fmt"
	"github.com/PandaManPMC/base58"
	"github.com/PandaManPMC/txBuilder/xno/bip32"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ed25519"
)

func ValidSolanaAddress(addr string) bool {
	decoded, err := base58.Decode(addr)
	if err != nil {
		return false
	}
	return len(decoded) == 32
}

func ImportWallet(mnemonic string, offsetPath int) (pk, address string, err error) {
	seed := bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)

	// 4. Derive: m/44'/501'/0'/0'
	purpose, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	coinType, _ := purpose.NewChildKey(bip32.FirstHardenedChild + 501)
	account, _ := coinType.NewChildKey(bip32.FirstHardenedChild + uint32(offsetPath))
	change, _ := account.NewChildKey(bip32.FirstHardenedChild + 0)

	privateKey := ed25519.NewKeyFromSeed(change.Key[:32])
	pub := privateKey.Public().(ed25519.PublicKey)

	address = base58.Encode(pub)

	return hex.EncodeToString(privateKey), address, nil
}

func ImportPrivateKeyFromHex(privateHex string) (ed25519.PrivateKey, error) {
	privateBytes, err := hex.DecodeString(privateHex)
	if err != nil {
		return nil, err
	}
	if len(privateBytes) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key length: %d", len(privateBytes))
	}
	return ed25519.PrivateKey(privateBytes), nil
}

func PrivateKeyToAddress(pk ed25519.PrivateKey) (string, error) {
	pub := pk.Public().(ed25519.PublicKey)
	address := base58.Encode(pub)
	return address, nil
}

func PrivateKeyStrToAddress(privateHex string) (string, error) {
	pk, err := ImportPrivateKeyFromHex(privateHex)
	if nil != err {
		return "", err
	}
	pub := pk.Public().(ed25519.PublicKey)
	address := base58.Encode(pub)
	return address, nil
}

// PrivateKeyHexToByteArray 私钥转为字节数组字符串，比如导入 Solflare 钱包就需要这个格式私钥
func PrivateKeyHexToByteArray(privateHex string) (string, error) {
	pk, err := ImportPrivateKeyFromHex(privateHex)
	if nil != err {
		return "", err
	}

	arr := "["
	for inx, v := range pk {
		arr = fmt.Sprintf("%s%d", arr, v)
		if inx != len(pk)-1 {
			arr = fmt.Sprintf("%s,", arr)
		}
	}
	arr = fmt.Sprintf("%s]", arr)

	return arr, nil
}

func Encode(b []byte) string {
	return base58.Encode(b)
}

func Decode(b string) ([]byte, error) {
	return base58.Decode(b)
}

func FindAssociatedTokenAddress(walletAddress, tokenMintAddress common.PublicKey) (common.PublicKey, uint8, error) {
	return common.FindAssociatedTokenAddress(walletAddress, tokenMintAddress)
}
