package cache

import (
	"app/model"
	"github.com/alecthomas/log4go"
	"time"
)

type UserNav struct {
	//c    *cache.CacheTemplate
	User *model.User //every template need it't own objs,but how to pass paramters like userId?
	//all sessions use the same cache?clone a template each for them?
}

//some jobs of controller was done by cache,rendering ought to be finished by controller
func (this *UserNav) Render(args *Args) string {
	t, ok := template_tree["common/user_nav"]
	if !ok {
		log4go.Error("can't find template cache")
		return ""
	}

	//nil user,return a fixed string
	if args.User == nil {
		return `<li><%= link_to( t("common.register"), new_user_registration_path) %></li>
  <li><%= link_to( t("common.login"), new_user_session_path ) %></li>
	`
	}

	//check if cache hits
	if c := t.IsHit(this.key(args.User.UserName)); c != nil {
		c.lock.RLock()
		defer c.lock.RUnlock()
		c.hot += 1
		c.last_visit = time.Now()
		return c.data
	}

	//render and save the render results

	return ""
}

func (this *UserNav) key(user_name string) string {
	return ""
}
