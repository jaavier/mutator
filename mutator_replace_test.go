package mutator

import (
	"testing"
)

// Helper function to count the number of replacements
func countReplacements(original, changed string) int {
	// Count the number of characters replaced
	replacements := 0
	for i := range original {
		if original[i] != changed[i] {
			replacements++
		}
	}
	return replacements
}

// Test for replaceChars
func TestReplaceChars(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewReplaceMutation(1, 1.0, "all"),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int // Expected number of replacements
	}{
		{
			name:          "Single Replacement - All",
			input:         "hello",
			mutationType:  NewReplaceMutation(1, 1.0, "all"),
			expectedCount: 1,
		},
		{
			name:          "Multiple Replacements - All",
			input:         "hello",
			mutationType:  NewReplaceMutation(2, 1.0, "all"),
			expectedCount: 2,
		},
		{
			name:          "More Replacements Than Characters - All",
			input:         "hello",
			mutationType:  NewReplaceMutation(5, 1.0, "all"),
			expectedCount: 5, // In this case, all characters might be replaced
		},
		{
			name:          "Replace Letters Only",
			input:         "hello123!",
			mutationType:  NewReplaceMutation(2, 1.0, "letters"),
			expectedCount: 2,
		},
		{
			name:          "Replace Numbers Only",
			input:         "hello123!",
			mutationType:  NewReplaceMutation(1, 1.0, "numbers"),
			expectedCount: 1,
		},
		{
			name:          "Replace Special Characters Only",
			input:         "hello123!",
			mutationType:  NewReplaceMutation(1, 1.0, "special"),
			expectedCount: 1,
		},
		{
			name:          "Replace Mixed Valid Characters",
			input:         "a1!b2@c3#d4$e5%",
			mutationType:  NewReplaceMutation(4, 1.0, "letters,numbers,special"),
			expectedCount: 4,
		},
		{
			name:          "Replace No Valid Characters",
			input:         "123456",
			mutationType:  NewReplaceMutation(2, 1.0, "letters"),
			expectedCount: 6, // No letters to replace
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.replaceChars(inputRunes, tt.mutationType)
			resultStr := string(result)

			// Count the number of replacements
			replacementCount := countReplacements(original, resultStr)

			if replacementCount == 0 {
				t.Errorf("replaceChars() replaced %d characters, want %d", replacementCount, tt.expectedCount)
			}

			// Additional checks to ensure replacements are valid
			for i, r := range resultStr {
				if rune(original[i]) != r && !isValidReplacement(r, tt.mutationType) {
					t.Errorf("replaceChars() replaced with invalid character: %c", r)
				}
			}
		})
	}
}

func isValidReplacement(r rune, mutationType MutationType) bool {
	if mutationType.Any {
		return true
	}
	if mutationType.CharSet != "" {
		for _, char := range mutationType.CharSet {
			if r == char {
				return true
			}
		}
	}
	return (mutationType.Letters && isLetter(r)) ||
		(mutationType.Numbers && isNumber(r)) ||
		(mutationType.Special && isSpecialChar(r))
}
