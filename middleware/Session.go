package middleware

import "github.com/aerogo/aero"

// Session writes the session data to the store if it has been modified.
func Session(ctx *aero.Context, next func()) {
	// Handle the request first
	next()

	// Update session if it has been modified
	if ctx.HasSession() && ctx.Session().Modified() {
		ctx.App.Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
	}
}
