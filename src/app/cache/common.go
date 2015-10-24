package cache

import (
	//"app/model"
	//"bytes"
	"github.com/alecthomas/log4go"
	//"io"
	//"time"
)

type UserNav struct {
	//c    *cache.CacheTemplate
	//User *model.User //every template need it't own objs,but how to pass paramters like userId?
	//all sessions use the same cache?clone a template each for them?
}

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

	//nil user,return a fixed string
	if args.User == nil {
		//user == nil,so lang can't be access
		//put Lang in cookie(session),then Lang can be selected before register or login
		log4go.Debug("not login")
		return _user_nav_not_login
	}

	//check if hit
	k := this.key(args.User.UserName)
	if s, ok := tpl.IsHit(k); ok {
		return s
	}

	//if not hit,render it
	log4go.Debug("render")
	return tpl.Render(k, args)
}

func (this *UserNav) key(user_name string) string {
	return user_name
}
