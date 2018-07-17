package main

import (
	"path"
	"strings"

	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/muilab/materia/components/css"
	"github.com/muilab/materia/components/js"
	"github.com/muilab/materia/layout"
	"github.com/muilab/materia/mui"
	"github.com/muilab/materia/pages/home"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Certificate
	app.Security.Load("security/server.crt", "security/server.key")

	// Setup routes
	route(app)

	// API
	mui.API.Install(app)

	// Close the database node on shutdown
	app.OnEnd(mui.Node.Close)

	// Prefetch all collections
	mui.DB.Prefetch()

	// Send "Link" header for Cloudflare on HTML responses
	app.Use(func(ctx *aero.Context, next func()) {
		if !strings.HasPrefix(ctx.URI(), "/_/") && strings.Contains(ctx.Request().Header().Get("Accept"), "text/html") {
			ctx.Response().Header().Set("Link", "</styles>; rel=preload; as=style,</scripts>; rel=preload; as=script")
		}

		next()
	})

	return app
}

func route(app *aero.Application) {
	l := layout.New(app)
	l.Render = fullpage.Render

	// Pages
	l.Page("/", home.Get)

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
