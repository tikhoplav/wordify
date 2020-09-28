package wordify

import (
	"testing"
	"fmt"
	"strings"
)

func TestMain(t *testing.T) {
	source := `This is a pretty long string.
		With a couple of sentences in it.
		And a couple of new lines.
		This also has a link https://google.com.
		How about (hyphens) and {brackets}, [huh?].
		And finally the <html><tag>markdown</tag></html>.`

	skipShort := func (word string) string {
		if len(word) > 2 {
			return word
		}
		return ""
	}

	toLower := func (word string) string {
		return strings.ToLower(word)
	}

	trimChars := func (word string) string {
		word = strings.TrimRight(word, ". ")
		word = strings.TrimRight(word, ", ")
		word = strings.TrimRight(word, "! ")
		word = strings.TrimRight(word, "? ")
		word = strings.TrimRight(word, ": ")
		return word
	}

	words, err := Wordify(source, trimChars, skipShort, toLower)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	fmt.Printf("%#v", words)
}