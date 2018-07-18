package library

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the frontpage.
func Get(ctx *aero.Context) string {
	user := mui.GetUserFromContext(ctx)
	allMaterials := mui.AllMaterials()

	sort.Slice(allMaterials, func(i, j int) bool {
		return allMaterials[i].Created > allMaterials[j].Created
	})

	return ctx.HTML(components.Library(allMaterials, user))
}
