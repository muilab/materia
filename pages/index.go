package pages

import (
	"path"

	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/muilab/materia/components/css"
	"github.com/muilab/materia/components/js"
	"github.com/muilab/materia/layout"
	"github.com/muilab/materia/pages/about"
	"github.com/muilab/materia/pages/api/upload"
	"github.com/muilab/materia/pages/books"
	"github.com/muilab/materia/pages/contact"
	"github.com/muilab/materia/pages/library"
	"github.com/muilab/materia/pages/login"
	"github.com/muilab/materia/pages/material"
	"github.com/muilab/materia/pages/materialset"
	"github.com/muilab/materia/pages/pricing"
	"github.com/muilab/materia/pages/report"
	"github.com/muilab/materia/pages/user"
	"github.com/muilab/materia/pages/workshop"
)

// Install configures the application routes.
func Install(app *aero.Application) {
	l := layout.New(app)
	l.Render = fullpage.Render

	// Pages
	l.Page("/", library.Get)
	l.Page("/books", books.Get)
	l.Page("/workshop", workshop.Get)
	l.Page("/report", report.Get)
	l.Page("/pricing", pricing.Get)
	l.Page("/contact", contact.Get)
	l.Page("/about", about.Get)
	l.Page("/login", login.Get)
	l.Page("/user/:id", user.Get)
	l.Page("/material/:id", material.Get)
	l.Page("/material/:id/edit", material.Edit)
	l.Page("/materialset/:id", materialset.Get)
	l.Page("/materialset/:id/edit", materialset.Edit)

	// API
	app.Get("/material/:id/download", material.Download)
	app.Post("/api/upload/material/:id/image", upload.MaterialImage)
	app.Post("/api/upload/material/:id/sample", upload.MaterialSampleImage)

	// Script bundle
	scriptBundle := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	// CSS bundle
	cssBundle := css.Bundle()

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(cssBundle)
	})

	// Static files
	app.Get("/images/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File(path.Join("images", ctx.Get("file")))
	})

	// Manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})
}
