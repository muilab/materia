package book

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	user := mui.GetUserFromContext(ctx)
	id := ctx.Get("id")
	set, err := mui.GetBook(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material set not found")
	}

	return ctx.HTML(components.Book(set, user))
}

// Edit displays the editing interface.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	set, err := mui.GetBook(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material set not found")
	}

	return ctx.HTML(components.EditBook(set))
}
