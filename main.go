package main

import (
	"net/http"
	"os"

	"gojwtexampleapi/config"
	"gojwtexampleapi/handler"
	"gojwtexampleapi/middleware"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv(config.ACCESS_SECRET, "any_key_123")

	r := mux.NewRouter()
	r.Use(middleware.CORSMiddleware)
	r.HandleFunc("/sign", handler.SignHandler()).Methods("POST")
	r.HandleFunc("/user", handler.MiddlewareVerifyTokenHandler(handler.GetUsersHandler())).Methods("GET")

	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	s.ListenAndServe()

	// user := User{
	// 	ID:   1,
	// 	Name: "gabs",
	// }
	// tokenStr, _ := MakeToken(user)
	// userMetada, err := ExtractTokenMetadata(tokenStr)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(userMetada)
}
