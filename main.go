package main

import (
	"strings"

	"github.com/aerogo/aero"

	"github.com/muilab/materia/mui"
	"github.com/muilab/materia/pages"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Certificate
	app.Security.Load("security/server.crt", "security/server.key")

	// Setup routes
	pages.Install(app)

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
