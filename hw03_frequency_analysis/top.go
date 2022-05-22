package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const defaultCount = 10

type item struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	re := regexp.MustCompile(`^[']|[!?.']$`)
	sliceOfStrings := strings.Fields(s)
	mapCounter := make(map[string]int)
	items := []item{}
	for _, word := range sliceOfStrings {
		word = strings.ToLower(word)
		word = re.ReplaceAllString(word, "")
		if word == "-" {
			continue
		}
		if _, inMap := mapCounter[word]; inMap {
			mapCounter[word]++
			for idx, i := range items {
				if i.Word == word {
					items = append(items[:idx], items[idx+1:]...)
					items = append(items, item{word, mapCounter[word]})
				}
			}
		} else {
			mapCounter[word] = 1
			items = append(items, item{word, 1})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return items[i].Word < items[j].Word
		}
		return items[i].Count > items[j].Count
	})

	result := []string{}
	for i := 0; i < defaultCount; {
		if i < len(items) {
			result = append(result, items[i].Word)
			i++
		} else {
			break
		}
	}
	return result
}
