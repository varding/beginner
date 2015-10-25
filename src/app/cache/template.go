package cache

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/log4go"
	"html/template"
	"sync"
	"time"
)

//add lang_id as part of the key,so different lang cached respectively
func CacheRender(path string, k string, args *Args, r ArgReader) template.HTML {
	tpl, ok := template_tree[path]
	if !ok {
		log4go.Error("can't find template cache")
		return ""
	}

	lang := "0"
	if args.Session != nil {
		l, ok := args.Session.Values["lang"]
		if ok {
			lang = l.(string)
		}
	}

	key := fmt.Sprintf("%s%s", lang, k)
	return tpl.Render(key, args, r)
}

/////////////////////////////////////////
//template.HTML will not be escaped in the template
type HtmlFragmentCache struct {
	data       template.HTML //cached content
	last_visit time.Time     //last visit time
	hot        int           //visit cnt since last time
	lock       sync.RWMutex
}

type TemplateCache struct {
	path   string                        //template path
	t      *template.Template            //parsed template
	cached map[string]*HtmlFragmentCache //cached result of render output,a timeout algrithm must be used
	lock   sync.RWMutex                  //lock before visit or modify cached
}

func new_template(path string) *TemplateCache {
	t, err := template.ParseFiles(path)
	if err != nil {
		log4go.Error("parse template path:%s,err:%v", path, err)
		return nil
	}
	return &TemplateCache{t: t, cached: make(map[string]*HtmlFragmentCache), path: path}
}

//if cache hited,no args is needed,this could save loading time(from database)
//the hit string could also be empty,so use bool indicate hit or not
func (this *TemplateCache) is_hit(key string) (template.HTML, bool) {
	//the whole map was locked?
	this.lock.RLock()
	c, ok := this.cached[key]
	this.lock.RUnlock()
	if ok {
		log4go.Debug("cache hit!,key:%s", key)
		c.lock.Lock()
		defer c.lock.Unlock()
		c.hot += 1
		c.last_visit = time.Now()
		return c.data, true
	}
	return template.HTML(""), false
}

// is it possible that two goroutine update the same cache??how to prevent this sutation??
func (this *TemplateCache) save(key string, value template.HTML) {
	s := &HtmlFragmentCache{data: value, last_visit: time.Now()}
	this.lock.Lock()
	this.cached[key] = s
	this.lock.Unlock()
}

//if not hit then arg must be prepared for render
func (this *TemplateCache) Render(key string, args *Args, r ArgReader) template.HTML {
	//check if hit
	if s, ok := this.is_hit(key); ok {
		return s
	}

	//if not hit,render it
	log4go.Debug("render")

	buf := bytes.NewBuffer(nil)
	//render and save the render results

	if this.t == nil {
		log4go.Error("template nil")
		return ""
	}

	r.ReadArgs(args)

	this.t.Execute(buf, args)
	//s := string(buf)
	//save the result to cache
	s := template.HTML(buf.String())
	this.save(key, s)
	return s
}

func (this *TemplateCache) CheckOutOfDate(t time.Time) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for k, v := range this.cached {
		if v.last_visit.Before(t) {
			delete(this.cached, k)
		}
	}
}

var template_tree map[string]*TemplateCache

func init() {
	template_tree = make(map[string]*TemplateCache)
	template_tree["common/user_nav"] = new_template("view/common/user_nav.html")
	template_tree["common/nav_bar"] = new_template("view/common/nav_bar.html")
	//template_tree["layouts/application"] = new_template("view/layouts/application.html")
}
