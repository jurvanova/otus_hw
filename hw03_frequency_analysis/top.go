package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	splitString := strings.Fields(s)

	wordCount := wordsCount(splitString)

	sortedKeys := sortKeys(wordCount)

	if len(sortedKeys) < 10 {
		return sortedKeys
	}
	return sortedKeys[:10]
}

func sortKeys(wordCount map[string]int) []string {
	keys := make([]string, 0, len(wordCount))

	for key := range wordCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if wordCount[keys[i]] == wordCount[keys[j]] {
			return keys[i] < keys[j]
		}
		return wordCount[keys[i]] > wordCount[keys[j]]
	})
	return keys
}

func wordsCount(words []string) map[string]int {
	count := make(map[string]int, len(words))

	for i := range words {
		if value, ok := count[words[i]]; !ok {
			count[words[i]] = 1
		} else {
			count[words[i]] = value + 1
		}
	}
	return count
}
