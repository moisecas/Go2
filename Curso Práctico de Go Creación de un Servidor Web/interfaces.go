package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro y camino"
}

func (pez) nada() string {
	return "soy un pez y estoy nadando"
}

func (pajaro) vuela() string {
	return "soy un pajaro y vuelo"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func nadarPez(pe pez) {
	fmt.Println(pe.nada())
}

func volarPajaro(pa pajaro) {
	fmt.Println(pa.vuela())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	nadarPez(pe)
	pa := pajaro{}
	volarPajaro(pa)

}
