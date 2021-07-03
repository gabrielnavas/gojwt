package middleware

import (
	"errors"
	"net/http"
	"strings"

	"gojwtexampleapi/infra"
	"gojwtexampleapi/usecase"
)

func MiddlewareVerifyTokenHandler(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
			return
		}

		// need a factory here. :)
		infra := infra.NewGetDataInfra()
		usecase := usecase.NewVerifyToken(infra)

		_, err = usecase.Handle(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
			return
		}
		handler(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	var strSplit []string = strings.Split(bearerToken, " ")
	if len(strSplit) == 0 {
		return "", errors.New("need on header: Authorization: <token>")
	}
	token := strings.TrimSpace(strSplit[1])
	return token, nil
}
