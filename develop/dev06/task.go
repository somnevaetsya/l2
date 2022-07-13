package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Flags struct {
	f []int
	s bool
	d string
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var example = []string{
	"Winter: white: snow: frost",
	"Spring: green: grass: warm",
	"Summer: colorful: blossom: hot",
	"Autumn: yellow: leaves: cool",
	"AAWDAWD AWDOKGDSODGOSD: awdawdAWDAWD",
	"AWDOKGDSODGOSD adfdsgsgwed awdwa d",
}

func ParseFlags() (flags Flags) {
	flag.BoolVar(&flags.s, "s", false, "separated")
	flag.StringVar(&flags.d, "d", "", "delimiter")
	var buf string
	flag.StringVar(&buf, "f", "", "fields")
	flag.Parse()
	var fields []int
	split := strings.Split(buf, ",")
	for _, item := range split {
		newSplit := strings.Split(item, "-")

		firstItem, err := strconv.Atoi(newSplit[0])
		if err != nil {
			break
		}
		if len(newSplit) > 1 {
			secondItem, err := strconv.Atoi(newSplit[1])
			if err != nil {
				break
			}
			for j := firstItem; j <= secondItem; j++ {
				if !contains(fields, j) {
					fields = append(fields, j)
				}
			}
		} else {
			fields = append(fields, firstItem)
		}
	}
	flags.f = fields
	return
}

func Cut(input []string, flags Flags) []string {
	var delimiter string
	if flags.d == "" {
		delimiter = "\t"
	} else {
		delimiter = flags.d
	}
	var result []string
	for _, oneString := range input {
		array := strings.Split(oneString, delimiter)
		if len(array) == 1 && flags.s {
			continue
		}
		for _, field := range flags.f {
			if field > len(array) {
				break
			}
			result = append(result, array[field-1])
		}
	}
	return result
}

func main() {
	flags := ParseFlags()
	fmt.Println(flags)
	result := Cut(example, flags)
	for _, item := range result {
		fmt.Println(item)
	}
}
