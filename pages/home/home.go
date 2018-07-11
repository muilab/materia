package home

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
)

// Get the frontpage.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Home())
}
