package controller

import (
	"github.com/alecthomas/log4go"
	"html/template"
	"net/http"
)

func TopicIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./view/topics/index.html", "./view/layouts/application.html")
	if err != nil {
		log4go.Error("can't load template:%v", err)
		return
	}
	err2 := t.Execute(w, nil)
	if err2 != nil {
		log4go.Error("render err:%v", err2)
	}
}

func TopicCreate(w http.ResponseWriter, r *http.Request) {

}

func TopicDelete(id uint32, w http.ResponseWriter, r *http.Request) {

}

func TopicEdit(id uint32, w http.ResponseWriter, r *http.Request) {

}
func TopicNew(w http.ResponseWriter, r *http.Request) {

}

func TopicShow(id uint32, w http.ResponseWriter, r *http.Request) {

}

func TopicUpdate(id uint32, w http.ResponseWriter, r *http.Request) {

}
