package main

import (
	"flag"
	"github.com/Maddosaurus/gophov/loader/imagemeta"
)

func main() {
	xmpPath := flag.String("x", "", "Filepath to an XMP file you want to extract info of")
	photoPath := flag.String("p", "", "Filepath to a photo file you want to extract info of")

	flag.Parse()

	if isFlagPassed("x") {
		imagemeta.ParseXMP(*xmpPath)
	}
	if isFlagPassed("p") {
		imagemeta.ParsePhoto(*photoPath)
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
