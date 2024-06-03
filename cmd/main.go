package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yanard18/cookiemonster"
)

func main() {
	args, err := cookiemonster.ParseArgs()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	encryptedKey, err := cookiemonster.ParseLocalState(args.StateFile)
	if err != nil {
		log.Fatalf("Error parsing local state file: %v", err)
	}

	key, err := cookiemonster.DecryptDPAPI(encryptedKey)
	if err != nil {
		log.Fatalf("Error decrypting DPAPI blob of local state encrypted key: %v", err)
	}

	if args.CookiesFile != "" {
		cookies, err := cookiemonster.ParseCookies(args.CookiesFile)
		if err != nil {
			fmt.Printf("[-] %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[+] Found %d cookies\n", len(cookies))

		for _, cookie := range cookies {

			// Decrypt the cookie value
			decrypted, err := cookiemonster.DecryptEncryptedCookieValue([]byte(cookie.Value), key)
			if err != nil {
				fmt.Printf("[-] %v\n", err)
				continue
			}

			fmt.Printf("[+] Host: %s\n", cookie.Host)
			fmt.Printf("    Name: %s\n", cookie.Name)
			fmt.Printf("    Value: %s\n", decrypted)
			fmt.Printf("    Path: %s\n", cookie.Path)
			fmt.Printf("    IsSecure: %t\n", cookie.IsSecure)
			fmt.Printf("    IsHttpOnly: %t\n", cookie.IsHttpOnly)
			fmt.Printf("    CreationUtc: %d\n", cookie.CreationUtc)
			fmt.Printf("    ExpiryUtc: %d\n", cookie.ExpiryUtc)
			fmt.Println("--------------------------------------------------")
		}
	}

}
