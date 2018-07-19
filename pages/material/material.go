package material

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Get the page.
func Get(ctx *aero.Context) string {
	user := mui.GetUserFromContext(ctx)
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	return ctx.HTML(components.Material(material, user))
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

// Download downloads the material files.
func Download(ctx *aero.Context) string {
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Material not found")
	}

	if !material.HasImage() {
		return ctx.Error(http.StatusBadRequest, "Material has no image")
	}

	ctx.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s%s"`, material.ID, material.Image.Extension))
	return ctx.File(fmt.Sprintf("images/materials/original/%s%s", material.ID, material.Image.Extension))
}
