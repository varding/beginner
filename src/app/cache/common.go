package cache

import (
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
func (this *UserNav) Render(args *Args) template.HTML {
	//nil user,return a fixed string
	if args.User == nil {
		//lang can't be access when user == nil
		log4go.Debug("not login")
		return _user_nav_not_login
	}
	key := args.User.UserName
	return CacheRender("common/user_nav", key, args, this)
}

func (this *UserNav) ReadArgs(args *Args) {
	// this shoud take effect next step
	// if args.User == nil {
	// 	user.Load(xxx)
	// }

}

//////////////////////////

type NavBar struct {
}

func (this *NavBar) Render(args *Args) template.HTML {
	return CacheRender("common/user_nav", "", args, this)
}

//nothing need to fill
func (this *NavBar) ReadArgs(args *Args) {
}

//all cached template should have a lang speciafication
