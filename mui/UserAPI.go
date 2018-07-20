package mui

import (
	"errors"

	"github.com/aerogo/aero"
	"github.com/aerogo/api"
)

// Force interface implementations
var (
	_ api.Editable = (*User)(nil)
)

// Authorize returns an error if the given API POST request is not authorized.
func (user *User) Authorize(ctx *aero.Context, action string) error {
	editor := GetUserFromContext(ctx)

	if editor == nil {
		return errors.New("Not logged in")
	}

	if action == "edit" && editor.ID != user.ID {
		return errors.New("Can't edit data from other users")
	}

	return nil
}

// Save saves the user object in the database.
func (user *User) Save() {
	DB.Set("User", user.ID, user)
}
