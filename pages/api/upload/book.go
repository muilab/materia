package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/mui"
)

// BookImage is the endpoint for uploading book images.
func BookImage(ctx *aero.Context) string {
	id := ctx.Get("id")
	book, err := mui.GetBook(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, err)
	}

	body, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusBadRequest, err)
	}

	book.SetImageBytes(body)
	book.Save()

	return ctx.Text("ok")
}
