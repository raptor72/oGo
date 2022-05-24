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
	mapCounter := make(map[string]item)
	items := &[]item{}

	for _, word := range sliceOfStrings {
		word = strings.ToLower(word)
		word = re.ReplaceAllString(word, "")
		if word == "-" {
			continue
		}
		if _, inMap := mapCounter[word]; inMap {
			it := mapCounter[word]
			it.Count++
			mapCounter[word] = it
		} else {
			mapCounter[word] = item{Word: word, Count: 1}
		}
	}
	for _, value := range mapCounter {
		*items = append(*items, value)
	}
	itemSlice := *items

	sort.Slice(itemSlice, func(i, j int) bool {
		if itemSlice[i].Count == itemSlice[j].Count {
			return itemSlice[i].Word < itemSlice[j].Word
		}
		return itemSlice[i].Count > itemSlice[j].Count
	})

	result := []string{}
	for i := 0; i < defaultCount; {
		if i < len(itemSlice) {
			result = append(result, itemSlice[i].Word)
			i++
		} else {
			break
		}
	}
	return result
}
