package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)


func main() { 
	port := "8080"
	go tailwindMinify()
    http.HandleFunc("/", root)
    http.HandleFunc("/css", css)
    http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)
	fmt.Println("Server started at port" + port)
	
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
	


    http.ListenAndServe(":" + port, nil)
} 

func root(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
        errorHandler(w, req, http.StatusNotFound)
        return
    }
	
	index, err := template.ParseFiles("index.html", "templates/home.html")
	if err != nil {
		log.Fatal(err) 
	} 
	data := make(map[string]string)
	data["title"] = "Jedders | Home"
	index.Execute(w, data)
}
func login(w http.ResponseWriter, req *http.Request) {
	if(req.Method == "POST"){
	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	fmt.Println(username)
	fmt.Println(password)




	// database storing session ids and whatnot
		
	}


	index, err := template.ParseFiles("index.html", "templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]string)
	data["title"] = "Jedders | Login"
	index.Execute(w, data)
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

