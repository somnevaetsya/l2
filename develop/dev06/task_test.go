package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleTest = []string{
	"Winter: white: snow: frost",
	"Spring: green: grass: warm",
	"Summer: colorful: blossom: hot",
	"Autumn: yellow: leaves: cool",
	"AAWDAWD AWDOKGDSODGOSD: awdawdAWDAWD",
	"AWDOKGDSODGOSD adfdsgsgwed awdwa d",
}

func TestFirst(t *testing.T) {
	flags := Flags{f: []int{1, 2, 3, 4}, s: true, d: ":"}
	require.Equal(t, []string{"Winter", " white", " snow", " frost", "Spring", " green", " grass", " warm", "Summer", " colorful", " blossom", " hot", "Autumn", " yellow", " leaves", " cool", "AAWDAWD AWDOKGDSODGOSD", " awdawdAWDAWD"}, Cut(exampleTest, flags), "firsttest")
}

func TestSecond(t *testing.T) {
	flags := Flags{f: []int{1, 2}, s: false, d: ":"}
	require.Equal(t, []string{"Winter", " white", "Spring", " green", "Summer", " colorful", "Autumn", " yellow", "AAWDAWD AWDOKGDSODGOSD", " awdawdAWDAWD", "AWDOKGDSODGOSD adfdsgsgwed awdwa d"}, Cut(exampleTest, flags), "firsttest")
}

func TestThird(t *testing.T) {
	flags := Flags{f: []int{1, 2}, s: false, d: " "}
	require.Equal(t, []string{"Winter:", "white:", "Spring:", "green:", "Summer:", "colorful:", "Autumn:", "yellow:", "AAWDAWD", "AWDOKGDSODGOSD:", "AWDOKGDSODGOSD", "adfdsgsgwed"}, Cut(exampleTest, flags), "firsttest")
}

func TestForth(t *testing.T) {
	flags := Flags{f: []int{1, 2}, s: false, d: ""}
	require.Equal(t, []string{"Winter: white: snow: frost", "Spring: green: grass: warm", "Summer: colorful: blossom: hot", "Autumn: yellow: leaves: cool", "AAWDAWD AWDOKGDSODGOSD: awdawdAWDAWD", "AWDOKGDSODGOSD adfdsgsgwed awdwa d"}, Cut(exampleTest, flags), "firsttest")
}

func TestFifth(t *testing.T) {
	flags := Flags{f: []int{1, 2}, s: true, d: ""}
	require.Equal(t, []string(nil), Cut(exampleTest, flags), "firsttest")
}
