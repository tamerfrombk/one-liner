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

func init() {
	flag.Usage = func () {
		fmt.Fprintf(flag.CommandLine.Output(), "one-liner is a program for collapsing input from stdin into a single line on stdout:\n")
		flag.PrintDefaults()
	}
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

	readBuffer := make([]byte, READ_BUFFER_LEN)
	for n, err := reader.Read(readBuffer); n > 0; n, err = reader.Read(readBuffer) {
		if err != nil {
			return err
		}

		// OPTIMIZATION(time): Write(buf) will write the entire contents of buf.
		// This means that if we manage to read as many bytes as the buffer holds,
		// then we can write the buffer directly.
		// If we read less than the read buffer size, we're forced to allocate a smaller
		// array.
		var writeBuffer []byte = nil;
		if n == READ_BUFFER_LEN {
			writeBuffer = removeNewline(readBuffer, n)
		} else {
			writeBuffer = removeNewline(readBuffer[:n], n)
		}
		w.Write(writeBuffer)
	}

	return nil
}

func Run(cmdLine []string) int {
	args := ParseArgs()
	if args.IsHelp {
		flag.Usage()
		return 0
	}

	if err := PrintOneLine(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return 1
	}

	return 0
}

func removeNewline(buf []byte, n int) []byte {
	for i := 0; i < n; i++ {
		if buf[i] == '\n' {
			buf[i] = ' '
		}
	}

	return buf
}
