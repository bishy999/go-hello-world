package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

const msg string = "hello world from my simple golang webapp running on a docker container"

type response struct {
	Status   string
	Message string
	Lyrics  []string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", msg)
}

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v1 := response{
		Status:  "ok",
		Message: msg,
		Lyrics:  []string{"You say goodbey and i say hello", "i am the eggman coo coo ca choo"},
	}
	j, err := json.Marshal(v1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}
