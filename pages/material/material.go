package material

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	return ctx.HTML(components.Material(material))
}

// Edit displays the editing interface.
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	return ctx.HTML(components.EditMaterial(material))
}
