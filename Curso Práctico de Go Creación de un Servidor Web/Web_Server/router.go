package main

import (
	"net/http"
)

type Router struct { //se encarga de manejar todos los mapeos
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

//creamos función de tipo reciba para que podamos accederla
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
} //necesitamos el request

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) { //implementamos todo lo que se necesita para manejar los request
	//escritor y un objeto llamado request
	handler, exist := r.FindHandler(request.URL.Path) //esta función tiene dos parametros escritor y request, para saber cual es el path que esta buscando
	if !exist {                                       //para saber si esxiste o no
		w.WriteHeader(http.StatusNotFound) //saber si una ruta existe o no
		return                             //rompemos todo si no existe
	}
	handler(w, request)
	//mapa de reglas del router

}
