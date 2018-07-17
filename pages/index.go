package pages

import (
	"path"

	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/muilab/materia/components/css"
	"github.com/muilab/materia/components/js"
	"github.com/muilab/materia/layout"
	"github.com/muilab/materia/pages/contact"
	"github.com/muilab/materia/pages/home"
	"github.com/muilab/materia/pages/pricing"
)

// Install configures the application routes.
func Install(app *aero.Application) {
	l := layout.New(app)
	l.Render = fullpage.Render

	// Pages
	l.Page("/", home.Get)
	l.Page("/pricing", pricing.Get)
	l.Page("/contact", contact.Get)

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
