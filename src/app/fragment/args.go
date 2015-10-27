package fragment

import (
	"app/model"
	"github.com/astaxie/beego/session"
	"net/http"
)

type RenderArgs struct {
	User    *model.User
	Session session.SessionStore
	Req     *http.Request
	Local   map[interface{}]interface{}
}

func NewRenderArgs(s session.SessionStore, r *http.Request, l map[interface{}]interface{}) *RenderArgs {
	return &RenderArgs{
		Session: s,
		Req:     r,
		Local:   l,
	}
}

// called in route.go
func (this *RenderArgs) Release(w http.ResponseWriter) {
	this.Session.SessionRelease(w)
}
