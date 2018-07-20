package mui

import (
	"fmt"

	"github.com/aerogo/nano"
)

// Material is a material that can be used for CG and manufacturing.
type Material struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name" editable:"true"`
	Description string    `json:"description" editable:"true"`
	Image       ImageFile `json:"image"`
	SampleIDs   []ID      `json:"samples"`

	HasCreator
	HasEditor
}

// Link returns the permalink for this object.
func (material *Material) Link() string {
	return "/material/" + material.ID
}

// HasImage tells you whether the material has an image.
func (material *Material) HasImage() bool {
	return material.Image.Extension != ""
}

// Samples returns all the samples for this material.
func (material *Material) Samples() []*MaterialSample {
	objects := DB.GetMany("MaterialSample", material.SampleIDs)
	result := make([]*MaterialSample, len(objects))

	for index, obj := range objects {
		result[index] = obj.(*MaterialSample)
	}

	return result
}

// ImageLink returns the image URL of the material.
func (material *Material) ImageLink(size string) string {
	if !material.HasImage() {
		return "/images/errors/404.png"
	}

	extension := ".jpg"

	if size == "original" {
		extension = material.Image.Extension
	}

	return fmt.Sprintf("/images/materials/%s/%s%s?%d", size, material.ID, extension, material.Image.LastModified)
}

// GetMaterial returns a single material by the given ID.
func GetMaterial(id string) (*Material, error) {
	obj, err := DB.Get("Material", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Material), nil
}

// StreamMaterials returns a stream of all materials.
func StreamMaterials() chan *Material {
	channel := make(chan *Material, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Material") {
			channel <- obj.(*Material)
		}

		close(channel)
	}()

	return channel
}

// AllMaterials returns a slice of all blog posts.
func AllMaterials() []*Material {
	var all []*Material

	for obj := range StreamMaterials() {
		all = append(all, obj)
	}

	return all
}

// FilterMaterials filters all materials by a custom function.
func FilterMaterials(filter func(*Material) bool) []*Material {
	var filtered []*Material

	for obj := range StreamMaterials() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
