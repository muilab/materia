package mui

import "fmt"

// MaterialSample represents a material sample.
type MaterialSample struct {
	ID    string    `json:"id"`
	Image ImageFile `json:"image"`
}

// ImageLink returns the image URL of the sample.
func (sample *MaterialSample) ImageLink(size string) string {
	extension := ".jpg"

	if size == "original" {
		extension = sample.Image.Extension
	}

	return fmt.Sprintf("/images/samples/%s/%s%s?%d", size, sample.ID, extension, sample.Image.LastModified)
}
