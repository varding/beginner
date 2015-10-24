package cache

import (
	//"cache"
	"github.com/alecthomas/log4go"
	tpl "html/template"
	//"net/http"
)

//this is special,the root of template,not in the tree_cache
//every controller owns one copy
//tree caches the renderd output,if not,render it and save it by mux.lock
//it only affects corresponding key,value of the cache_render
type Application struct {
	t       *tpl.Template //not cached template
	Content string        //controller must fill this field to complete render
	UserNav string        //fill by applicaion itself if render is required
}

func NewApp() *Application {
	t, err := tpl.ParseFiles("view/layouts/application.html")
	if err != nil {
		log4go.Error("can't find template:%v", err)
		return nil
	}
	return &Application{t: t}
}

func (this *Application) Render(args *Args) {

}
