package main

import (
	"fmt"
)

// Warning messages
var (
	// Usage - shows how to use program
	Usage = "\nUSAGE: ./ascii-art \"Your text goes here\"[| shadow | standard | thinkertoy] (if not indicated standard font will be used)\n"
)

const (
	// Height - determines the height of the symbol
	Height = 8
	// Width - determines the width of the symbol
	Width = 9
)

var NewLine = []string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"}

// Colors
var (
	// Reset - reset colors
	Reset = "\033[0m"
	// Red - set color to red
	Red = "\033[31m"
	// Green - set color to green
	Green = "\033[32m"
	// Yellow - set color to yellow
	Yellow = "\033[33m"
	// Blue - set color to blue
	Blue = "\033[34m"
	// Purple - set color to purple
	Purple = "\033[35m"
	// Cyan - set color to cyan
	Cyan = "\033[36m"
	// Gray - set color to gray
	Gray = "\033[37m"
	// White - set color to white
	White = "\033[97m"
)

func errorHandler(err error) {
	fmt.Printf(Red+"\nOops, something went wrong :("+Reset+"\n%s%s\n", err.Error(), Usage)
}
