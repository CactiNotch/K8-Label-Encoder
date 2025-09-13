package encoder

import "fmt"

// ValidateLabel checks if a label meets Kubernetes label requirements.
// Validates:
//   - Maximum length of 63 characters
// Returns an error if validation fails, nil otherwise.
func ValidateLabel(label string) error {
	if len(label) > 63 {
		return fmt.Errorf("label length exceeds kubernetes maximum of 63 characters")
	}
	return nil
}
