package layout

import (
	"github.com/alecthomas/log4go"
	"html/template"
	"net/http"
)

//this is special,the root of template,not in the tree_cache
//every controller owns one copy
//tree caches the renderd output,if not,render it and save it by mux.lock
//it only affects corresponding key,value of the cache_render
//http://stackoverflow.com/questions/18175630/go-template-executetemplate-include-html
type Application struct {
	t       *template.Template //not cached template
	Content template.HTML      //controller must fill this field to complete render
	UserNav template.HTML      //fill by applicaion itself if render is required
	NavBar  template.HTML

	PageTitle       string
	MetaKeywords    string
	MetaDescription string
}

func NewApp() *Application {
	t, err := template.ParseFiles("view/layout/application.html")
	if err != nil {
		log4go.Error("load template err:%v", err)
		return nil
	}
	return &Application{t: t}
}

//topics controller fill all the member of application obj and call Render to write the execute result to w
func (this *Application) Render(w http.ResponseWriter) {
	if err := this.t.Execute(w, this); err != nil {
		log4go.Error("app execute err:%v", err)
	}
}
