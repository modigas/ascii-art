package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(name string) (map[byte][]string, error) {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		err = errors.New("Invalid Font Name")
		return nil, err
	}
	data := make(map[byte][]string)
	content := strings.Split(string(dat), "\n\n")
	for i, el := range content {
		if i == 0 {
			el = el[1:]
		}
		data[byte(i+32)] = strings.Split(el, "\n")
	}
	return data, nil
}

func formStringASCII(word string, abc map[byte][]string) (string, error) {
	var err error
	words := strings.Split(word, "\\n")
	str := ""
	for _, w := range words {
		res := make([][]string, 0)
		for _, r := range w {
			if r < 32 || r > 126 {
				err = fmt.Errorf("Invalid Symbol")
				return str, err
			}
			res = append(res, abc[byte(r)])
		}
		for i := 0; i < 8; i++ {
			for _, el := range res {
				str += el[i]
			}
			str += "\n"
		}
	}
	return str, err
}
