package cache

import (
	"bytes"
	"github.com/alecthomas/log4go"
	"html/template"
	"sync"
	"time"
)

type StringCache struct {
	data       string    //cached content
	last_visit time.Time //last visit time
	hot        int       //visit cnt since last time
	lock       sync.RWMutex
}

type TemplateCache struct {
	path   string                  //template path
	t      *template.Template      //parsed template
	cached map[string]*StringCache //cached result of render output,a timeout algrithm must be used
	lock   sync.RWMutex            //lock before visit or modify cached
}

func NewTemplate(path string) *TemplateCache {
	t, err := template.ParseFiles(path)
	if err != nil {
		log4go.Error("parse template path:%s,err:%v", path, err)
		return nil
	}
	return &TemplateCache{t: t, cached: make(map[string]*StringCache), path: path}
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

//if cache hited,no args is needed,this could save loading time(from database)
//the hited string could also be empty,so use bool indicate hit or not
func (this *TemplateCache) IsHit(key string) (string, bool) {
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
	return "", false
}

// is it possible that two goroutine update the same cache??how to prevent this sutation??
func (this *TemplateCache) Save(key, value string) {
	s := &StringCache{data: value, last_visit: time.Now()}
	this.lock.Lock()
	this.cached[key] = s
	this.lock.Unlock()
}

//if not hit then arg must be prepared for render
func (this *TemplateCache) Render(key string, arg interface{}) string {
	buf := bytes.NewBuffer(nil)
	//bufio.NewWriter(w)
	//how to convert to io.Writer?
	//render and save the render results

	if this.t == nil {
		log4go.Error("template nil")
		return ""
	}

	this.t.Execute(buf, arg)
	//s := string(buf)
	//save the result to cache
	s := buf.String()
	this.Save(key, s)
	return s
}

// func new(path string) *CacheTemplate {
// 	return &CacheTemplate{path: path}
// }

// //path:template path
// //key :cache key,for example content of topics/100=>content_topics_100
// //or reply number of 10 of topics/100,reply_10_topics_100
// func Get(path, key string) *CacheTemplate {
// 	// if t, ok := template_tree[path]; ok {
// 	// 	return t
// 	// }
// 	if c, ok := template_tree[path]; !ok {
// 		return nil
// 	} else {
// 		if c.Check(key) {
// 			c.String(key)
// 		} else {

// 		}
// 	}
// }

// func Save(path, key string, args map[string]string) *CacheTemplate {

// }

// //
var template_tree map[string]*TemplateCache

// func Walk() *CacheTemplate {
// 	//walk template directory and format tree architecture
// }

// func (this *CacheTemplate) Check(key string) bool {
// 	//cache ok
// 	_, ok := this.cache_render[key]
// 	return ok
// }

// func (this *CacheTemplate) String(key string) string {
// 	//render
// 	if this.t != nil {
// 		this.t.Execute(wr, data)
// 	}
// }

func init() {
	template_tree = make(map[string]*TemplateCache)
	template_tree["common/user_nav"] = NewTemplate("view/common/user_nav.html")
}
