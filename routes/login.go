package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"golang.org/x/crypto/bcrypt"
)
func sendJsonErr (w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
func LoginGet(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/login/" {
		w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "404 page not found")
        return
    }
	index, err := template.ParseFiles("index.html", "templates/login.html", "templates/navbar.html")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]string)
	data["title"] = "Jedders | Login"
	index.Execute(w, data)
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	type UserRequest struct {
		Username                      string `json:"username"`
		Password                      string `json:"password"`
	}
	decoder := json.NewDecoder(req.Body)

	var user UserRequest
	err := decoder.Decode(&user)
	if err != nil {
		sendJsonErr(w, "Server Error decoding JSON", http.StatusInternalServerError)
    }
	if(len(user.Username) < 1 || len(user.Password) < 1){
		sendJsonErr(w, "Username and password must be at least 1 character long", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		sendJsonErr(w, "Server Error: could not hash password", http.StatusInternalServerError)
		return
	}
	fmt.Println("Hashed password:", string(hash))

	dbUser := models.User{Username: user.Username, Password: string(hash)}
	result := initializers.ConnectDB().Create(&dbUser)
	if result.Error != nil {
		sendJsonErr(w, "Server Error: could not create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Successfully created user"})



}