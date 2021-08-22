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
	
	const READ_BUFFER_LEN = 64 * 1024

	reader := bufio.NewReader(r)

	buf := make([]byte, READ_BUFFER_LEN)
	for n, err := reader.Read(buf); n > 0; n, err = reader.Read(buf) {
		if err != nil {
			return err
		}

		if n == READ_BUFFER_LEN {
			w.Write(clean(buf, n))
		} else {
			temp := buf[:n]
			w.Write(clean(temp, n))
		}
	}

	return nil
}

func Run(cmdLine []string) int {
	args := ParseArgs()
	if args.IsHelp {
		flag.PrintDefaults()
		return 0
	}

	if err := PrintOneLine(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return 1
	}

	return 0
}

func clean(buf []byte, n int) []byte {
	for i := 0; i < n; i++ {
		if buf[i] == '\n' {
			buf[i] = ' '
		}
	}

	return buf
}
