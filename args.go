package cookiemonster

import (
	"flag"
	"log"
	"os"
)

type Args struct {
	StateFile   string
	CookiesFile string
}

func ParseArgs() (Args, error) {
	var args Args
	flag.StringVar(&args.StateFile, "state", "", "path to the state file")
	flag.StringVar(&args.CookiesFile, "cookies", "", "path to the cookies file")
	flag.Parse()

	if _, err := os.Stat(args.StateFile); os.IsNotExist(err) {
		log.Fatalf("State file does not exist: %v", err)
	}

	if _, err := os.Stat((args.CookiesFile)); os.IsNotExist(err) {
		log.Fatalf("Cookies file does not exist: %v", err)
	}

	return args, nil
}
