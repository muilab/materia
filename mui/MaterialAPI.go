package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/muilab/materia/mui/utils"
)

// Force interface implementations
var (
	_ api.Newable  = (*Material)(nil)
	_ api.Editable = (*Material)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (set *Material) Authorize(ctx *aero.Context, action string) error {
	user := GetUserFromContext(ctx)

	if user == nil {
		return errors.New("Not logged in")
	}

	if action == "edit" && user.ID != set.CreatedBy {
		return errors.New("Can't edit data from other users")
	}

	return nil
}

// Create sets the data for a new material set with data we received from the API request.
func (set *Material) Create(ctx *aero.Context) error {
	set.ID = GenerateID("Material")
	set.Name = "Untitled"
	set.Created = utils.DateTimeUTC()
	set.CreatedBy = GetUserFromContext(ctx).ID
	return nil
}

// Save saves the material set object in the database.
func (set *Material) Save() {
	DB.Set("Material", set.ID, set)
}
