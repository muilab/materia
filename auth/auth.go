package auth

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/mui"
)

const newUserStartRoute = "/"

// Install installs the authentication routes in the application.
func Install(app *aero.Application) {
	// Logout
	app.Get("/logout", func(ctx *aero.Context) string {
		if ctx.HasSession() {
			user := mui.GetUserFromContext(ctx)

			if user != nil {
				authLog.Info("User logged out", user.ID, ctx.RealIP(), user.Email, user.RealName())
			}

			ctx.Session().Set("userId", nil)
		}

		return ctx.Redirect("/")
	})
}
