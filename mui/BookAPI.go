package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/muilab/materia/mui/utils"
)

// Force interface implementations
var (
	_ api.Newable   = (*Book)(nil)
	_ api.Editable  = (*Book)(nil)
	_ api.Deletable = (*Book)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (book *Book) Authorize(ctx *aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if (action == "edit" || action == "delete") && user.ID != book.CreatedBy {
		return errors.New("Can't edit data from other users")
	}

	return nil
}

// Create books the data for a new material book with data we received from the API request.
func (book *Book) Create(ctx *aero.Context) error {
	book.ID = GenerateID("Book")
	book.Created = utils.DateTimeUTC()
	book.CreatedBy = GetUserFromContext(ctx).ID
	book.MaterialIDs = []ID{}
	return nil
}

// DeleteInContext deletes the book in the given context.
func (book *Book) DeleteInContext(ctx *aero.Context) error {
	return book.Delete()
}

// Delete deletes the book from the database.
func (book *Book) Delete() error {
	// Delete all image files
	for _, output := range bookImageOutputs {
		output.Delete(book.ID)
	}

	DB.Delete("Book", book.ID)
	return nil
}

// Save saves the material book object in the database.
func (book *Book) Save() {
	DB.Set("Book", book.ID, book)
}
