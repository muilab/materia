package mui

import (
	"bytes"
	"image"
	"path"
	"time"

	"github.com/blitzprog/imageoutput"
)

// Define the material sample image outputs
var materialSampleImageOutputs = []imageoutput.Output{
	// Original at full size
	&imageoutput.OriginalFile{
		Directory: path.Join(Root, "images/samples/original/"),
		Width:     0,
		Height:    0,
	},

	// JPEG - Large
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/samples/large/"),
		Width:     MaterialImageLargeWidth,
		Height:    MaterialImageLargeHeight,
		Quality:   MaterialImageJPEGQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/samples/medium/"),
		Width:     MaterialImageMediumWidth,
		Height:    MaterialImageMediumHeight,
		Quality:   MaterialImageJPEGQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusMedium,
	},

	// WebP - Large
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/samples/large/"),
		Width:     MaterialImageLargeWidth,
		Height:    MaterialImageLargeHeight,
		Quality:   MaterialImageWebPQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/samples/medium/"),
		Width:     MaterialImageMediumWidth,
		Height:    MaterialImageMediumHeight,
		Quality:   MaterialImageWebPQuality + MaterialImageQualityBonusLowDPI + MaterialImageQualityBonusMedium,
	},
}

// SetImageBytes accepts a byte buffer that represents an image file and updates the sample image.
func (sample *MaterialSample) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return sample.SetImage(&imageoutput.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the sample image to the given MetaImage.
func (sample *MaterialSample) SetImage(metaImage *imageoutput.MetaImage) error {
	var lastError error

	// Save the different image formats
	for _, output := range materialSampleImageOutputs {
		err := output.Save(metaImage, sample.ID)

		if err != nil {
			lastError = err
		}
	}

	sample.Image.Extension = metaImage.Extension()
	sample.Image.Width = metaImage.Image.Bounds().Dx()
	sample.Image.Height = metaImage.Image.Bounds().Dy()
	sample.Image.AverageColor = GetAverageColor(metaImage.Image)
	sample.Image.LastModified = time.Now().Unix()
	return lastError
}
