package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/muilab/materia/mui/utils"
)

// Force interface implementations
var (
	_ api.Newable   = (*Material)(nil)
	_ api.Editable  = (*Material)(nil)
	_ api.Deletable = (*Material)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (material *Material) Authorize(ctx *aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if action == "edit" && user.ID != material.CreatedBy {
		return errors.New("Can't edit data from other users")
	}

	return nil
}

// Create sets the data for a new material set with data we received from the API request.
func (material *Material) Create(ctx *aero.Context) error {
	material.ID = GenerateID("Material")
	material.Created = utils.DateTimeUTC()
	material.CreatedBy = GetUserFromContext(ctx).ID
	return nil
}

// DeleteInContext deletes the material in the given context.
func (material *Material) DeleteInContext(ctx *aero.Context) error {
	return material.Delete()
}

// Delete deletes the material from the database.
func (material *Material) Delete() error {
	// Delete the material from all the books that contained it
	for book := range StreamBooks() {
		if book.Remove(material.ID) {
			book.Save()
		}
	}

	// Delete all samples for this material
	for _, sample := range material.Samples() {
		sample.Delete()
	}

	// Delete all image files
	for _, output := range materialImageOutputs {
		output.Delete(material.ID)
	}

	DB.Delete("Material", material.ID)
	return nil
}

// Save saves the material set object in the database.
func (material *Material) Save() {
	DB.Set("Material", material.ID, material)
}
