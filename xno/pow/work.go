package pow

import (
	"encoding/binary"
	"golang.org/x/crypto/blake2b"
	"hash"
	"math/rand"
	"runtime"
)

// GenerateCPU generates proof-of-work using the CPU.
func GenerateCPU(data []byte, target uint64) (work []byte, err error) {
	n := runtime.NumCPU()
	ch := make(chan []byte, n)
	hash := make([]hash.Hash, n)
	for i := 0; i < n; i++ {
		if hash[i], err = blake2b.New(8, nil); err != nil {
			return
		}
	}
	done := false
	x := rand.Uint64()
	for i := 0; i < n; i++ {
		go func(i int) {
			work := make([]byte, 8)
			for x := x + uint64(i); !done; x += uint64(n) {
				binary.BigEndian.PutUint64(work, x)
				hash[i].Reset()
				hash[i].Write(work)
				hash[i].Write(data)
				if binary.LittleEndian.Uint64(hash[i].Sum(nil)) >= target {
					done = true
					ch <- work
				}
			}
		}(i)
	}
	return <-ch, nil
}
