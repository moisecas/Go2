package main

import (
	"net/http"
)

type Server struct { //creación del servidor con una propiedad de tipo string
	port   string
	router *Router
}

func NewServer(port string) *Server { //instanciar el servidor recibe el puerto y devuelve el servidor, apunta
	return &Server{ //la función recibe el puerto y apunta al server (struct)
		port:   port, //capaces de crear el servidor
		router: NewRouter(),
	}

}
func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler // va a ser capaz de asociar al handler

}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)

	}
	return f
	//requiere que iteremos
}

func (s *Server) Listen() error { //manejo de error de conexión
	//se agrego el struct router
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

//para correr los archivos de la carpeta se usa go run * o go run . en caso de error go: cannot find main module; see ‘go help modules’ ejecutar go env -w GO111MODULE=off
