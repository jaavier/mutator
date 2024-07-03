// Archivo: mutator_duplicate_test.go

package mutator

import (
	"testing"
	"unicode"
)

// Helper function to count the number of letters in the result
func countLetters(runes []rune) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range runes {
		if unicode.IsLetter(r) {
			counts[r]++
		}
	}
	return counts
}

// Helper function to count the number of duplicated letters
func countDuplications(original []rune, changed []rune) int {
	origCounts := countLetters(original)
	changedCounts := countLetters(changed)

	duplicationCount := 0
	for r, origCount := range origCounts {
		if changedCount, found := changedCounts[r]; found {
			if changedCount > origCount {
				duplicationCount += changedCount - origCount
			}
		}
	}

	return duplicationCount
}

// Test for duplicateChars
func TestDuplicateChars(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewDuplicateMutation(1, 1.0),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int // Expected number of duplications
	}{
		{
			name:          "Single duplicate",
			input:         "hello",
			mutationType:  NewDuplicateMutation(1, 1.0),
			expectedCount: 1,
		},
		{
			name:          "Multiple duplicates",
			input:         "hello",
			mutationType:  NewDuplicateMutation(2, 1.0),
			expectedCount: 2,
		},
		{
			name:          "More duplicates than characters",
			input:         "hello",
			mutationType:  NewDuplicateMutation(5, 1.0),
			expectedCount: 5, // This might not be accurate due to randomness
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.duplicateChars(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of duplications
			duplicationCount := countDuplications([]rune(original), []rune(resultStr))

			if duplicationCount != tt.expectedCount {
				t.Errorf("duplicateChars() duplicated %d letters, want %d", duplicationCount, tt.expectedCount)
			}
		})
	}
}
