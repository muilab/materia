package mui

import (
	"github.com/aerogo/nano"
)

// MaterialSet is a set of materials.
type MaterialSet struct {
	ID          ID     `json:"id"`
	Name        string `json:"name" editable:"true"`
	MaterialIDs []ID   `json:"materials"`

	HasDraft
	HasCreator
	HasEditor
}

// Materials returns all materials of the given set.
func (set *MaterialSet) Materials() []*Material {
	objects := DB.GetMany("Material", set.MaterialIDs)
	materials := make([]*Material, len(objects))

	for index, obj := range objects {
		materials[index] = obj.(*Material)
	}

	return materials
}

// Link returns the permalink for this object.
func (set *MaterialSet) Link() string {
	return "/materialset/" + set.ID
}

// ImageLink returns the image URL of the main material.
func (set *MaterialSet) ImageLink(size string) string {
	if len(set.MaterialIDs) > 0 {
		return set.MainMaterial().ImageLink(size)
	}

	// return fmt.Sprintf("/images/materials/%s/%d.jpg", size, rand.Intn(10)+1)
	return "/images/errors/404.png"
}

// MainMaterial returns the main material, if available.
func (set *MaterialSet) MainMaterial() *Material {
	return nil
}

// Remove material with the given ID from the material list.
func (set *MaterialSet) Remove(materialID string) bool {
	for index, item := range set.MaterialIDs {
		if item == materialID {
			set.MaterialIDs = append(set.MaterialIDs[:index], set.MaterialIDs[index+1:]...)
			return true
		}
	}

	return false
}

// GetMaterialSet returns a single material by the given ID.
func GetMaterialSet(id string) (*MaterialSet, error) {
	obj, err := DB.Get("MaterialSet", id)

	if err != nil {
		return nil, err
	}

	return obj.(*MaterialSet), nil
}

// StreamMaterialSets returns a stream of all materials.
func StreamMaterialSets() chan *MaterialSet {
	channel := make(chan *MaterialSet, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("MaterialSet") {
			channel <- obj.(*MaterialSet)
		}

		close(channel)
	}()

	return channel
}

// AllMaterialSets returns a slice of all blog posts.
func AllMaterialSets() []*MaterialSet {
	var all []*MaterialSet

	for obj := range StreamMaterialSets() {
		all = append(all, obj)
	}

	return all
}

// FilterMaterialSets filters all material sets by a custom function.
func FilterMaterialSets(filter func(*MaterialSet) bool) []*MaterialSet {
	var filtered []*MaterialSet

	for obj := range StreamMaterialSets() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
