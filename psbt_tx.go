package txBuilder

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

type TxInput struct {
	TxId              string
	VOut              uint32
	Sequence          uint32
	Amount            int64
	Address           string
	PrivateKey        string
	NonWitnessUtxo    string
	MasterFingerprint uint32
	DerivationPath    string
	PublicKey         string
}
type TxInputs []*TxInput

func (inputs TxInputs) UtxoViewpoint(net *chaincfg.Params) (UtxoViewpoint, error) {
	view := make(UtxoViewpoint, len(inputs))
	for _, v := range inputs {
		h, err := chainhash.NewHashFromStr(v.TxId)
		if err != nil {
			return nil, err
		}
		changePkScript, err := AddrToPkScript(v.Address, net)
		if err != nil {
			return nil, err
		}
		view[wire.OutPoint{Index: v.VOut, Hash: *h}] = changePkScript
	}
	return view, nil
}

type TxOutput struct {
	Address           string
	Amount            int64
	IsChange          bool
	MasterFingerprint uint32
	DerivationPath    string
	PublicKey         string
}

type ToSignInput struct {
	Index              int    `json:"index"`
	Address            string `json:"address"`
	PublicKey          string `json:"publicKey"`
	SigHashTypes       []int  `json:"sighashTypes"`
	DisableTweakSigner bool   `json:"disableTweakSigner"`
}

type SignPsbtOption struct {
	AutoFinalized bool           `json:"autoFinalized"`
	ToSignInputs  []*ToSignInput `json:"toSignInputs"`
}

func AddrToPkScript(addr string, network *chaincfg.Params) ([]byte, error) {
	address, err := btcutil.DecodeAddress(addr, network)
	if err != nil {
		return nil, err
	}

	return txscript.PayToAddrScript(address)
}

func PayToPubKeyHashScript(pubKeyHash []byte) ([]byte, error) {
	return txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).AddOp(txscript.OP_HASH160).
		AddData(pubKeyHash).AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).
		Script()
}

func PayToWitnessPubKeyHashScript(pubKeyHash []byte) ([]byte, error) {
	return txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(pubKeyHash).Script()
}

type OutPoint struct {
	TxId string `json:"txId"`
	VOut uint32 `json:"vOut"`
}

type PsbtInput struct {
	TxId   string `json:"txId"`
	Amount int64  `json:"amount"`
	VOut   uint32 `json:"vOut"`
}

type PsbtOutput struct {
	VOut     uint32 `json:"vOut"`
	Amount   int64  `json:"amount"`
	Address  string `json:"address"`
	PkScript string `json:"pkScript"`
}

type PsbtInputOutputs struct {
	UnSignedTx string        `json:"un_signed_tx"`
	Input      []*PsbtInput  `json:"input"`
	Output     []*PsbtOutput `json:"output"`
}

type GenerateMPCPSbtTxRes struct {
	Psbt         string   `json:"psbt"`
	PsbtTx       string   `json:"psbtTx"`
	SignHashList []string `json:"signHashList"`
}
