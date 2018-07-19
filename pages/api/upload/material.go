package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/mui"
)

// MaterialImage is the endpoint for uploading material images.
func MaterialImage(ctx *aero.Context) string {
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, err)
	}

	body, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	material.SetImageBytes(body)
	material.Save()
	return ctx.Text("ok")
}

// MaterialSampleImage is the endpoint for uploading material sample images.
func MaterialSampleImage(ctx *aero.Context) string {
	id := ctx.Get("id")
	material, err := mui.GetMaterial(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, err)
	}

	body, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	material.AddSampleImageBytes(body)
	material.Save()
	return ctx.Text("ok")
}
