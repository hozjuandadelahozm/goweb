package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estructuras

// type Usuarios struct {
// 	UserName string
// 	Edad     int
// 	Activo   bool
// 	Admin    bool
// 	Cursos   []Curso
// }

type Usuarios struct {
	UserName string
	Edad     int
}

type Curso struct {
	Nombre string
}

//Funciones

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una función"
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler error
func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Función de render template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		// http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handlerError(rw)
	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	// funciones := template.FuncMap{
	// 	"saludar": Saludar,
	// }

	// fmt.Fprintln(rw, "Hola Mundo")
	// template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	usuario := Usuarios{"Juan", 25}

	// template.Execute(rw, usuario)

	renderTemplate(rw, "inde.html", usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	renderTemplate(rw, "registro.html", nil)
}

func main() {

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Servidor
	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	fmt.Println("El servidor esta corriendo en el puerto 5000")
	fmt.Println("Run server: http://localhost:5000/")

	log.Fatal(server.ListenAndServe())

}
