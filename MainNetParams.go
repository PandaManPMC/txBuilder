package txBuilder

import (
	"fmt"
	"github.com/PandaManPMC/txBuilder/dogeNetParams"
	"github.com/PandaManPMC/txBuilder/ltcNetParams"
	"github.com/PandaManPMC/txBuilder/rvnNetParams"
	"github.com/btcsuite/btcd/chaincfg"
)

func init() {
	if e := chaincfg.Register(&ltcNetParams.MainNetParams); nil != e {
		fmt.Println(e)
	}
	if e := chaincfg.Register(&dogeNetParams.MainNetParams); nil != e {
		fmt.Println(e)
	}
	if e := chaincfg.Register(&rvnNetParams.MainNetParams); nil != e {
		fmt.Println(e)
	}
}
