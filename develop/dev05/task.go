package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Flags struct {
	A      int
	B      int
	C      int
	c      bool
	i      bool
	v      bool
	F      bool
	n      bool
	Params []string
}

func ParseFlag() (flags Flags) {
	flag.IntVar(&flags.A, "A", 0, "after")
	flag.IntVar(&flags.B, "B", 0, "before")
	flag.IntVar(&flags.C, "C", 0, "context")
	flag.BoolVar(&flags.c, "c", false, "count")
	flag.BoolVar(&flags.i, "i", false, "ignore-case")
	flag.BoolVar(&flags.v, "v", false, "invert")
	flag.BoolVar(&flags.F, "F", false, "fixed")
	flag.BoolVar(&flags.n, "n", false, "line num")
	flag.Parse()
	flags.Params = flag.Args()
	return
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Grep(flags Flags) {
	file, err := os.Open(flags.Params[1])

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if flags.n {
		for i := 0; i < len(data); i++ {
			data[i] = fmt.Sprintf("%d:%s", i+1, data[i])
		}
	}
	var result []string
	for i, item := range data {
		var matched bool
		if flags.F {
			if flags.i {
				matched = strings.Contains(strings.ToLower(item), strings.ToLower(flags.Params[0]))
			} else {
				matched = strings.Contains(item, flags.Params[0])
			}
		} else {
			if flags.i {
				matched, _ = regexp.MatchString(strings.ToLower(flags.Params[0]), strings.ToLower(item))
			} else {
				matched, _ = regexp.MatchString(flags.Params[0], item)
			}
		}
		if matched && flags.v != true {
			if flags.A != 0 {
				if i+flags.A > len(data) {
					result = append(result, data[i:]...)
				} else {
					result = append(result, data[i:i+flags.A+1]...)
				}
			} else if flags.B != 0 {
				if i-flags.B < 0 {
					result = append(result, data[:i+1]...)
				} else {
					result = append(result, data[i-flags.B:i+1]...)
				}
			} else if flags.C != 0 {
				if i+flags.C > len(data) {
					if i-flags.C < 0 {
						result = append(result, data[:]...)
					} else {
						result = append(result, data[i-flags.C:]...)
					}
				} else {
					if i-flags.C < 0 {
						result = append(result, data[:i+flags.C+1]...)
					} else {
						result = append(result, data[i-flags.C:i+flags.C+1]...)
					}
				}
			} else {
				result = append(result, item)
			}
		} else if !matched && flags.v == true {
			if flags.A != 0 {
				if i+flags.A > len(data) {
					result = append(result, data[i:]...)
				} else {
					result = append(result, data[i:i+flags.A+1]...)
				}
			} else if flags.B != 0 {
				if i-flags.B < 0 {
					result = append(result, data[:i+1]...)
				} else {
					result = append(result, data[i-flags.B:i+1]...)
				}
			} else if flags.C != 0 {
				if i+flags.C > len(data) {
					if i-flags.C < 0 {
						result = append(result, data[:]...)
					} else {
						result = append(result, data[i-flags.C:]...)
					}
				} else {
					if i-flags.C < 0 {
						result = append(result, data[:i+flags.C+1]...)
					} else {
						result = append(result, data[i-flags.C:i+flags.C+1]...)
					}
				}
			} else {
				result = append(result, item)
			}
		}
	}
	if flags.c {
		fmt.Println(len(unique(result)))
	} else {
		for _, item := range unique(result) {
			fmt.Println(item)
		}
	}
	//fmt.Println(unique(result))
}

func main() {
	flags := ParseFlag()
	fmt.Println(flags)
	Grep(flags)
}
