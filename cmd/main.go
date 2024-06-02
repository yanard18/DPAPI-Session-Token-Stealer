package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yanard18/cookiemonster"
)

type Args struct {
	StateFile   string
	CookiesFile string
}

func parseArgs() Args {
	var args Args
	flag.StringVar(&args.StateFile, "state", "", "path to the state file")
	flag.StringVar(&args.CookiesFile, "cookies", "", "path to the cookies file")
	flag.Parse()
	return args
}

func main() {
	args := parseArgs()

	if _, err := os.Stat(args.StateFile); os.IsNotExist(err) {
		fmt.Println("State file does not exist")
		os.Exit(1)
		return
	}

	if _, err := os.Stat((args.CookiesFile)); os.IsNotExist(err) {
		fmt.Println("Cookies file does not exist")
		os.Exit(1)
		return
	}

	encryptedKey, err := cookiemonster.ParseLocalState(args.StateFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	key, err := cookiemonster.DecryptDPAPI(encryptedKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
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
