package handler

import "net/http"

func SayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello friend!"))
	}
}
