package user

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := mui.GetUserFromContext(ctx)
	viewUser, err := mui.GetUser(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found")
	}

	return ctx.HTML(components.User(viewUser, user))
}

// Edit the user.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := mui.GetUserFromContext(ctx)
	viewUser, err := mui.GetUser(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found")
	}

	if user != viewUser {
		return ctx.Error(http.StatusUnauthorized, "Can't edit data from other users")
	}

	return ctx.HTML(components.EditUser(viewUser))
}
