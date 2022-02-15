package main

func main() {
	server := NewServer(":3000") //traemos la función del archivo server junto al puerto
	server.Handle("/", HandleRoot)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth))
	server.Listen()

}
