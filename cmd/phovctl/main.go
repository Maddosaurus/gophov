/*
Phovctl provides CLI control of all common features of GoPhoV.
*/
package main

import (
	"flag"
	"github.com/Maddosaurus/gophov/filter"
	"github.com/Maddosaurus/gophov/loader"
	"github.com/Maddosaurus/gophov/loader/imagemeta"
	"log/slog"
	"os"
)

func main() {
	setupLogging()

	xmpPath := flag.String("x", "", "Filepath to an XMP file you want to extract info of")
	photoPath := flag.String("p", "", "Filepath to a photo file you want to extract info of")
	dir := flag.String("d", "", "Path to a folder to read everything in")
	rating := flag.Int("r", 0, "Filter for minimum number of stars")

	flag.Parse()

	var xmps []*loader.XMPSidecar

	if isFlagPassed("x") {
		imagemeta.ParseXMP(*xmpPath)
	}
	if isFlagPassed("p") {
		imagemeta.ParsePhoto(*photoPath)
	}
	if isFlagPassed("d") {
		xmps = imagemeta.LoadDir(*dir)
	}
	if isFlagPassed("r") {
		filter.MinStars(xmps, *rating)
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
