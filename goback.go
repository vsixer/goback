package main

import (
	"goback/auth"
)

func main() {

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// test, _ := json.Marshal(authToket)

	auth.Login()
}
