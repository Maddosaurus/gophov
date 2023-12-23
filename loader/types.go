package loader

import (
	"fmt"
	"time"
)

type XMPSidecar struct {
	Base            *BaseInfo
	ReferencedImage string
}

type Image struct {
	Base     *BaseInfo
	Filename string
}

type BaseInfo struct {
	Rating      int
	Label       string
	LensModel   string
	LensMake    string
	CameraModel string
	CameraMake  string
	CreatedDate time.Time
	ModifyDate  time.Time
}

func (i BaseInfo) String() string {
	return fmt.Sprintf("Rating: %d, Label: %s, Lens Model: %s, Lens Make: %s, Camera Model: %s, Camera Make: %s, Created: %s, Modified: %s",
		i.Rating, i.Label, i.LensModel, i.LensMake, i.CameraModel, i.CameraMake, i.CreatedDate.Local(), i.ModifyDate.Local())
}
