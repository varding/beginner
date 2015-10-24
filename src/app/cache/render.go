package cache

import (
	//"github.com/alecthomas/log4go"
	"app/model"
	"github.com/gorilla/sessions"
	"net/http"
)

//architecture is almost the same with view

//all template must implement this
type RenderCache interface {
	//IsHit(key string) bool  //check if the key is cached
	//String() string         //output renderd content,if the cached is presented then return directly,otherwise render it
	Render(*Args) string
	Local(id string) string //output cached field of obj
}

//
var cache_tree map[string]RenderCache

func init() {
	cache_tree = make(map[string]RenderCache)
	//cache_tree["app"] = NewApp()
}

type Args struct {
	User    *model.User
	Session *sessions.Session
	Req     *http.Request
}

////////////////////////////////
//find cache path and return string
//this should insert into func map
//two parameters,the first is the current render args obj,sencond is the path
//if starts with '#',it is the field of obj
//else it is the global cached template
func GlobalCache(path string) string {
	// if len(path) == 0 {
	// 	log4go.Error("empty path")
	// 	return ""
	// }

	// // //combined content
	// // if path[0] == '#' {

	// // }

	// if c, ok := cache_tree[path]; ok {
	// 	return c.String()
	// } else {
	// 	log4go.Error("can't find cache of path:", path)
	// 	return ""
	// }
	return ""
}

//get the cached field with name id
func LocalCache(c RenderCache, id string) {

}
