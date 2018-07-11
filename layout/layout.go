package fullpage

import (
	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	return components.Layout(ctx.App, ctx, content)
}
