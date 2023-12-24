package imagemeta

import (
	"github.com/Maddosaurus/gophov/loader"
	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/xmp"
	"log/slog"
	"os"
	"path/filepath"
)

// ParseXMP loads meta information from XMP files in a given path
func ParseXMP(xmpPath string) *loader.XMPSidecar {
	if len(xmpPath) < 2 {

		slog.Warn("Warning: XMP Path is too short")
		return nil
	}
	f, err := os.Open(xmpPath)
	if err != nil {
		slog.Error("failed to open file.", "err", err)
	}
	defer f.Close()

	e, err := xmp.ParseXmp(f)
	if err != nil {
		slog.Error("parser error.", "err", err)
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
			Filename:    filepath.Base(xmpPath),
		},
		ReferencedImage: e.CRS.RawFileName,
	}

	slog.Debug("Loaded item.", "item", x.Base)
	slog.Debug("Create XMP Sidecar.", "Sidecar", x.ReferencedImage)
	return x
}

// ParsePhoto loads meta information from a given photo
func ParsePhoto(photoPath string) *loader.Image {
	f, err := os.Open(photoPath)
	if err != nil {
		slog.Error("failed to open file", "err", err)
	}
	defer f.Close()

	e, err := imagemeta.Decode(f)
	if err != nil {
		slog.Error("parser error", err)
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
			Filename:    filepath.Base(photoPath),
		},
	}

	slog.Debug("Loaded item.", "item", i.Base)
	return i
}
