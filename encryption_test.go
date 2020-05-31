package encryption

import (
	"testing"
)

var testData = map[string]string{
	"hello":     "h2ll4",
	"hi there!": "h3 th2r2!",
	"the quick brown fox jumps over a lazy dog.": "th2 q53ck br4wn f4x j5mps 4v2r 1 l1zy d4g.",
}

func TestEncode(t *testing.T) {
	for phrase, expected := range testData {
		result := Encode(phrase)
		if result != expected {
			t.Errorf("test failed Encode - results not match\nGot:\n%v\nExpected:\n%v", result, expected)
			break
		}
	}
}

func TestDecode(t *testing.T) {
	for expected, phrase := range testData {
		result := Decode(phrase)
		if result != expected {
			t.Errorf("test failed Decode - results not match\nGot:\n%v\nExpected:\n%v", result, expected)
			break
		}
	}
}
