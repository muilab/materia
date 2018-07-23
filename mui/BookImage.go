package mui

import (
	"bytes"
	"image"
	"path"
	"time"

	"github.com/blitzprog/imageoutput"
)

const (
	// BookImageLargeWidth is the minimum width in pixels of a large book image.
	BookImageLargeWidth = 512

	// BookImageLargeHeight is the minimum height in pixels of a large book image.
	BookImageLargeHeight = 512

	// BookImageMediumWidth is the minimum width in pixels of a medium book image.
	BookImageMediumWidth = 256

	// BookImageMediumHeight is the minimum height in pixels of a medium book image.
	BookImageMediumHeight = 256

	// BookImageWebPQuality is the WebP quality of book images.
	BookImageWebPQuality = 70

	// BookImageJPEGQuality is the JPEG quality of book images.
	BookImageJPEGQuality = 70

	// BookImageQualityBonusLowDPI ...
	BookImageQualityBonusLowDPI = 12

	// BookImageQualityBonusLarge ...
	BookImageQualityBonusLarge = 10

	// BookImageQualityBonusMedium ...
	BookImageQualityBonusMedium = 15

	// BookImageQualityBonusSmall ...
	BookImageQualityBonusSmall = 15
)

// Define the book image outputs
var bookImageOutputs = []imageoutput.Output{
	// Original at full size
	&imageoutput.OriginalFile{
		Directory: path.Join(Root, "images/books/original/"),
		Width:     0,
		Height:    0,
	},

	// JPEG - Large
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/books/large/"),
		Width:     BookImageLargeWidth,
		Height:    BookImageLargeHeight,
		Quality:   BookImageJPEGQuality + BookImageQualityBonusLowDPI + BookImageQualityBonusLarge,
	},

	// JPEG - Medium
	&imageoutput.JPEGFile{
		Directory: path.Join(Root, "images/books/medium/"),
		Width:     BookImageMediumWidth,
		Height:    BookImageMediumHeight,
		Quality:   BookImageJPEGQuality + BookImageQualityBonusLowDPI + BookImageQualityBonusMedium,
	},

	// WebP - Large
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/books/large/"),
		Width:     BookImageLargeWidth,
		Height:    BookImageLargeHeight,
		Quality:   BookImageWebPQuality + BookImageQualityBonusLowDPI + BookImageQualityBonusLarge,
	},

	// WebP - Medium
	&imageoutput.WebPFile{
		Directory: path.Join(Root, "images/books/medium/"),
		Width:     BookImageMediumWidth,
		Height:    BookImageMediumHeight,
		Quality:   BookImageWebPQuality + BookImageQualityBonusLowDPI + BookImageQualityBonusMedium,
	},
}

// SetImageBytes accepts a byte buffer that represents an image file and updates the book image.
func (book *Book) SetImageBytes(data []byte) error {
	// Decode
	img, format, err := image.Decode(bytes.NewReader(data))

	if err != nil {
		return err
	}

	return book.SetImage(&imageoutput.MetaImage{
		Image:  img,
		Format: format,
		Data:   data,
	})
}

// SetImage sets the book image to the given MetaImage.
func (book *Book) SetImage(metaImage *imageoutput.MetaImage) error {
	var lastError error

	// Save the different image formats
	for _, output := range bookImageOutputs {
		err := output.Save(metaImage, book.ID)

		if err != nil {
			lastError = err
		}
	}

	book.Image.Extension = metaImage.Extension()
	book.Image.Width = metaImage.Image.Bounds().Dx()
	book.Image.Height = metaImage.Image.Bounds().Dy()
	book.Image.AverageColor = GetAverageColor(metaImage.Image)
	book.Image.LastModified = time.Now().Unix()
	return lastError
}
