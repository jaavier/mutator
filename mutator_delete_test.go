package mutator

import (
	"testing"
)

// Helper function to count the number of deletions
func countDeletions(original, changed string) int {
	// Count the number of characters removed
	return len(original) - len(changed)
}

// Test for deleteChars
func TestDeleteChars(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewDeleteMutation(1, 1.0, "all"),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int // Expected number of deletions
	}{
		{
			name:          "Single Deletion - All",
			input:         "hello",
			mutationType:  NewDeleteMutation(1, 1.0, "all"),
			expectedCount: 1,
		},
		{
			name:          "Multiple Deletions - All",
			input:         "hello",
			mutationType:  NewDeleteMutation(2, 1.0, "all"),
			expectedCount: 2,
		},
		{
			name:          "More Deletions Than Characters - All",
			input:         "hello",
			mutationType:  NewDeleteMutation(5, 1.0, "all"),
			expectedCount: 5, // In this case, all characters might be deleted
		},
		{
			name:          "Delete Letters Only",
			input:         "hello123!",
			mutationType:  NewDeleteMutation(2, 1.0, "letters"),
			expectedCount: 2,
		},
		{
			name:          "Delete Numbers Only",
			input:         "hello123!",
			mutationType:  NewDeleteMutation(1, 1.0, "numbers"),
			expectedCount: 1,
		},
		{
			name:          "Delete Special Characters Only",
			input:         "hello123!",
			mutationType:  NewDeleteMutation(1, 1.0, "special"),
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.deleteChars(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of deletions
			deletionCount := countDeletions(original, resultStr)

			if deletionCount != tt.expectedCount {
				t.Errorf("deleteChars() deleted %d characters, want %d", deletionCount, tt.expectedCount)
			}
		})
	}
}
