package encoder

// SpecialCharMap provides bidirectional mapping between special characters
// and their encoded representations for Kubernetes labels.
type SpecialCharMap struct {
    CharToCode map[string]string // Maps special characters to their encoded form
    CodeToChar map[string]string // Maps encoded form back to special characters
}

// Encodings represents a single character-to-code mapping pair.
type Encodings struct {
    Char, // Original special character
    Code string // Encoded representation
}


// NewSpecialCharMap maps special characters to their encoded representations
// Follows the url encoding pattern of _XX where XX is the hex value of the character
// Replacing % with _ to make it more Kubernetes label friendly
// Do not need to encode alphanumeric characters, "-", "_", and "." as they are safe for labels
func NewSpecialCharMap() *SpecialCharMap {
	charToCode := make(map[string]string)
	codeToChar := make(map[string]string)

	for _, mapping := range []Encodings {
		{" ", "_20"},
		{"!", "_21"},
		{"\"", "_22"},
		{"#", "_23"},
		{"$", "_24"},
		{"%", "_25"},
		{"&", "_26"},
		{"'", "_27"},
		{"(", "_28"},
		{")", "_29"},
		{"*", "_2A"},
		{"+", "_2B"},
		{",", "_2C"},
		{"/", "_2F"},
		{":", "_3A"},
		{";", "_3B"},
		{"<", "_3C"},
		{"=", "_3D"},
		{">", "_3E"},
		{"?", "_3F"},
		{"@", "_40"},
		{"[", "_5B"},
		{"\\", "_5C"},
		{"]", "_5D"},
		{"^", "_5E"},
		{"`", "_60"},
		{"{", "_7B"},
		{"|", "_7C"},
		{"}", "_7D"},
		{"~", "_7E"},
	} {
		charToCode[mapping.Char] = mapping.Code
		codeToChar[mapping.Code] = mapping.Char
	}

	return &SpecialCharMap{
        CharToCode: charToCode,
        CodeToChar: codeToChar,
    }
}