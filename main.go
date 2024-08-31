package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"text/template"
)

func root(w http.ResponseWriter, req *http.Request) {
	index := template.Must(template.ParseFiles("index.html"))
	index.Execute(w, nil)
}
func css(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "styles/output.css")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
} 



func tailwindMinify() {
	cmd := exec.Command("./tailwindcss", "-i", "styles/input.css", "-o", "styles/output.css", "--minify")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}

func main() { 
	fmt.Println("Server started at port 8090")
	tailwindMinify()
	
    http.HandleFunc("/", root)

    http.HandleFunc("/css", css)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}