package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) mover() string {
	return "soy un perro y camino"
}

func (pez) mover() string {
	return "soy un pez y estoy nadando"
}

func (pajaro) mover() string {
	return "soy un pajaro y vuelo"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)

}
