package main

import (
	"flag"
	"github.com/Maddosaurus/gophov/loader/imagemeta"
	"log/slog"
	"os"
)

func main() {
	setupLogging()

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

func setupLogging() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
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
