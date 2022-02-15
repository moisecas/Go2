package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string) //creación del canal con chan, de tipo string

	servidores := []string{ //variable servidores creada cmo slice de tipo string
		"http://uis.edu.co",
		"http://facebook.com",
		"http://instagram.com",
		"http://linkedin.com",
		"http://linkedin.com",
		"http://sitioparaprobarerror.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal) //go, para nuevo hilo, nuevo proceso independiente de main
			//canal, recibe un nuevo parametro
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	//alguien pendiente o escuchando ese cana, hay un canal para las 4 subrutinas, por el channel

	tiempoPaso := time.Since(inicio)
	fmt.Printf("tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no disponible =(")

	} else {
		fmt.Println(servidor, "estamos onfire :)")

	}

}
