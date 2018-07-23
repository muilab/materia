package mui

import (
	"fmt"

	"github.com/aerogo/nano"
)

// Book is a collection of materials.
type Book struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name" editable:"true"`
	Description string    `json:"description" editable:"true"`
	Image       ImageFile `json:"image"`
	MaterialIDs []ID      `json:"materials"`

	HasDraft
	HasCreator
	HasEditor
}

// Materials returns all materials of the given book.
func (book *Book) Materials() []*Material {
	objects := DB.GetMany("Material", book.MaterialIDs)
	materials := make([]*Material, len(objects))

	for index, obj := range objects {
		materials[index] = obj.(*Material)
	}

	return materials
}

// HasImage tells you whether the book has an image.
func (book *Book) HasImage() bool {
	return book.Image.Extension != ""
}

// ImageLink returns the image URL of the book.
func (book *Book) ImageLink(size string) string {
	if !book.HasImage() {
		return "/images/errors/404.png"
	}

	extension := ".jpg"

	if size == "original" {
		extension = book.Image.Extension
	}

	return fmt.Sprintf("/images/books/%s/%s%s?%d", size, book.ID, extension, book.Image.LastModified)
}

// Link returns the permalink for this object.
func (book *Book) Link() string {
	return "/book/" + book.ID
}

// Remove material with the given ID from the material list.
func (book *Book) Remove(materialID string) bool {
	for index, item := range book.MaterialIDs {
		if item == materialID {
			book.MaterialIDs = append(book.MaterialIDs[:index], book.MaterialIDs[index+1:]...)
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
