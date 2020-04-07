package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprint(w, "Hola mundo")
		http.Redirect(w, r, "/dos", http.StatusMovedPermanently) //301
	})
	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hola mundo, dos")
	})
	http.HandleFunc("/notFound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
