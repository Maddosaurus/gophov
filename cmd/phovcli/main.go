package main

import (
	"flag"

	"github.com/Maddosaurus/gophov/loader"
)

func main() {
	xmpPath := flag.String("x", "", "Filepath to an XMP file you want to extract info of")

	flag.Parse()

	if isFlagPassed("x") {
		loader.ParseXMP(*xmpPath)
	}
}

func isFlagPassed(name string) bool {
	passed := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})
	return passed
}
