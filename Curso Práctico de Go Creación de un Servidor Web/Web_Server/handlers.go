package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "acá estamos de nuevo")

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Este es el home, acá iniciamos")
}
