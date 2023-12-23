package imagemeta

import (
	"github.com/Maddosaurus/gophov/loader"
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/xmp"
	"log"
	"os"
)

// ParseXMP loads meta information from XMP files in a given path
func ParseXMP(xmpPath string) *loader.XMPSidecar {
	if len(xmpPath) < 2 {
		log.Printf("Warning: XMP Path is too short")
		return nil
	}
	f, err := os.Open(xmpPath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	e, err := xmp.ParseXmp(f)
	if err != nil {
		log.Fatalf("parser error: %v", err)
	}

	x := &loader.XMPSidecar{
		Base: &loader.BaseInfo{
			Rating:      int(e.Basic.Rating),
			Label:       e.Basic.Label,
			LensModel:   e.Aux.Lens,
			LensMake:    e.Aux.LensInfo,
			CameraModel: e.Tiff.Model,
			CameraMake:  e.Tiff.Make,
			CreatedDate: e.Basic.CreateDate,
			ModifyDate:  e.Basic.ModifyDate,
		},
		ReferencedImage: e.CRS.RawFileName,
	}

	log.Printf("Loaded item: %v", x.Base)
	log.Printf("Create XMP Sidecar: %+v", x.ReferencedImage)
	return x
}

// ParsePhoto loads meta information from a given photo
func ParsePhoto(photoPath string) *loader.Image {
	f, err := os.Open(photoPath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	e, err := imagemeta.Decode(f)
	if err != nil {
		log.Fatalf("parser error: %v", err)
	}

	i := &loader.Image{
		Base: &loader.BaseInfo{
			Rating:      int(e.Rating),
			Label:       "",
			LensModel:   e.LensModel,
			LensMake:    e.LensMake,
			CameraModel: e.Model,
			CameraMake:  e.Make,
			CreatedDate: e.CreateDate(),
			ModifyDate:  e.ModifyDate(),
		}, Filename: e.DocumentName,
	}

	log.Printf("Loaded item: %v", i.Base)
	log.Printf("Create Image: %+v", i.Filename)
	return i
}
