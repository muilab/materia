package mui

import (
	"bytes"
	"image"
	"path"
	"time"

	"github.com/blitzprog/imageoutput"
)

const (
	// MaterialImageLargeWidth is the minimum width in pixels of a large material image.
	MaterialImageLargeWidth = 512

	// MaterialImageLargeHeight is the minimum height in pixels of a large material image.
	MaterialImageLargeHeight = 512

	// MaterialImageMediumWidth is the minimum width in pixels of a medium material image.
	MaterialImageMediumWidth = 256

	// MaterialImageMediumHeight is the minimum height in pixels of a medium material image.
	MaterialImageMediumHeight = 256

	// MaterialImageWebPQuality is the WebP quality of material images.
	MaterialImageWebPQuality = 70

	// MaterialImageJPEGQuality is the JPEG quality of material images.
	MaterialImageJPEGQuality = 70

	// MaterialImageQualityBonusLowDPI ...
	MaterialImageQualityBonusLowDPI = 12

	// MaterialImageQualityBonusLarge ...
	MaterialImageQualityBonusLarge = 10

	// MaterialImageQualityBonusMedium ...
	MaterialImageQualityBonusMedium = 15

	// MaterialImageQualityBonusSmall ...
	MaterialImageQualityBonusSmall = 15
)

// Define the material image outputs
var materialImageOutputs = []imageoutput.Output{
	// Original at full size
	&imageoutput.OriginalFile{
		Directory: path.Join(Root, "images/materials/original/"),
		Width:     0,
		Height:    0,
	},

	// JPEG - Large
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/materials/large/"),
		Width:     MaterialImageLargeWidth,
		Height:    MaterialImageLargeHeight,
		Quality:   MaterialImageJPEGQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/materials/medium/"),
		Width:     MaterialImageMediumWidth,
		Height:    MaterialImageMediumHeight,
		Quality:   MaterialImageJPEGQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusMedium,
	},

	// WebP - Large
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/materials/large/"),
		Width:     MaterialImageLargeWidth,
		Height:    MaterialImageLargeHeight,
		Quality:   MaterialImageWebPQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/materials/medium/"),
		Width:     MaterialImageMediumWidth,
		Height:    MaterialImageMediumHeight,
		Quality:   MaterialImageWebPQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusMedium,
	},
}

// SetImageBytes accepts a byte buffer that represents an image file and updates the material image.
func (material *Material) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return material.SetImage(&imageoutput.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the material image to the given MetaImage.
func (material *Material) SetImage(metaImage *imageoutput.MetaImage) error {
	var lastError error

	// Save the different image formats
	for _, output := range materialImageOutputs {
		err := output.Save(metaImage, material.ID)

		if err != nil {
			lastError = err
		}
	}

	material.Image.Extension = metaImage.Extension()
	material.Image.Width = metaImage.Image.Bounds().Dx()
	material.Image.Height = metaImage.Image.Bounds().Dy()
	material.Image.AverageColor = GetAverageColor(metaImage.Image)
	material.Image.LastModified = time.Now().Unix()
	return lastError
}
