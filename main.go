package main

import (
	"fmt"
	"log"
	"net/http"

	"./mux"
)

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola usuario "+this.name)
}
func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde una funcion anonima")
}
func main() {
	user := &User{"henry"}
	mux := mux.CreateMux()
	mux.AddFun("/hola", hola)
	mux.AddHandle("/user", user)
	/* http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
	 	fmt.Fprint(w, "Hola mundo, dos")
	 })
	http.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})*/
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
