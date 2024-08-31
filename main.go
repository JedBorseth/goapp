package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)


func main() { 
	go tailwindMinify()
    http.HandleFunc("/", root)
    http.HandleFunc("/css", css)
    http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)
	fmt.Println("Server started at port 8090")
	
	dir := "public"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Failed to read", dir, err)
	}
	for _, file := range files {
		if(!file.IsDir()){
			filePath := filepath.Join(dir, file.Name())
			fmt.Println("Serving", filePath)
			http.HandleFunc("/public/" + file.Name(), func(w http.ResponseWriter, req *http.Request) {
				http.ServeFile(w, req, filePath)
			})

	}
}
	


    http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
        errorHandler(w, req, http.StatusNotFound)
        return
    }
	index := template.Must(template.ParseFiles("index.html"))
	index.Execute(w, nil)
}
func login(w http.ResponseWriter, req *http.Request) {
	// index := template.Must(template.ParseFiles("login.html"))
	// index.Execute(w, nil)
	fmt.Fprint(w, "login")
}
func css(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "styles/output.css")
}
func errorHandler(w http.ResponseWriter, req *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
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

