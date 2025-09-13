package encoder

import (
	"strings"
	"testing"
)

// TestEncodeDecodeSpecialChars verifies that special characters are correctly
// encoded and decoded, maintaining roundtrip consistency.
func TestEncodeDecodeSpecialChars(t *testing.T) {
	testCases := []struct {
		input   string
		encoded string
		decoded string
	}{
		{"hello world!", "hello_20world_21", "hello world!"},
		{"special_chars: _%&*", "special_chars_3A_20__25_26_2A", "special_chars: _%&*"},
		{"noSpecialChars", "noSpecialChars", "noSpecialChars"},
		{"mix_of-special.chars/and:more?", "mix_of-special.chars_2Fand_3Amore_3F", "mix_of-special.chars/and:more?"},
		{"label_with-safe-special.chars", "label_with-safe-special.chars", "label_with-safe-special.chars"},
	}
	for _, tc := range testCases {
		encoded := EncodeSpecialChars(tc.input)
		if encoded != tc.encoded {
			t.Errorf("EncodeSpecialChars(%q) = %q; want %q", tc.input, encoded, tc.encoded)
		}
		decoded := DecodeSpecialChars(encoded)
		if decoded != tc.decoded {
			t.Errorf("DecodeSpecialChars(%q) = %q; want %q", encoded, decoded, tc.decoded)
		}
	}
}

// TestValidateLabel verifies that label validation correctly identifies
// labels that exceed Kubernetes length restrictions.
func TestValidateLabel(t *testing.T) {
	longLabel := strings.Repeat("a", 64)
	if err := ValidateLabel(longLabel); err == nil {
		t.Error("Expected error for long label, got nil")
	}
}

// BenchmarkEncodeSpecialChars measures the performance of special character encoding.
func BenchmarkEncodeSpecialChars(b *testing.B) {
	input := "special_chars: _%&*"
	for b.Loop() {
		EncodeSpecialChars(input)
	}
}
