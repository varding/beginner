package controller

import (
	"app/model"
	"github.com/alecthomas/log4go"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

type BaseController struct {
	Session *sessions.Session //
	User    *model.User
}

var session_store = sessions.NewCookieStore([]byte("123467855"))

//先获取session
func (this *BaseController) Load(req *http.Request) bool {
	var err error
	this.Session, err = session_store.Get(req, "SID") //session id in cookie
	if err != nil {
		log4go.Error("session load err:", err)
		return false
	}

	if this.Session.IsNew {
		this.Session.Options.Domain = "www.beginner-mind.com"
		// MaxAge=0 means no 'Max-Age' attribute specified.
		// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
		// MaxAge>0 means Max-Age attribute present and given in seconds.
		this.Session.Options.MaxAge = 7 * 24 * 3600 //seven days
		this.Session.Options.HttpOnly = false
		this.Session.Options.Secure = true
	}

	return true
}

//load basic infomation of user
func (this *BaseController) LoadUser() {
	if userId, ok := this.Session.Values["uid"]; ok {
		this.User = model.LoadUser(userId.(string))
	}
}

//load all infomation of current user
func (this *BaseController) LoadUserAll() {
	if userId, ok := this.Session.Values["uid"]; ok {
		this.User = model.LoadUserAll(userId.(string))
	}
}

// all redirections
func (this *BaseController) Redirect(t *template.Template, msg string) {

}

func (this *BaseController) Redirect404(msg string) {

}

func (this *BaseController) RedirectFlash(t *template.Template, msg string) {

}
