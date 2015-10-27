package common

import (
	"app/fragment"
	"app/fragment/cache"
	"github.com/alecthomas/log4go"
	"html/template"
)

type UserNav struct {
}

//put Lang in cookie(session),then Lang can be selected before register or login
const _user_nav_not_login = `
<ul class="nav user-bar navbar-nav navbar-right">
  <li><a href="/account/sign_up">注册</a></li>
  <li><a href="/account/sign_in">登录</a></li>
</ul>
`

//some jobs of controller was done by cache,rendering ought to be finished by controller
func (this *UserNav) Render(args *fragment.RenderArgs) template.HTML {
	//nil user,return a fixed string
	if args.User == nil {
		//lang can't be access when user == nil
		log4go.Debug("not login")
		return _user_nav_not_login
	}
	key := args.User.UserName
	return cache.Render("common/user_nav", key, args, this)
}

func (this *UserNav) PrepareArgs(args *fragment.RenderArgs) {
	// this shoud take effect next step
	// if args.User == nil {
	// 	user.Load(xxx)
	// }

}
