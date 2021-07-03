package handler

import (
	"errors"
	"net/http"
	"strings"

	"gojwtexampleapi/infra"
	"gojwtexampleapi/usecase"
)

func MiddlewareVerifyTokenHandler(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := ExtractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
			return
		}
		usecase := usecase.ExtractTokenMetadata(infra.GetData)
		_, err = usecase(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
			return
		}
		handler(w, r)
	})
}

func ExtractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	var strSplit []string = strings.Split(bearerToken, " ")
	if len(strSplit) == 0 {
		return "", errors.New("need on header: Authorization: <token>")
	}
	token := strings.TrimSpace(strSplit[1])
	return token, nil
}
