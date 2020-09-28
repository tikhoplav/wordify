package wordify

import (
	"strings"
	"bufio"
)

func Wordify(source string, mappings ...func (string) string) ([]string, error) {
	var result []string
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		for _, m := range mappings {
			word = m(word)
			if word == "" {
				break
			}
		}
		if word != "" {
			result = append(result, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}