package main

import (
	"html/template"
	"log"
	"net/http"
)

type Curso struct {
	Name string
	Time int
}
type Usert struct {
	UserName string
	Age      int
	Activ    bool
	Tags     []string
	Cursos   []Curso
}

func main() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		//templ, _ := template.New("hola").Parse("hola mundo")
		templ, _ := template.ParseFiles("templates/index.html", "templates/footer.html")
		tags := []string{"go", "java", "c#"}
		cursos := []Curso{Curso{Name: "rails", Time: 12}, Curso{Name: "JS", Time: 24}}
		user := Usert{
			UserName: "henry",
			Age:      28,
			Activ:    true,
			Tags:     tags,
			Cursos:   cursos,
		}
		templ.Execute(w, user)
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
