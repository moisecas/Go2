package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)

}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}
func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("nombre", tarea.nombre)
		fmt.Println("descripcion", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {

			fmt.Println("nombre", tarea.nombre)
			fmt.Println("descripcion", tarea.descripcion)

		}

	}
}

func main() {
	t1 := &task{
		nombre:      "completar curso",
		descripcion: "completar el curso en esta semana",
	}
	t2 := &task{
		nombre:      "completar curso 2",
		descripcion: "completar el curso 2 en esta semana",
	}
	t3 := &task{
		nombre:      "completar curso 3",
		descripcion: "completar el curso 3 en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("tareas completadas")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["moises"] = lista

	t4 := &task{
		nombre:      "completar curso 4",
		descripcion: "completar el curso 4 en esta semana",
	}
	t5 := &task{
		nombre:      "completar curso 5",
		descripcion: "completar el curso 58 en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}
	mapaTareas["tayson"] = lista2

	fmt.Println("tareas de moises")
	mapaTareas["moises"].imprimirLista()
	fmt.Println("tareas de tayson")
	mapaTareas["tayson"].imprimirLista()

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("index", i, "nombre", lista.tasks[i].nombre)
	}
	for index, tarea := range lista.tasks {
		fmt.Println("index", index, "nombre", tarea.nombre)

	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}

}
