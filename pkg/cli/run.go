package cli

import ( 
	"flag"
	"fmt"
	"os"
)

type Args struct {
	IsHelp bool
}

func ParseArgs(args []string) (*Args, error) {
	helpPtr := flag.Bool("h", false, "displays this help message")

	flag.Parse()

	return &Args {
		IsHelp: *helpPtr,
	}, nil
}

func Run(cmdLine []string) int {
	args, err := ParseArgs(cmdLine)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if args.IsHelp {
		printHelp(0)
	}

	return 0
}

func printHelp(exitCode int) {
	flag.PrintDefaults()
	os.Exit(exitCode)
}