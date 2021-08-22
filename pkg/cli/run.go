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

func PrintOneLine(r io.Reader, w io.Writer) error {

	reader := bufio.NewReader(r)

	buf := make([]byte, 64 * 1024)
	for b, err := reader.Read(buf); b > 0; b, err = reader.Read(buf) {
		if err != nil {
			return err
		}

		w.Write(clean(buf, b))
	}

	return nil
}

func Run(cmdLine []string) int {
	args := ParseArgs()
	if args.IsHelp {
		printHelp(0)
	}

	if err := PrintOneLine(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return 1
	}

	return 0
}

func clean(b []byte, n int) []byte {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if b[i] == '\n' {
			buf[i] = ' '
		} else {
			buf[i] = b[i] 
		}
	}

	return buf
}

func printHelp(exitCode int) {
	flag.PrintDefaults()

	os.Exit(exitCode)
}