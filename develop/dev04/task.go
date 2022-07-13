package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	if lenS != lenT {
		return false
	}
	anagramMap := make(map[string]int)

	for i := 0; i < lenS; i++ {
		anagramMap[string(s[i])]++
	}
	for i := 0; i < lenT; i++ {
		anagramMap[string(t[i])]--
	}
	for i := 0; i < lenS; i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}
	return true
}
func GetAnagrams(input *[]string) *map[string][]string {
	result := make(map[string][]string)
	var flag bool
	for _, word := range *input {
		flag = false
		if len(word) <= 2 {
			continue
		}
		for k := range result {
			if isAnagram(strings.ToLower(word), strings.ToLower(k)) {
				flag = true
				result[k] = append(result[k], strings.ToLower(word))
			}
		}
		if flag == false {
			result[word] = append(result[word], strings.ToLower(word))
		}
	}
	for _, slice := range result {
		sort.Strings(slice)
	}
	return &result
}

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "сколит", "столик"}
	fmt.Println(GetAnagrams(&input))
}
