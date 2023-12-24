package imagemeta

import (
	"github.com/Maddosaurus/gophov/loader"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"
)

func LoadDir(dir string) []*loader.XMPSidecar {
	return iterateRecursive(dir)
}

func iterateRecursive(dir string) []*loader.XMPSidecar {
	var xmps []*loader.XMPSidecar
	filepath.Walk(dir, func(path string, item fs.FileInfo, err error) error {
		if err != nil {
			slog.Error("Failed to walk directory", "directory", path, "error", err)
		}
		slog.Info("Found entry", "name", item.Name(), "dir?", item.IsDir(), "extension", filepath.Ext(item.Name()))
		if item.IsDir() {
			return nil
		}
		if strings.Contains(strings.ToLower(filepath.Ext(item.Name())), ".xmp") {
			xmps = append(xmps, ParseXMP(path))
		}
		return nil
	})

	for _, x := range xmps {
		slog.Info("XMP Entry", "xmp", x)
	}
	return xmps
}
