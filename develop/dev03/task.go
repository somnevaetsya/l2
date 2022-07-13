package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Flags struct {
	k      int
	n      bool
	r      bool
	u      bool
	c      bool
	Params []string
}

func ParseFlag() (flags Flags) {
	flag.IntVar(&flags.k, "k", 0, "column")
	flag.BoolVar(&flags.n, "n", false, "numeric sort")
	flag.BoolVar(&flags.r, "r", false, "reverse order")
	flag.BoolVar(&flags.u, "u", false, "unique")
	flag.BoolVar(&flags.c, "c", false, "is sorted")
	flag.Parse()
	flags.Params = flag.Args()
	return
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func writeIsSorted(flag bool, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	if flag {
		fmt.Fprintln(w, "File sorted.")
	} else {
		fmt.Fprintln(w, "File isn`t sorted.")
	}

	return w.Flush()
}

func ReverseSlice(data *[]string) {
	for i, j := 0, len(*data)-1; i < j; i, j = i+1, j-1 {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
	}
}

func unique(input []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func StraightSort(data []string) {
	sort.Slice(data, func(i, j int) bool {
		return strings.ToLower(data[i]) < strings.ToLower(data[j])
	})
}

func NumericSort(data *[]string) {
	numeric := GetNumeric(*data)
	alphabet := GetString(*data)
	sort.Slice(numeric, func(i, j int) bool {
		convI, _ := strconv.Atoi(numeric[i])
		convJ, _ := strconv.Atoi(numeric[j])
		return convI < convJ
	})
	StraightSort(alphabet)
	var result []string
	result = append(result, alphabet...)
	result = append(result, numeric...)
	fmt.Println(result)
	*data = result
}

func GetNumeric(data []string) []string {
	var numeric []string
	for _, item := range data {
		if _, err := strconv.Atoi(item); err == nil {
			numeric = append(numeric, item)
		}
	}
	return numeric
}

func GetString(data []string) []string {
	var numeric []string
	for _, item := range data {
		if _, err := strconv.Atoi(item); err != nil {
			numeric = append(numeric, item)
		}
	}
	return numeric
}

func IsSorted(input, data []string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] != data[i] {
			return false
		}
	}
	return true
}

func ColumnSort(data []string, column int) {
	sort.Slice(data, func(i, j int) bool {
		columnsI := strings.Split(data[i], " ")
		columnsJ := strings.Split(data[j], " ")
		return columnsI[column-1] < columnsJ[column-1]
	})
}

func main() {
	flags := ParseFlag()
	data, err := readLines(flags.Params[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	var toCheck []string
	StraightSort(data)
	if flags.k != 0 {
		ColumnSort(data, flags.k)
		if flags.c {
			ColumnSort(toCheck, flags.k)
		}
	}
	if flags.n {
		NumericSort(&data)
		if flags.c {
			NumericSort(&toCheck)
		}
	}
	if flags.r {
		ReverseSlice(&data)
		if flags.c {
			ReverseSlice(&toCheck)
		}
	}
	if flags.u {
		unique(data)
		if flags.c {
			unique(toCheck)
		}
	}
	if flags.c {
		err = writeIsSorted(IsSorted(data, toCheck), flags.Params[1])
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		err = writeLines(data, flags.Params[1])
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
