package github.com/jaavier/mutator

import (
	"testing"
)

func TestCreateDict(t *testing.T) {
	content := "word1 word2 word2 word3 word3 word3 word4 word4 word4 word4 word5 word5 word5 word5 word5 word6 word6 word6 word6 word6 word6 word7"
	totalWords := 2
	kind := "letters"

	dict := CreateDict(content, totalWords, kind)
	if len(dict) != totalWords {
		t.Errorf("Expected %d words, but got %d", totalWords, len(dict))
	}

	for word, count := range dict {
		if len(word) < 4 {
			t.Errorf("Word %s is shorter than 4 characters", word)
		}
		if count < 5 {
			t.Errorf("Word %s has fewer than 5 repetitions", word)
		}
	}
}

func TestCreateDictWithNumbers(t *testing.T) {
	content := "word1 1234 1234 1234 1234 1234 word2 5678 5678 5678 5678 5678"
	totalWords := 2
	kind := "numbers"

	dict := CreateDict(content, totalWords, kind)

	if len(dict) != totalWords {
		t.Errorf("Expected %d words, but got %d", totalWords, len(dict))
	}

	for word, count := range dict {
		if len(word) < 4 {
			t.Errorf("Word %s is shorter than 4 characters", word)
		}
		if count < 5 {
			t.Errorf("Word %s has fewer than 5 repetitions", word)
		}
		if !isValidWord(word, false, true, false) {
			t.Errorf("Word %s does not match the specified kind (numbers)", word)
		}
	}
}

func TestCreateDictWithSpecialCharacters(t *testing.T) {
	content := "hello! hello! hello! hello! hello! $world $world $world $world $world"
	totalWords := 1
	kind := "special"

	dict := CreateDict(content, totalWords, kind)

	if len(dict) != totalWords {
		t.Errorf("Expected %d words, but got %d", totalWords, len(dict))
	}

	for word, count := range dict {
		if len(word) < 4 {
			t.Errorf("Word %s is shorter than 4 characters", word)
		}
		if count < 5 {
			t.Errorf("Word %s has fewer than 5 repetitions", word)
		}
		if !isValidWord(word, false, false, true) {
			t.Errorf("Word %s does not match the specified kind (special characters)", word)
		}
	}
}

func TestCreateDictWithMixedCharacters(t *testing.T) {
	content := "word1 word2 word2 word3 word3 word3 word4 word4 word4 word4 word5 word5 word5 word5 word5 hello! hello! hello! hello! hello!"
	totalWords := 2
	kind := "letters,special"

	dict := CreateDict(content, totalWords, kind)

	if len(dict) != totalWords {
		t.Errorf("Expected %d words, but got %d", totalWords, len(dict))
	}

	for word, count := range dict {
		if len(word) < 4 {
			t.Errorf("Word %s is shorter than 4 characters", word)
		}
		if count < 5 {
			t.Errorf("Word %s has fewer than 5 repetitions", word)
		}
		if !isValidWord(word, true, false, true) {
			t.Errorf("Word %s does not match the specified kind (letters and special characters)", word)
		}
	}
}
