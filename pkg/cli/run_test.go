package cli

import (
	"testing"
	"bytes"
	"strings"
)

func TestPrintOneLine_LFText_LF(t *testing.T) {
    testCases := map[string]string {
		" hello ": " hello ",
		"\nhello": " hello",
		"": "",
		"foo": "foo",
		"f": "f",
		"\n": " ",
		"hello\n ": "hello  ",
	}

	runTestCases(t, testCases, "\n")
}

func TestPrintOneLine_CRLFText_CRLF(t *testing.T) {
    testCases := map[string]string {
		" hello ": " hello ",
		"\r\nhello": " hello",
		"": "",
		"foo": "foo",
		"f": "f",
		"\r\n": " ",
		"hello\r\n ": "hello  ",
	}

	runTestCases(t, testCases, "\r\n")
}

func TestPrintOneLine_CRLFText_LF(t *testing.T) {
    testCases := map[string]string {
		" hello ": " hello ",
		"\r\nhello": "\r hello",
		"": "",
		"foo": "foo",
		"f": "f",
		"\r\n": "\r ",
		"hello\r\n ": "hello\r  ",
	}

	runTestCases(t, testCases, "\n")
}

func TestPrintOneLine_LFText_CRLF(t *testing.T) {
    testCases := map[string]string {
		" hello ": " hello ",
		"\nhello": "\nhello",
		"": "",
		"foo": "foo",
		"f": "f",
		"\n": "\n",
		"hello\n ": "hello\n ",
	}

	runTestCases(t, testCases, "\r\n")
}

func runTestCases(t *testing.T, testCases map[string]string, newLine string) {
	for input, expected := range testCases {
		reader := strings.NewReader(input)
		writer := bytes.NewBufferString("")

		if err := PrintOneLine(reader, writer, newLine); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected + newLine {
			t.Fatalf("expected '%s', got '%s'", expected + newLine, actual)
		}
	}
}