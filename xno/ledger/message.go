package ledger

import (
	"encoding/binary"
)

func appendPath(buf []byte, path []uint32) []byte {
	var p [4]byte
	buf = append(buf, byte(len(path)))
	for _, i := range path {
		binary.BigEndian.PutUint32(p[:], 0x80000000|i)
		buf = append(buf, p[:]...)
	}
	return buf
}

// GetAddress returns the public key and the encoded address for the given BIP32 path.
func GetAddress(path []uint32) (pubkey []byte, address string, err error) {
	d, err := getDevice()
	if err != nil {
		return
	}
	defer d.Close()
	payload := []byte{0xa1, 0x02, 0x00, 0x00, 0x00}
	payload = appendPath(payload, path)
	payload[4] = byte(len(payload)) - 5
	resp, err := send(d, payload)
	if err != nil {
		return
	}
	pubkey = resp[:32]
	address = string(resp[33 : 33+resp[32]])
	return
}
