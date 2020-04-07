package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hola mundo, el metodo usado es:"+r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprint(w, "Mundo, el metodo usado es:"+r.Method)
		case "POST":
			fmt.Fprint(w, "Metodo POST")
		default:
			http.Error(w, "Metodo no soportado", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola mundo, dos")
	})
	http.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
