package cookiemonster

import (
	"encoding/binary"
	"fmt"
)

type DPAPIBlob struct {
	Version             uint32
	GUID                [16]byte
	EncryptionAlgortihm uint32
	EncryptionKeySize   uint32
}

func ParseDPAPIBlob(blob []byte) (*DPAPIBlob, error) {
	if len(blob) < 20 {
		return nil, fmt.Errorf("blob is too short")
	}

	d := &DPAPIBlob{}
	d.Version = binary.LittleEndian.Uint32(blob[:4])
	copy(d.GUID[:], blob[4:20])
	d.EncryptionAlgortihm = binary.LittleEndian.Uint32(blob[20:24])
	d.EncryptionKeySize = binary.LittleEndian.Uint32(blob[24:28])

	return d, nil
}

// ExtractGUID returns the GUID as a string
func ExtractGUID(guid [16]byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		binary.LittleEndian.Uint32(guid[0:4]),
		binary.LittleEndian.Uint16(guid[4:6]),
		binary.LittleEndian.Uint16(guid[6:8]),
		binary.BigEndian.Uint16(guid[8:10]),
		guid[10:])
}
