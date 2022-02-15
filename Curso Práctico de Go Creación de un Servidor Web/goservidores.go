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
	for _, servidor := range servidores {
		go revisarServidor(servidor, canal) //go, para nuevo hilo, nuevo proceso independiente de main
		//canal, recibe un nuevo parametro
	}
	//alguien pendiente o escuchando ese cana, hay un canal para las 4 subrutinas, por el channel
	for i := 0; i < len(servidores); i++ { //len devuelve un valor de longitud almacenado.
		fmt.Println(<-canal)
		//el for es para consultar los 4 sitios, for inicia en 0 hasta que i sea menor que la longitud de los servidores
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no disponible =(")
		canal <- servidor + "no esta on, disponible" //<- para usar el canal, menciono el canal y el mensaje, transmitimos la data
	} else {
		fmt.Println(servidor, "estamos onfire :)")
		canal <- servidor + "esta ok, fucionando"
	}

}
