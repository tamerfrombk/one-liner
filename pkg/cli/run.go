package cli

import ( 
	"bufio"
	"io"
	"flag"
	"fmt"
	"os"
)

type Args struct {
	IsHelp bool
}

func ParseArgs() (*Args) {
	helpPtr := flag.Bool("h", false, "displays this help message")

	flag.Parse()

	return &Args {
		IsHelp: *helpPtr,
	}
}

func Clean(s string) string {
	ret := ""
	for _, c := range(s) {
		if c != '\n' {
			ret += string(c)
		}
	}

	return ret
}

func PrintOneLine(r io.Reader) error {
	scanner := bufio.NewScanner(r); 
	for scanner.Scan() {
		fmt.Print(Clean(scanner.Text()))
	}

	return scanner.Err()
}

func Run(cmdLine []string) int {
	args := ParseArgs()
	if args.IsHelp {
		printHelp(0)
	}

	if err := PrintOneLine(os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return 1
	}

	return 0
}

func printHelp(exitCode int) {
	flag.PrintDefaults()

	os.Exit(exitCode)
}