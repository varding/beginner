package route

import (
	"app/conf"
	"app/controller"
	"app/fragment"
	"github.com/alecthomas/log4go"
	"github.com/astaxie/beego/session"
	"net"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

var server *http.Server

func new_render_args(w http.ResponseWriter, r *http.Request) *fragment.RenderArgs {
	//get session,if this is the first time visit ,a new session is returned
	s, err := globalSessions.SessionStart(w, r)
	if err != nil {
		log4go.Error("can't get session:%v", err)
	}

	//default is zh-CN
	lang := "zh-CN"
	if l := s.Get("lang"); l != nil {
		lang = l.(string)
	} else {
		s.Set("lang", "zh-CN") //set default lang
	}

	return fragment.NewRenderArgs(s, r, conf.Local(lang))
}

/*
method=delete是模拟的
http://stackoverflow.com/questions/21739122/rails-delete-method-it-doesnt-work
GET "/patients/1?confirm=Are+you+sure%3F&method=delete"
也就是要先判断query参数里是否有method参数
*/
func http_handler(w http.ResponseWriter, r *http.Request) {
	args := new_render_args(w, r)
	defer args.Release(w) //release sessionStore

	start := time.Now()
	log4go.Debug("req:%s", r.URL.Path)
	if validate_path(w, r) == false {
		return
	}

	//root
	if r.URL.Path == "/" {
		return
	}

	//去掉前后的slash
	p := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(p, "/")

	switch parts[0] {
	case "topics":
		handle_topic(parts, w, r, args)
	case "users":
	case "test":
		//controller.TestUserNav(w, r)
		controller.TestApp(w, r, args)
	}
	d := time.Now().Sub(start)
	log4go.Info("time cost:%s", d.String())
}

func handle_topic(parts []string, w http.ResponseWriter, r *http.Request, args *fragment.RenderArgs) {
	parts_cnt := len(parts)
	m := r.Method
	log4go.Debug("handle topics,path parts:%q", parts)
	switch parts_cnt {
	case 1:
		switch m {
		case "GET":
			controller.TopicIndex(w, r, args)
		case "POST":
			controller.TopicCreate(w, r, args)
		}
	case 2:
		if parts[1] == "new" {
			controller.TopicNew(w, r, args)
		} else {
			id, err := strconv.ParseUint(parts[1], 10, 64)
			if err != nil {
				log4go.Error("topic id parse err:", err)
				//redirect server error
				return
			}
			switch m {
			case "GET":
				controller.TopicShow(uint32(id), w, r, args)
			case "POST":
				controller.TopicUpdate(uint32(id), w, r, args)
			case "DELETE":
				controller.TopicDelete(uint32(id), w, r, args)
			}
		}
	case 3:
		if parts[1] == "edit" {
			id, err := strconv.ParseUint(parts[3], 10, 64)
			if err != nil {
				log4go.Error("topic id parse err:", err)
				//redirect server error
				return
			}
			controller.TopicEdit(uint32(id), w, r, args)
		}
	}
}

var globalSessions *session.Manager

func Run() {
	var err1 error
	globalSessions, err1 = session.NewManager("file", `{"cookieName":"sid","gclifetime":3600,"ProviderConfig":"./tmp"}`)
	if err1 != nil {
		log4go.Error("session manager init failed")
	}

	go globalSessions.GC()

	server = &http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      http.HandlerFunc(http_handler),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	//https://github.com/revel/revel/blob/master/server.go
	listener, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log4go.Critical("Failed to listen:", err)
	}
	log4go.Info("server running :9090")
	log4go.Critical("Failed to serve:", server.Serve(listener))
}

func validate_path(w http.ResponseWriter, req *http.Request) bool {
	//https://github.com/gorilla/mux/blob/master/mux.go
	if p := cleanPath(req.URL.Path); p != req.URL.Path {

		// Added 3 lines (Philip Schlump) - It was dropping the query string and #whatever from query.
		// This matches with fix in go 1.2 r.c. 4 for same problem.  Go Issue:
		// http://code.google.com/p/go/issues/detail?id=5252
		url := *req.URL
		url.Path = p
		p = url.String()

		w.Header().Set("Location", p)
		w.WriteHeader(http.StatusMovedPermanently)
		return false
	}
	return true
}

func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}
