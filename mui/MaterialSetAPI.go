package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
	"github.com/muilab/materia/mui/utils"
)

// Force interface implementations
var (
	_ api.Newable  = (*MaterialSet)(nil)
	_ api.Editable = (*MaterialSet)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (set *MaterialSet) Authorize(ctx *aero.Context, action string) error {
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
func (set *MaterialSet) Create(ctx *aero.Context) error {
	set.ID = GenerateID("MaterialSet")
	set.Name = "Untitled"
	set.IsDraft = true
	set.Created = utils.DateTimeUTC()
	set.CreatedBy = GetUserFromContext(ctx).ID
	set.MaterialIDs = []string{}
	return nil
}

// Save saves the material set object in the database.
func (set *MaterialSet) Save() {
	DB.Set("MaterialSet", set.ID, set)
}
