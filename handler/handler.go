package handler

import (
	"encoding/json"
	"gojwtexampleapi/entity"
	"gojwtexampleapi/usecase"
	"net/http"
)

var users []entity.User = make([]entity.User, 0)

func SignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyJson struct {
			Name string `json:"name"`
		}
		json.NewDecoder(r.Body).Decode(&bodyJson)

		if bodyJson.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("missing name")
			return
		}

		user := entity.User{
			ID:   uint64(len(users) + 1),
			Name: bodyJson.Name,
		}
		token, err := usecase.MakeToken(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users = append(users, user)

		response := struct {
			User  entity.User `json:"user"`
			Token string      `json:"token"`
		}{
			User:  user,
			Token: token,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	}
}
