package mui

import "fmt"

// MaterialSample represents a material sample.
type MaterialSample struct {
	ID         ID        `json:"id"`
	MaterialID ID        `json:"materialId"`
	Image      ImageFile `json:"image"`
}

// Material returns the material of the sample.
func (sample *MaterialSample) Material() *Material {
	material, _ := GetMaterial(sample.MaterialID)
	return material
}

// ImageLink returns the image URL of the sample.
func (sample *MaterialSample) ImageLink(size string) string {
	extension := ".jpg"

	if size == "original" {
		extension = sample.Image.Extension
	}

	return fmt.Sprintf("/images/samples/%s/%s%s?%d", size, sample.ID, extension, sample.Image.LastModified)
}
