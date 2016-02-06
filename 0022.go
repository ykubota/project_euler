package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename string = "p022_names.txt"

func main() {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	names := strings.SplitN(string(buf), ",", -1)
	for i, v := range names {
		names[i] = strings.SplitN(v, "\"", -1)[1]
	}
	sort.Strings(names)
	a := 0
	for i, v := range names {
		a += alphabetvalue(v) * (i + 1)
	}

	fmt.Println(a)
}

func alphabetvalue(t string) int {
	a := 0
	for _, v := range t {
		a += int(v) - 64
	}
	return a
}
