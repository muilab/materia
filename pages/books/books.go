package books

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	user := mui.GetUserFromContext(ctx)

	books := mui.FilterBooks(func(book *mui.Book) bool {
		if user != nil && user.ID == book.CreatedBy {
			return true
		}

		return book.Public
	})

	sort.Slice(books, func(i, j int) bool {
		return books[i].Created > books[j].Created
	})

	return ctx.HTML(components.Books(books, user))
}
