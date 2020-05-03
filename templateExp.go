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

var (
	tags   = []string{"go", "java", "c#"}
	cursos = []Curso{Curso{Name: "rails", Time: 12}, Curso{Name: "JS", Time: 24}}
	user   = Usert{
		UserName: "henry",
		Age:      28,
		Activ:    true,
		Tags:     tags,
		Cursos:   cursos,
	}
)

//var templ = template.Must(template.ParseFiles("templates/index.html", "templates/footer.html", "templates/load.html"))
var templates = template.Must(template.ParseGlob("templates/*.html"))
var internalError = template.Must(template.ParseFiles("templates/error/internalError.html"))

func handleError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	internalError.Execute(w, status)
}
func renderPage(w http.ResponseWriter, page string, data interface{}) {
	//templ, _ := template.New("hola").Parse("hola mundo")
	w.Header().Set("Content-type", "text/html")
	if err := templates.ExecuteTemplate(w, page, user); err != nil {
		handleError(w, http.StatusInternalServerError)
	}
}
func main() {

	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "index.html", nil)
	})
	staticsFiles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", staticsFiles))
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
