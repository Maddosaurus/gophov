// Package loader implements functions to load metadata from image and sidecar files.
package loader

import (
	"time"
)

var (
	_ MetaItem = (*XMPSidecar)(nil)
	_ MetaItem = (*Image)(nil)
)

// MetaItem represents common functions all gophov items have
type MetaItem interface {
	GetReferencedImage() string
	GetInfo() *BaseInfo
}

// XMPSidecar represents information specific to XMP sidecar files
type XMPSidecar struct {
	Base            *BaseInfo
	ReferencedImage string
}

// Image represents information specific to images
type Image struct {
	Base *BaseInfo
}

// BaseInfo represents common data amongst all managed types
type BaseInfo struct {
	Rating      int
	Label       string
	LensModel   string
	LensMake    string
	CameraModel string
	CameraMake  string
	CreatedDate time.Time
	ModifyDate  time.Time
	Filename    string
}

func (x XMPSidecar) GetReferencedImage() string {
	return x.ReferencedImage
}

func (x XMPSidecar) GetInfo() *BaseInfo {
	return x.Base
}

func (i Image) GetReferencedImage() string {
	return i.Base.Filename
}

func (i Image) GetInfo() *BaseInfo {
	return i.Base
}

//func (i BaseInfo) String() string {
//	return fmt.Sprintf("Rating: %d, Label: %s, Lens Model: %s, Lens Make: %s, Camera Model: %s, Camera Make: %s, Created: %s, Modified: %s",
//		i.Rating, i.Label, i.LensModel, i.LensMake, i.CameraModel, i.CameraMake, i.CreatedDate.Local(), i.ModifyDate.Local())
//}
