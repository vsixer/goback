package main

import (
	"goback/core/output"
	"net/http"
)

func main() {

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// test, _ := json.Marshal(authToket)

	// auth.Login()
}

func handler(w http.ResponseWriter, r *http.Request) {
	output.Ok()
}

func initHTTPServer() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8081", nil)
}
