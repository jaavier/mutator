package mutator

import (
	"strings"
	"testing"
)

// Helper function to count the number of insertions in the result string
func countInsertions(original, changed string) int {
	return len(changed) - len(original)
}

// Test for insertChars
func TestInsertChars(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewInsertMutation(1, 1.0, "all"),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int // Expected number of insertions
	}{
		{
			name:          "Single Insertion - All",
			input:         "hello",
			mutationType:  NewInsertMutation(1, 1.0, "all"),
			expectedCount: 1,
		},
		{
			name:          "Multiple Insertions - All",
			input:         "hello",
			mutationType:  NewInsertMutation(2, 1.0, "all"),
			expectedCount: 2,
		},
		{
			name:          "More Insertions Than Characters - All",
			input:         "hello",
			mutationType:  NewInsertMutation(5, 1.0, "all"),
			expectedCount: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.insertChars(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of insertions
			insertionCount := countInsertions(original, resultStr)

			if insertionCount != tt.expectedCount {
				t.Errorf("insertChars() inserted %d characters, want %d", insertionCount, tt.expectedCount)
			}
		})
	}
}

// Test for insertBoundaryValues
func TestInsertBoundaryValues(t *testing.T) {
	mutator := New(&Config{})

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Insert Boundary Value",
			input:    "hello",
			expected: "hello0", // Or any other boundary value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			result := mutator.insertBoundaryValues(inputRunes)
			resultStr := string(result)

			// Check if the result contains the expected boundary value
			if !containsBoundaryValue(resultStr) {
				t.Errorf("insertBoundaryValues() did not insert a boundary value correctly: %s got %s", tt.input, resultStr)
			}
		})
	}
}

// Test for insertHTML
func TestInsertHTML(t *testing.T) {
	mutator := New(&Config{})

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Insert HTML Tag",
			input:    "hello",
			expected: "hello<div>", // Or any other HTML tag
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			result := mutator.insertHTML(inputRunes, MutationType{
				Type:   HtmlElement,
				Length: 1,
			})
			resultStr := string(result)

			// Check if the result contains an HTML tag
			if !containsHTMLTag(resultStr) {
				t.Errorf("insertHTML() did not insert an HTML tag correctly")
			}
		})
	}
}

// Test for insertFakeWord
func TestInsertFakeWord(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewInsertMutation(1, 1.0, "fakeword"),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int
	}{
		{
			name:          "Insert Fake Word",
			input:         "hello",
			mutationType:  NewInsertMutation(1, 1.0, "fakeword"),
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.insertFakeWord(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of insertions
			insertionCount := countInsertions(original, resultStr)

			if insertionCount != tt.expectedCount {
				t.Errorf("insertFakeWord() inserted %d characters, want %d", insertionCount, tt.expectedCount)
			}
		})
	}
}

// Test for insertRandomJSON
func TestInsertRandomJSON(t *testing.T) {
	mutator := New(&Config{})

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Insert Random JSON",
			input:    "hello",
			expected: "hello{\"key\":\"value\"}", // Expected JSON string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			result := mutator.insertRandomJSON(inputRunes, 1)
			resultStr := string(result)

			// Check if the result contains JSON
			if !strings.Contains(resultStr, "{") || !strings.Contains(resultStr, "") {
				t.Errorf("insertRandomJSON() did not insert JSON correctly: %s", resultStr)
			}
		})
	}
}

// Test for insertRandomWord
func TestInsertRandomWord(t *testing.T) {
	mutator := New(&Config{
		MutationTypes: []MutationType{
			NewInsertMutation(1, 1.0, "randomword"),
		},
	})

	tests := []struct {
		name          string
		input         string
		mutationType  MutationType
		expectedCount int
	}{
		{
			name:          "Insert Random Word",
			input:         "hello",
			mutationType:  NewFakeWordMutation(1, 1.0),
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputRunes := []rune(tt.input)
			original := string(inputRunes)
			result := mutator.insertFakeWord(inputRunes, tt.mutationType)
			resultStr := string(result)
			// Count the number of insertions
			insertionCount := countInsertions(original, resultStr)
			if insertionCount != tt.expectedCount {
				t.Errorf("insertRandomWord() inserted %d characters, want %d", insertionCount, tt.expectedCount)
			}
		})
	}
}

// Helper functions to check if certain values are present in the result
func containsBoundaryValue(result string) bool {
	boundaryValues := []string{
		"0", "1", "-1", "32767", "-32768", "2147483647", "-2147483648", "9223372036854775807", "-9223372036854775808",
	}
	for _, value := range boundaryValues {
		if strings.Contains(result, value) {
			return true
		}
	}
	return false
}

func containsHTMLTag(result string) bool {
	htmlTags := []string{
		"<div>", "</div>", "<span>", "</span>", "<a href='#'>", "</a>",
		"<b>", "</b>", "<i>", "</i>", "<p>", "</p>",
		"<br>", "<img src='#'/>", "<h1>", "</h1>",
	}
	for _, tag := range htmlTags {
		if strings.Contains(result, tag) {
			return true
		}
	}
	return false
}
