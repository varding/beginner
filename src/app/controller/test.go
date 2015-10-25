package controller

import (
	"app/cache"
	"app/conf"
	"app/model"
	"fmt"
	"github.com/alecthomas/log4go"
	"net/http"
)

func TestUserNav(w http.ResponseWriter, r *http.Request) {
	log4go.Debug("TestIndex")
	n := &cache.UserNav{}
	u := &model.User{}
	u.UserName = "ding"

	//check lang and give the right locals
	a := &cache.Args{User: u, Local: conf.Local("zh-CN")}
	s := n.Render(a)
	w.Write([]byte(s))
}

func TestApp(w http.ResponseWriter, r *http.Request) {
	app := cache.NewApp()
	user_nav := cache.UserNav{}
	nav_bar := cache.NavBar{}
	//args
	u := &model.User{}
	u.UserName = "ding"

	session, err := session_store.Get(r, "SID") //session id in cookie
	if err != nil {
		log4go.Error("get session err:%v", err)
	} else {
		fmt.Println(session)
		session.Values["lang"] = "zh-CN"
	}

	//check lang and give the right locals
	a := &cache.Args{User: u, Local: conf.Local("zh-CN"), Session: session}
	app.UserNav = user_nav.Render(a)
	app.NavBar = nav_bar.Render(a)

	app.Content = "hello application"
	app.Render(w)
}
