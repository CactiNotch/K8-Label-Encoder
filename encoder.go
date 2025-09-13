package encoder

import (
	"strings"
)

var specialChars = NewSpecialCharMap()

// EncodeSpecialChars converts special characters in a label to their encoded form.
// It preserves alphanumeric characters, "-", "_", and "." as they are valid in
// Kubernetes labels. Returns the encoded string.
func EncodeSpecialChars(label string) string {
	var builder strings.Builder
	builder.Grow(len(label) * 3)

	for _, ch := range label {
		if code, ok := specialChars.CharToCode[string(ch)]; ok {
			builder.WriteString(code)
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}


// DecodeSpecialChars converts encoded special characters back to their original form.
// It processes encoded sequences starting with '_' followed by two characters,
// leaving all other characters unchanged. Returns the decoded string.
func DecodeSpecialChars(encodedLabel string) string {
	var builder strings.Builder
	builder.Grow(len(encodedLabel))

	for i := 0; i < len(encodedLabel); i++ {
		if encodedLabel[i] == '_' && i+2 < len(encodedLabel) {
			code := encodedLabel[i : i+3]
			if char, ok := specialChars.CodeToChar[code]; ok {
				builder.WriteString(char)
				i += 2
				continue
			}
		}
		builder.WriteByte((encodedLabel[i]))
	}
	return builder.String()
}
