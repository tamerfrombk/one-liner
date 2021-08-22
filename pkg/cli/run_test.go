package cli

import "testing"

func TestClean(t *testing.T) {
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
		if actual := Clean([]byte(input)); actual != expected {
			t.Fatalf("expected '%s', got '%s'", expected, actual)
		}
	}
}