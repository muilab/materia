package middleware

import (
	"strings"

	"github.com/aerogo/aero"
)

// LinkHeader sends the headers required for HTTP/2 push by proxies like Cloudflare.
func LinkHeader(ctx *aero.Context, next func()) {
	if !strings.HasPrefix(ctx.URI(), "/_/") && strings.Contains(ctx.Request().Header().Get("Accept"), "text/html") {
		ctx.Response().Header().Set("Link", "</styles>; rel=preload; as=style,</scripts>; rel=preload; as=script")
	}

	next()
}
