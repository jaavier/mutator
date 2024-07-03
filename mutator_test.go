package github.com/jaavier/mutator

// import (
// 	"testing"
// )

// func TestSwapChars(t *testing.T) {
// 	m := &Mutator{}
// 	runes := []rune("abcdef")
// 	mutationType := MutationType{Length: 3}

// 	originalRunes := make([]rune, len(runes))
// 	copy(originalRunes, runes)

// 	result := m.swapChars(runes, mutationType)

// 	if len(result) != len(runes) {
// 		t.Errorf("Expected length %d, but got %d", len(runes), len(result))
// 	}

// 	// Check if the result runes are not in the same order
// 	if string(result) == string(originalRunes) {
// 		t.Errorf("Expected different order of runes, but got the same")
// 	}
// }

// func TestInsertHTML(t *testing.T) {
// 	m := &Mutator{}
// 	runes := []rune("Hello world")

// 	result := m.insertHTML(runes)

// 	if len(result) <= len(runes) {
// 		t.Errorf("Expected length greater than %d, but got %d", len(runes), len(result))
// 	}

// 	if string(result) == string(runes) {
// 		t.Errorf("Expected different runes after HTML insertion")
// 	}
// }

// func TestInsertFakeWord(t *testing.T) {
// 	m := &Mutator{}
// 	runes := []rune("Hello world")
// 	mutationType := MutationType{Length: 5}

// 	result := m.insertFakeWord(runes, mutationType)

// 	if len(result) <= len(runes) {
// 		t.Errorf("Expected length greater than %d, but got %d", len(runes), len(result))
// 	}

// 	if string(result) == string(runes) {
// 		t.Errorf("Expected different runes after fake word insertion")
// 	}
// }

// func TestReverseSubstring(t *testing.T) {
// 	m := &Mutator{}
// 	runes := []rune("Javier")
// 	mutationType := MutationType{Length: 1}

// 	result := m.reverseSubstring(runes, mutationType)

// 	if len(result) != len(runes) {
// 		t.Errorf("Expected length %d, but got %d", len(runes), len(result))
// 	}

// 	// Check if the substring is reversed
// 	reversed := m.reverseSubstring(runes, mutationType)
// 	if string(result) != string(reversed) {
// 		t.Errorf("Expected substring to be reversed. %s but got %s", string(result), string(reversed))
// 	}
// }

// func TestInsertRandomJSON(t *testing.T) {
// 	m := &Mutator{}
// 	runes := []rune("Hello world")
// 	mutationType := MutationType{Length: 3}

// 	result := m.insertRandomJSON(runes, mutationType.Length)

// 	if len(result) <= len(runes) {
// 		t.Errorf("Expected length greater than %d, but got %d", len(runes), len(result))
// 	}

// 	if string(result) == string(runes) {
// 		t.Errorf("Expected different runes after random JSON insertion")
// 	}
// }
