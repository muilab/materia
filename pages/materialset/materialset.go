package materialset

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	set, err := mui.GetMaterialSet(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	return ctx.HTML(components.MaterialSet(set))
}

// Edit displays the editing interface.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	set, err := mui.GetMaterialSet(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	return ctx.HTML(components.EditMaterialSet(set))
}
