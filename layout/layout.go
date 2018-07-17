package fullpage

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	user := mui.GetUserFromContext(ctx)
	return components.Layout(ctx.App, ctx, content, user)
}
