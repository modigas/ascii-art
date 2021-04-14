package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 && len(args) < 4 {
		fontName := "standard"
		if len(args) == 3 {
			fontName = args[2]
		}
		abc, err := readFile(fontName + ".txt")
		if err != nil {
			errorHandler(err)
			return
		}
		str, err := formStringASCII(args[1], abc)
		if err != nil {
			errorHandler(err)
			return
		}
		fmt.Println(str)
	} else {

	}
}
