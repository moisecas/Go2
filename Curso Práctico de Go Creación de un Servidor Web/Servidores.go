package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	servidores := []string{
		"http://uis.edu.co",
		"http://facebook.com",
		"http://instagram.com",
		"http://linkedin.com",
		"http://linkedin.com",
		"http://sitioparaprobarerror.com",
	}
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	tiempoPaso := time.Since(inicio)
	fmt.Printf("tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no disponible =(")
	} else {
		fmt.Println(servidor, "estamos onfire :)")
	}

}
