package mui

import "github.com/aerogo/nano"

// Book contains a collection of material sets.
type Book struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	MaterialSetIDs []string `json:"materialSets"`

	HasCreator
}

// MaterialSets returns all material sets of the given book.
func (book *Book) MaterialSets() []*MaterialSet {
	objects := DB.GetMany("MaterialSet", book.MaterialSetIDs)
	sets := make([]*MaterialSet, len(objects))

	for index, obj := range objects {
		sets[index] = obj.(*MaterialSet)
	}

	return sets
}

// GetBook returns a single book by the given ID.
func GetBook(id string) (*Book, error) {
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

// FilterBooks filters all books by a custom function.
func FilterBooks(filter func(*Book) bool) []*Book {
	var filtered []*Book

	for obj := range StreamBooks() {
		if filter(obj) {
			filtered = append(filtered, obj)
		}
	}

	return filtered
}
