package cache

import (
	"github.com/alecthomas/log4go"
)

type UserNav struct {
	args *Args
}

//put Lang in cookie(session),then Lang can be selected before register or login
const _user_nav_not_login = `
<ul class="nav user-bar navbar-nav navbar-right">
  <li><a href="/account/sign_up">注册</a></li>
  <li><a href="/account/sign_in">登录</a></li>
</ul>
`

//some jobs of controller was done by cache,rendering ought to be finished by controller
func (this *UserNav) Render(args *Args) string {
	tpl, ok := template_tree["common/user_nav"]
	if !ok {
		log4go.Error("can't find template cache")
		return ""
	}
	this.args = args //save args for ReadArgs

	//nil user,return a fixed string
	if args.User == nil {
		//lang can't be access when user == nil
		log4go.Debug("not login")
		return _user_nav_not_login
	}

	k := this.key(args.User.UserName)
	return tpl.Render(k, this)
}

func (this *UserNav) ReadArgs() *Args {
	return this.args
}

func (this *UserNav) key(user_name string) string {
	return user_name
}
