package github.com/jaavier/mutator

import (
	"testing"
)

// Helper function to count the number of picked characters
func countPickedChars(original, picked string) int {
	// Picked characters should be equal to the length of the picked string
	return len(picked)
}

// Test for pickChars
func TestPickChars(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int             // Expected number of picked characters
		expectedChars func(rune) bool // Function to validate picked characters
	}{
		{
			name:          "Pick Letters Only",
			input:         "a1!b2@c3#d4$e5%",
			mutationType:  NewPickMutation(3, 1.0, "letters"),
			expectedCount: 3,
			expectedChars: isLetter,
		},
		{
			name:          "Pick Numbers Only",
			input:         "a1!b2@c3#d4$e5%",
			mutationType:  NewPickMutation(2, 1.0, "numbers"),
			expectedCount: 2,
			expectedChars: isNumber,
		},
		{
			name:          "Pick Special Characters Only",
			input:         "a1!b2@c3#d4$e5%",
			mutationType:  NewPickMutation(4, 1.0, "special"),
			expectedCount: 4,
			expectedChars: isSpecialChar,
		},
		{
			name:          "Pick Mixed Valid Characters",
			input:         "a1!b2@c3#d4$e5%",
			mutationType:  NewPickMutation(5, 1.0, "letters,numbers,special"),
			expectedCount: 5,
			expectedChars: func(r rune) bool {
				return isLetter(r) || isNumber(r) || isSpecialChar(r)
			},
		},
		{
			name:          "Pick More Than Available",
			input:         "abc",
			mutationType:  NewPickMutation(5, 1.0, "letters"),
			expectedCount: 3, // Since we only have 3 letters available
			expectedChars: isLetter,
		},
		{
			name:          "Pick No Valid Characters",
			input:         "123456!",
			mutationType:  NewPickMutation(2, 1.0, "special"),
			expectedCount: 1, // No letters to pick
			expectedChars: isSpecialChar,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize Mutator with a default configuration
			mutator := New(&Config{
				MutationTypes: []MutationType{tt.mutationType},
			})

			inputRunes := []rune(tt.input)
			result := mutator.pickChars(inputRunes, tt.mutationType)
			resultStr := string(result)

			// Count the number of picked characters
			pickedCount := countPickedChars(tt.input, resultStr)

			// fmt.Println(tt.expectedCount, pickedCount)

			if pickedCount != tt.expectedCount {
				t.Errorf("pickChars() picked %d characters, want %d", pickedCount, tt.expectedCount)
			}

			// Check if all characters in result are valid according to the mutationType
			for _, r := range resultStr {
				if !tt.expectedChars(r) {
					t.Errorf("pickChars() result contains invalid character: %c", r)
				}
			}
		})
	}
}
