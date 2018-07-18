package mui

import "github.com/aerogo/nano"

// MaterialSet is a set of materials.
type MaterialSet struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	MaterialIDs []string `json:"materials"`

	HasCreator
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