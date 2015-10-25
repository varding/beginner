package cache

import (
	"app/model"
	"github.com/gorilla/sessions"
	"net/http"
)

type Args struct {
	User    *model.User
	Session *sessions.Session
	Req     *http.Request
	Local   map[interface{}]interface{}
}

//////////////////////////////
//called before render in templateCache,callee should prepare data for render(cache miss)
type ArgReader interface {
	ReadArgs() *Args
}
