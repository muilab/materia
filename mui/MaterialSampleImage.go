package mui

import (
	"bytes"
	"fmt"
	"image"
	"time"

	"github.com/blitzprog/imageoutput"
)

// AddSampleImageBytes accepts a byte buffer that represents an image file and adds it as a sample image.
func (material *Material) AddSampleImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return material.AddSampleImage(&imageoutput.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// AddSampleImage adds a sample image to the material.
func (material *Material) AddSampleImage(metaImage *imageoutput.MetaImage) error {
	var lastError error

	// Save the different image formats
	for _, output := range materialImageOutputs {
		err := output.Save(metaImage, fmt.Sprintf("%s-sample-%d", material.ID, len(material.Samples)))

		if err != nil {
			lastError = err
		}
	}

	sample := ImageFile{
		Extension:    metaImage.Extension(),
		Width:        metaImage.Image.Bounds().Dx(),
		Height:       metaImage.Image.Bounds().Dy(),
		AverageColor: GetAverageColor(metaImage.Image),
		LastModified: time.Now().Unix(),
	}

	material.Samples = append(material.Samples, sample)
	return lastError
}
