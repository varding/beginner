package controller

import (
	"app/fragment"
	"app/fragment/common"
	"app/fragment/layout"
	"app/model"
	"github.com/alecthomas/log4go"
	"net/http"
)

func TestUserNav(w http.ResponseWriter, r *http.Request, args *fragment.RenderArgs) {
	log4go.Debug("TestIndex")
	n := &common.UserNav{}
	s := n.Render(args)
	w.Write([]byte(s))
}

func TestApp(w http.ResponseWriter, r *http.Request, args *fragment.RenderArgs) {
	app := layout.NewApp()
	user_nav := common.UserNav{}
	nav_bar := layout.NavBar{}
	//args
	u := &model.User{}
	u.UserName = "ding"

	// session, err := session_store.Get(r, "SID") //session id in cookie
	// if err != nil {
	// 	log4go.Error("get session err:%v", err)
	// } else {
	// 	fmt.Println(session)
	// 	session.Values["lang"] = "zh-CN"
	// }

	//check lang and give the right locals
	//a := &cache.RenderArgs{User: u, Local: conf.Local("zh-CN"), Session: session}
	app.UserNav = user_nav.Render(args)
	app.NavBar = nav_bar.Render(args)

	app.Content = "hello application"
	app.Render(w)
}
