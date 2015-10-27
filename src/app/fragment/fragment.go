package fragment

import (
	"html/template"
)

/*
how to create a fragment with cache
e.g view/common/user_nav.html
1. define a struct UserNav in cache/common package
2. impl Render
this function will be called by parent fragment or controller to get output html fragment
func (this *UserNav) Render(args *cache.RenderArgs) template.HTML {
	//nil user,return a fixed string
	if args.User == nil {
		//lang can't be access when user == nil
		log4go.Debug("not login")
		return _user_nav_not_login
	}
	key := args.User.UserName
	return cache.CacheRender("common/user_nav", key, args, this)
}
3. impl PrepareArgs
func (this *UserNav) PrepareArgs(args *cache.RenderArgs) {
	// this shoud take effect next step
	// if args.User == nil {
	// 	user.Load(xxx)
	// }
}
4. add templatecache entry in cache/template.go
template_tree["common/user_nav"] = new_template("view/common/user_nav.html")
*/

//////////////////////////////
// all fragment must impl this,if cache miss cache.CacheRender call this function to prepare args for real render
//called before render in templateCache,callee should prepare data for render(cache miss)
type PrepareRenderArgs interface {
	PrepareArgs(*RenderArgs)
}

//all fragment must impl this,parent fragment or controller call this to get output html
type Fragment interface {
	Render(args *RenderArgs) template.HTML
	//PrepareArgs(*RenderArgs)
}
