package library

import (
	"sort"
	"strings"

	"github.com/aerogo/aero"
	"github.com/muilab/materia/components"
	"github.com/muilab/materia/mui"
)

// Search the material library.
func Search(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.ToLower(term)
	user := mui.GetUserFromContext(ctx)

	materials := mui.FilterMaterials(func(material *mui.Material) bool {
		if term == "" {
			return true
		}

		if strings.Contains(strings.ToLower(material.Name), term) {
			return true
		}

		if strings.Contains(strings.ToLower(material.Description), term) {
			return true
		}

		return false
	})

	sort.Slice(materials, func(i, j int) bool {
		return materials[i].Created > materials[j].Created
	})

	return ctx.HTML(components.Library(materials, user))
}
