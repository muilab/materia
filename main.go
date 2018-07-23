package main

import (
	"os"

	"github.com/aerogo/aero"

	"github.com/aerogo/session-store-nano"
	"github.com/muilab/materia/auth"
	"github.com/muilab/materia/middleware"
	"github.com/muilab/materia/mui"
	"github.com/muilab/materia/pages"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Sessions
	app.Sessions.Duration = 3600 * 24 * 30 * 6
	app.Sessions.Store = nanostore.New(mui.DB.Collection("Session"))

	// Certificate
	app.Security.Load("security/server.crt", "security/server.key")

	// Content security policy
	app.ContentSecurityPolicy.Set("img-src", "https: data:")

	// Development config
	production := os.Getenv("PRODUCTION")

	if production == "" {
		app.Config.Domain = "beta.materia.mui.jp"
	}

	// Pages
	pages.Install(app)

	// Auth
	auth.Install(app)

	// API
	mui.API.Install(app)

	// Rewrite
	app.Rewrite(rewrite)

	// Middleware
	app.Use(
		middleware.Session,
		middleware.LinkHeader,
	)

	// Close the database node on shutdown
	app.OnEnd(mui.Node.Close)

	// Prefetch all collections
	mui.DB.Prefetch()

	return app
}
