package books

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	books := mui.AllBooks()
	return ctx.HTML(components.Books(books))
}
