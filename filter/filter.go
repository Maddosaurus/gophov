package filter

import (
	"github.com/Maddosaurus/gophov/loader"
	"log/slog"
)

func MinStars(xmps []*loader.XMPSidecar, numStars int) []*loader.XMPSidecar {
	var filtered []*loader.XMPSidecar
	for _, x := range xmps {
		if x.Base.Rating >= numStars {
			filtered = append(filtered, x)
			slog.Debug("Appended item to filtered XMPs", "item", x, "rating", x.Base.Rating)
		}
	}
	return filtered
}
