package main

import (
	"fmt"

	"net/http"
	"os"
	"path/filepath"

	"example.com/m/v2/initializers"
	"example.com/m/v2/routes"
)

func init() {
	initializers.LoadEnv() 
	initializers.ConnectDB()
}


func main() { 
	go initializers.TailwindCompile()
    mux := http.NewServeMux()

    mux.HandleFunc("/", routes.Root)
    
	mux.HandleFunc("/styles.css", css)
    
	mux.HandleFunc("/headers", headers)
	
	mux.HandleFunc("GET /login/", routes.LoginGet)
	mux.HandleFunc("/signUp", routes.SignUp)
	
	
	dir := "public"
	files, err := os.ReadDir(dir) 
	if err != nil {
		fmt.Println("Failed to read", dir, err)
	}
	for _, file := range files {
		if(!file.IsDir()){
			filePath := filepath.Join(dir, file.Name())
			fmt.Println("Serving", filePath)
			mux.HandleFunc("/public/" + file.Name(), func(w http.ResponseWriter, req *http.Request) {
				http.ServeFile(w, req, filePath)
			})

	}
}
	


	fmt.Println("Server started at port " + os.Getenv("PORT"))
	fmt.Println("http://localhost:" + os.Getenv("PORT"))
    http.ListenAndServe(":" + os.Getenv("PORT"), mux)
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






