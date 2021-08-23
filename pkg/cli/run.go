package cli

import ( 
	"bufio"
	"errors"
	"io"
	"flag"
	"fmt"
	"os"
	"runtime"
)

type Args struct {
	IsHelp bool
	LineEnding string
}

func init() {
	flag.Usage = func () {
		fmt.Fprintf(flag.CommandLine.Output(), "one-liner is a program for collapsing input from stdin into a single line on stdout:\n")
		flag.PrintDefaults()
	}
}

func ParseArgs() (*Args, error) {
	helpPtr := flag.Bool("h", false, "displays this help message")
	lfPtr := flag.Bool("lf", false, "force Linux style \\n input newline")
	crlfPtr := flag.Bool("crlf", false, "force Windows style \\r\\n input newline")

	flag.Parse()

	if *lfPtr && *crlfPtr {
		return nil, errors.New("line endings cannot be both \\n and \\r\\n")
	} 
		
	var lineEnding string;
	if *lfPtr != *crlfPtr {
		if *lfPtr {
			lineEnding = "\n"
		} else {
			lineEnding = "\r\n"
		}
	} else {
		// both are false, default to system newline
		lineEnding = determineSystemNewline()
	}
	

	return &Args {
		IsHelp: *helpPtr,
		LineEnding: lineEnding,
	}, nil
}

func PrintOneLine(r io.Reader, w io.Writer, lineEnding string) error {
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
			writeBuffer = removeNewline(readBuffer, n, lineEnding)
		} else {
			writeBuffer = removeNewline(readBuffer[:n], n, lineEnding)
		}
		w.Write(writeBuffer)
	}
	w.Write([]byte(lineEnding))

	return nil
}

func Run(cmdLine []string) int {
	args, err := ParseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return 1
	}

	if args.IsHelp {
		flag.Usage()
		return 0
	}

	if err := PrintOneLine(os.Stdin, os.Stdout, args.LineEnding); err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return 1
	}

	return 0
}

func determineSystemNewline() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n";
}

func removeNewline(buf []byte, n int, newLine string) []byte {
	if newLine == "\r\n" {
		return removeCRLF(buf, n)
	} else {
		return removeLF(buf, n)
	}
}

func removeCRLF(buf []byte, n int) []byte {
	out := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if (buf[i] == '\r') {
			if next := i + 1; next < n && buf[next] == '\n' {
				out = append(out, ' ')
				i++; // skip the '\n'
			} else {
				out = append(out, '\r')
			}
		} else {
			out = append(out, buf[i])
		}
	}

	return out
}

func removeLF(buf []byte, n int) []byte {
	for i := 0; i < n; i++ {
		if buf[i] == '\n' {
			buf[i] = ' '
		}
	}

	return buf
}
