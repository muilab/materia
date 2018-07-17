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
	user, err := mui.GetUser(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found")
	}

	return ctx.HTML(components.User(user))
}
