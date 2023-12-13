package loader

import (
	"log"
	"os"

	"github.com/evanoberholster/imagemeta/xmp"
)

// ParseXMP loads meta information from XMP files in a given path
func ParseXMP(xmppath string) {
	if len(xmppath) < 2 {
		log.Printf("Warning: XMP Path is too short")
		return
	}
	f, err := os.Open(xmppath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	e, err := xmp.ParseXmp(f)
	if err != nil {
		log.Fatalf("parser error: %v", err)
	}

	log.Printf("Label: %s, Rating: %d, Lens: %s, Make: %s, Model: %s", e.Basic.Label, e.Basic.Rating, e.Aux.Lens, e.Tiff.Make, e.Tiff.Model)
}
