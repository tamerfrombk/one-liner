package cli

import (
	"testing"
	"bytes"
	"strings"
)

func TestPrintOneLine(t *testing.T) {
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

		if err := PrintOneLine(reader, writer); err != nil {
			t.Fatal(err)
		}

		if actual := writer.String(); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
}