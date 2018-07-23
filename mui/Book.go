package mui

import (
	"github.com/aerogo/nano"
)

// Book is a collection of materials.
type Book struct {
	ID          ID     `json:"id"`
	Name        string `json:"name" editable:"true"`
	Description string `json:"description" editable:"true"`
	MaterialIDs []ID   `json:"materials"`

	HasDraft
	HasCreator
	HasEditor
}

// Materials returns all materials of the given set.
func (set *Book) Materials() []*Material {
	objects := DB.GetMany("Material", set.MaterialIDs)
	materials := make([]*Material, len(objects))

	for index, obj := range objects {
		materials[index] = obj.(*Material)
	}

	return materials
}

// Link returns the permalink for this object.
func (set *Book) Link() string {
	return "/book/" + set.ID
}

// Remove material with the given ID from the material list.
func (set *Book) Remove(materialID string) bool {
	for index, item := range set.MaterialIDs {
		if item == materialID {
			set.MaterialIDs = append(set.MaterialIDs[:index], set.MaterialIDs[index+1:]...)
			return true
		}
	}

	return false
}

// GetBook returns a single book by the given ID.
func GetBook(id ID) (*Book, error) {
	obj, err := DB.Get("Book", id)

	if err != nil {
		return nil, err
	}

	return obj.(*Book), nil
}

// StreamBooks returns a stream of all books.
func StreamBooks() chan *Book {
	channel := make(chan *Book, nano.ChannelBufferSize)

	go func() {
		for obj := range DB.All("Book") {
			channel <- obj.(*Book)
		}

		close(channel)
	}()

	return channel
}

// AllBooks returns a slice of all books.
func AllBooks() []*Book {
	var all []*Book

	for obj := range StreamBooks() {
		all = append(all, obj)
	}

	return all
}

// FilterBooks filters all material sets by a custom function.
func FilterBooks(filter func(*Book) bool) []*Book {
	var filtered []*Book

	for obj := range StreamBooks() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
