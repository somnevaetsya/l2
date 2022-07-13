package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFirst(t *testing.T) {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "сколит", "столик"}
	require.Equal(t, &map[string][]string{"листок": {"листок", "сколит", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}}, GetAnagrams(&input), "firsttest")
}

func TestSecond(t *testing.T) {
	input := []string{"апельсин", "спаниель", "лексор", "кресло", "склеро", "скорел"}
	require.Equal(t, &map[string][]string{"апельсин": {"апельсин", "спаниель"}, "лексор": {"кресло", "лексор", "склеро", "скорел"}}, GetAnagrams(&input), "firsttest")
}

func TestThird(t *testing.T) {
	input := []string{"куоср", "кроус", "курос", "сокур", "сурок", "уркос"}
	require.Equal(t, &map[string][]string{"куоср": {"кроус", "куоср", "курос", "сокур", "сурок", "уркос"}}, GetAnagrams(&input), "firsttest")
}

func TestForth(t *testing.T) {
	input := []string{"а", "о", "б", "покраснение", "пенсионерка"}
	require.Equal(t, &map[string][]string{"покраснение": {"пенсионерка", "покраснение"}}, GetAnagrams(&input), "firsttest")
}

func TestAll(t *testing.T) {
	input := []string{"апельсин", "спаниель", "лексор", "кресло", "склеро", "скорел", "куоср", "кроус", "курос", "сокур", "сурок", "уркос", "а", "о", "б", "покраснение", "пенсионерка", "пятак", "пятка", "тяпка", "листок", "слиток", "сколит", "столик"}
	require.Equal(t, &map[string][]string{"апельсин": {"апельсин", "спаниель"}, "куоср": {"кроус", "куоср", "курос", "сокур", "сурок", "уркос"}, "лексор": {"кресло", "лексор", "склеро", "скорел"}, "листок": {"листок", "сколит", "слиток", "столик"}, "покраснение": {"пенсионерка", "покраснение"}, "пятак": {"пятак", "пятка", "тяпка"}}, GetAnagrams(&input), "firsttest")
}
