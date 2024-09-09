package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Root(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "404 page not found")
        return
    }
	
	index, err := template.ParseFiles("index.html", "templates/home.html", "templates/navbar.html")
	if err != nil {
		log.Fatal(err) 
	} 
	data := make(map[string]string)
	data["title"] = "Jedders | Home"
	index.Execute(w, data)
}