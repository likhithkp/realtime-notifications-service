package handler

import (
	"encoding/json"
	"net/http"
	"realtime-notifications-service/services"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not a valid method", http.StatusMethodNotAllowed)
		return
	}

	user := new(User)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Name == "" || user.Password == "" {
		http.Error(w, "All fields are mandatory", http.StatusBadRequest)
		return
	}

	res := services.CreateUserService(user.Name, user.Email, user.Password)

	if res != nil {
		json.NewEncoder(w).Encode(&res)
		return
	}

}
