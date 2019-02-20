package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func getport() string{
    p := os.Getenv("PORT")

   if p!="" {
      return ":" + p
   }
   return ":8080"
}

func main() {
    port := getport()

	http.HandleFunc("/", idx)
	http.HandleFunc("/link1", lk1)
	http.HandleFunc("/link2", lk2)
	http.HandleFunc("/link3", lk3)
	http.HandleFunc("/link4", lk4)
	http.HandleFunc("/link5", lk5)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(port, nil)

	if err != nil {
	   panic(err)
	}
}



func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}
}

func lk1(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link1 Page",
	}

	err := tpl.ExecuteTemplate(w, "link1.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func lk2(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link2 Page",
	}

	err := tpl.ExecuteTemplate(w, "link2.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func lk3(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link3 Page",
	}

	err := tpl.ExecuteTemplate(w, "link3.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func lk4(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link4 Page",
	}

	err := tpl.ExecuteTemplate(w, "link4.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func lk5(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link5 Page",
	}

	err := tpl.ExecuteTemplate(w, "link5.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
