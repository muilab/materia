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
	allMaterialSets := mui.AllMaterialSets()

	sort.Slice(allMaterialSets, func(i, j int) bool {
		return allMaterialSets[i].Created > allMaterialSets[j].Created
	})

	return ctx.HTML(components.Library(allMaterialSets, user))
}
