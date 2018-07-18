package mui

import "github.com/aerogo/nano"

// Material is a material that can be used for CG and manufacturing.
type Material struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`

	HasCreator
	HasEditor
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
