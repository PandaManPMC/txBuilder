package txBuilder

import (
	"github.com/PandaManPMC/txBuilder/btcWal"
	"github.com/PandaManPMC/txBuilder/ethWal"
	"github.com/PandaManPMC/txBuilder/sol"
	"github.com/PandaManPMC/txBuilder/tronWal"
	"github.com/PandaManPMC/txBuilder/xno"
	"testing"
)

func TestAddressValidate(t *testing.T) {
	t.Log("ValidateAddress BTC")
	t.Log(btcWal.IsValidBTCAddress("bc1q54zg0nqjh2sqkv9yk93682w9jtcnp53gg9myyz9rz3uaez5s4dasczkjtm"))
	t.Log(btcWal.IsValidBTCAddress("bc1q5c7qck7up05hwjau7shqlvkhd46xeqfav9c93x"))
	t.Log()

	t.Log("ValidateAddress LTC")
	t.Log(btcWal.IsValidLTCAddress("ltc1q9xsxup3px68gppudjtfr67cmjl0uat4fpj2hwa"))
	t.Log(btcWal.IsValidLTCAddress("ltc1qxp5fpy3km65qan9r2kv93rsvpnt7he0ewgnc9j"))
	t.Log()

	t.Log("ValidateAddress DOGE")
	t.Log(btcWal.IsValidDOGEAddress("DP1JGhd4e2Dsiaoq67Zoq61wJkp2E38S9P"))
	t.Log(btcWal.IsValidDOGEAddress("D8RQc9X9i98B3pbhG3hbAaT6vFXX3HoYGU"))
	t.Log()

	t.Log("ValidateAddress RVN")
	t.Log(btcWal.IsValidRVNAddress("RGqNpd2Jc26szpRdKaSFziEw3mqmAXitKn"))
	t.Log(btcWal.IsValidRVNAddress("RPxbnu3kt8ZcN5HLH7ACVUz8Eh1uPPZyaK"))
	t.Log()

	t.Log("ValidateAddress EVN")
	t.Log(ethWal.ValidAddress("0x6ef25ea3f4cceae27d57cbe9a4cfdd2ded2b2740"))
	t.Log(ethWal.ValidAddress("0x19647d4f0f3ea905f635850c8a6745282a6ae1e6"))
	t.Log(ethWal.ValidAddress("0x1e1c6ce778091208dc9057522c7049b255fa0a19"))
	t.Log()

	t.Log("ValidateAddress SOL")
	t.Log(sol.ValidSolanaAddress("ECMPJt8PUwwDVtBBhxAhogy7gkiwaNgQVEApmCx1ZFpM"))
	t.Log(sol.ValidSolanaAddress("3eDrHFaNk6dfWXgkPJM1v4ybjDQTFg4CZpCrmP75ywx1"))
	t.Log(sol.ValidSolanaAddress("GAqJNshjkWkDnjohbn73bpQiQm6oM2tbQ152MrH49EvK"))
	t.Log()

	t.Log("ValidateAddress NANO")
	t.Log(xno.IsXNOAddress("nano_38zggnkwefudity7igsw5jh1y7hogxbpm5ihwbuju6hgsboni35a7ppfzb1q"))
	t.Log(xno.IsXNOAddress("nano_1udrywwns1jj74yhjcabr3y4zqmez8j4h5rodyptcwbc39xf44hwzq5dgzpa"))
	t.Log()

	t.Log("ValidateAddress TRON")
	t.Log(tronWal.ValidAddress("THFR2uZQoPech7NTuJmkvkMmFo9fcABr2U"))
	t.Log(tronWal.ValidAddress("TXbvZxk1wcxZFkTQyVze4nXXdF9JDg5RUC"))
	t.Log(tronWal.ValidAddress("TT8CgE55UQou36LwVCVDwWTboHvX1EcFMm"))
	t.Log()

}

func TestAddressValidateFail(t *testing.T) {
	t.Log("ValidateAddress BTC")
	t.Log(btcWal.IsValidBTCAddress("bc1q54zg0nqjh2sqkv9yk93682w9jtcnp53gg9myyz9rz3uaez5s4dasczkjt1"))
	t.Log(btcWal.IsValidBTCAddress("bc1q54zg0nqjh2sqkv9yk93682w9jtcnp53gg9myyz9rz3uaez5s4dasczkjt"))
	t.Log()

	t.Log("ValidateAddress LTC")
	t.Log(btcWal.IsValidLTCAddress("ltc1q9xsxup3px68gppudjtfr67cmjl0uat4fpj2hw1"))
	t.Log(btcWal.IsValidLTCAddress("ltc1q9xsxup3px68gppudjtfr67cmjl0uat4fpj2hw"))
	t.Log()

	t.Log("ValidateAddress DOGE")
	t.Log(btcWal.IsValidDOGEAddress("DP1JGhd4e2Dsiaoq67Zoq61wJkp2E38S91"))
	t.Log(btcWal.IsValidDOGEAddress("DP1JGhd4e2Dsiaoq67Zoq61wJkp2E38S9"))
	t.Log()

	t.Log("ValidateAddress RVN")
	t.Log(btcWal.IsValidRVNAddress("RGqNpd2Jc26szpRdKaSFziEw3mqmAXitK1"))
	t.Log(btcWal.IsValidRVNAddress("RGqNpd2Jc26szpRdKaSFziEw3mqmAXitK"))
	t.Log()

	t.Log("ValidateAddress EVN")
	t.Log(ethWal.ValidAddress("0x19647d4f0f3ea905f635850c8a6745282a6ae1e"))
	t.Log()

	t.Log("ValidateAddress SOL")
	t.Log(sol.ValidSolanaAddress("ECMPJt8PUwwDVtBBhxAhogy7gkiwaNgQVEApmCx1ZFp1a"))
	t.Log(sol.ValidSolanaAddress("ECMPJt8PUwwDVtBBhxAhogy7gkiwaNgQVEApmCx1ZFp1"))
	t.Log(sol.ValidSolanaAddress("ECMPJt8PUwwDVtBBhxAhogy7gkiwaNgQVEApmCx1ZFp"))
	t.Log(sol.ValidSolanaAddress("ECMPJt8PUwwDVtBBhxAhogy7gkiwaNgQVEApmCx1ZF"))
	t.Log()

	t.Log("ValidateAddress NANO")
	t.Log(xno.IsXNOAddress("nano_38zggnkwefudity7igsw5jh1y7hogxbpm5ihwbuju6hgsboni35a7ppfzb10"))
	t.Log(xno.IsXNOAddress("nano_38zggnkwefudity7igsw5jh1y7hogxbpm5ihwbuju6hgsboni35a7ppfzb1"))
	t.Log()

	t.Log("ValidateAddress TRON")
	t.Log(tronWal.ValidAddress("THFR2uZQoPech7NTuJmkvkMmFo9fcABr2"))
	t.Log(tronWal.ValidAddress("THFR2uZQoPech7NTuJmkvkMmFo9fcABr21"))
	t.Log()
}
