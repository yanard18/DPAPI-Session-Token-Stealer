package cookiemonster

import (
	"flag"
	"log"
	"os"
)

type Args struct {
	StateFile    string
	CookiesFile  string
	Output       string
	Auto         bool
	KillBrowsers bool
}

func ParseArgs() (Args, error) {
	var args Args
	flag.StringVar(&args.StateFile, "state", "", "path to the state file")
	flag.StringVar(&args.CookiesFile, "cookies", "", "path to the cookies file")
	flag.StringVar(&args.Output, "output", "", "output file")
	flag.BoolVar(&args.Auto, "auto", false, "Scan for cookies automatically")
	flag.BoolVar(&args.KillBrowsers, "kill", false, "Kill browser processes, cookies of running instance of browsers cannot be decrypted")
	flag.Parse()

	if args.Auto {
		return args, nil
	}

	if _, err := os.Stat(args.StateFile); os.IsNotExist(err) {
		log.Fatalf("State file does not exist: %v", err)
	}

	if _, err := os.Stat((args.CookiesFile)); os.IsNotExist(err) {
		log.Fatalf("Cookies file does not exist: %v", err)
	}

	return args, nil
}
