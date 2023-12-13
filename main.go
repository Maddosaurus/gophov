package main

import (
	"log"
	"os"

	"github.com/evanoberholster/imagemeta/xmp"
)

var (
	xmppath  = "/home/mat/Downloads/img/_DSC3433.xmp"
	xmppath3 = "/home/mat/Downloads/img/_DSC3433_3.xmp"
)

func main() {
	parseXMP(xmppath)
	parseXMP(xmppath3)
}

func parseXMP(xmppath string) {
	f, err := os.Open(xmppath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	e, err := xmp.ParseXmp(f)
	if err != nil {
		log.Fatalf("parser error: %v", err)
	}

	log.Printf("Label: %s, Rating: %d\nFull Meta: %v", e.Basic.Label, e.Basic.Rating, e)
}
