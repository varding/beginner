package layout

import (
	"app/fragment"
	"app/fragment/cache"
	"html/template"
)

type NavBar struct {
}

func (this *NavBar) Render(args *fragment.RenderArgs) template.HTML {
	return cache.Render("layout/nav_bar", "", args, this)
}

//nothing need to fill
func (this *NavBar) PrepareArgs(args *fragment.RenderArgs) {
}

//all cached template should have a lang speciafication
