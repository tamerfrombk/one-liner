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

	for input, expected := range testCases {
		reader := strings.NewReader(input)
		writer := bytes.NewBufferString("")

		if err := PrintOneLine(reader, writer, "\n"); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
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

	for input, expected := range testCases {
		reader := strings.NewReader(input)
		writer := bytes.NewBufferString("")

		if err := PrintOneLine(reader, writer, "\r\n"); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
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

	for input, expected := range testCases {
		reader := strings.NewReader(input)
		writer := bytes.NewBufferString("")

		if err := PrintOneLine(reader, writer, "\n"); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
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

	for input, expected := range testCases {
		reader := strings.NewReader(input)
		writer := bytes.NewBufferString("")

		if err := PrintOneLine(reader, writer, "\r\n"); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
}