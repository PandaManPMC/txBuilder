package xno_test

import (
	"github.com/PandaManPMC/txBuilder/xno"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNanoAmount(t *testing.T) {
	for _, s := range []string{
		"133246497.546603",
		"1000.000000",
		"0.100000",
		"0.000001",
	} {
		n, err := xno.NanoAmountFromString(s)
		require.Nil(t, err)
		assert.Equal(t, s, n.String())
	}
	for _, s := range []string{
		"0.0000000000000000000000000000001",
	} {
		_, err := xno.NanoAmountFromString(s)
		require.NotNil(t, err)
	}
}
