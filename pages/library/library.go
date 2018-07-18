package library

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the frontpage.
func Get(ctx *aero.Context) string {
	user := mui.GetUserFromContext(ctx)
	allMaterialSets := mui.AllMaterialSets()
	return ctx.HTML(components.Library(allMaterialSets, user))
}
