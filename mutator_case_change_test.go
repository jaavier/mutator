// Archivo: mutator_change_case_test.go
package mutator

import (
	"fmt"
	"testing"
	"unicode"
)

// Function to count the number of changes in the case of letters
func countCaseChanges(original []rune, changed []rune) int {
	count := 0
	for i := range original {
		if unicode.IsLetter(original[i]) && original[i] != changed[i] && unicode.IsLetter(changed[i]) {
			if (unicode.IsUpper(original[i]) && unicode.IsLower(changed[i])) || (unicode.IsLower(original[i]) && unicode.IsUpper(changed[i])) {
				count++
			}
		}
	}
	return count
}

// Test for changeCaseChars
func TestChangeCaseChars(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewCaseChangeMutation(1, 1.0),
		},
	})

	tests := []struct {
		name          MutationAction
		input         string
		mutationType  MutationType
		expectedCount int // Expected number of case changes
	}{
		{
			name:          CaseChange,
			input:         "hello",
			mutationType:  NewCaseChangeMutation(1, 1.0),
			expectedCount: 1,
		},
		{
			name:          CaseChange,
			input:         "hello",
			mutationType:  NewCaseChangeMutation(2, 1.0),
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%x", tt.name), func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.changeCaseChars(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of changes
			changeCount := countCaseChanges([]rune(original), []rune(resultStr))

			if changeCount != tt.expectedCount {
				t.Errorf("changeCaseChars() changed %d characters, want %d", changeCount, tt.expectedCount)
			}
		})
	}
}
