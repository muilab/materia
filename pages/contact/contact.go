package contact

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
)

// Get the page.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Contact())
}
