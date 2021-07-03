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
	r.HandleFunc("/protected", middleware.MiddlewareVerifyTokenHandler(handler.SayHello())).Methods("GET")

	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	s.ListenAndServe()
}
