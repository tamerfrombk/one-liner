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

func ParseArgs() *Args {
	helpPtr := flag.Bool("h", false, "displays this help message")

	flag.Parse()

	return &Args {
		IsHelp: *helpPtr,
	}
}

func Clean(b []byte) string {
	for i, r := range b {
		if r == byte('\n') {
			b[i] = byte(' ')
		}
	}

	return string(b)
}

func PrintOneLine(r io.Reader) error {

	reader := bufio.NewReader(r)

	buf := make([]byte, 64 * 1024)
	for b, err := reader.Read(buf); b > 0; b, err = reader.Read(buf) {
		if err != nil {
			return err
		}
		
		fmt.Print(Clean(buf))
	}

	return nil
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