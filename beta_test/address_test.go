package beta_test

import (
	"fmt"
	"github.com/PandaManPMC/txBuilder"
	"github.com/PandaManPMC/txBuilder/ethWal"
	"github.com/PandaManPMC/txBuilder/tronWal"
	"testing"
)

func TestRandAddress(t *testing.T) {
	fmt.Println("ETH :", txBuilder.RandomETH())
	fmt.Println("SOL :", txBuilder.RandomSOL())
	fmt.Println("DOGE:", txBuilder.RandomDOGE())
	fmt.Println("LTC :", txBuilder.RandomLTC())
	fmt.Println("RVN :", txBuilder.RandomRVN())
	fmt.Println("TRX :", txBuilder.RandomTRX())

	t.Log(ethWal.ValidAddress(txBuilder.RandomETH()))
	t.Log(tronWal.ValidAddress(txBuilder.RandomTRX()))
}
