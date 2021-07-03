package handler

import (
	"encoding/json"
	"gojwtexampleapi/entity"
	"gojwtexampleapi/infra"
	"gojwtexampleapi/usecase"
	"net/http"

	"github.com/google/uuid"
)

// fictitious sign, no need username or password
// only return the token
var userID = uuid.New().String()
var userName = "James"

var users []entity.User = make([]entity.User, 0)

func SignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		infra := infra.NewMakeToken()
		usecase := usecase.NewCreateTokenImpl(infra)

		token, err := usecase.Handle(userID, userName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
	}
}
