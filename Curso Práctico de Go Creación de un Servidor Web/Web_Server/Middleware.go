package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("autenticación ok")
			if flag {
				f(w, r)
			} else {
				return
			}
		}

	}

	//estamos haciendo esto para evitar repetir cod

}
