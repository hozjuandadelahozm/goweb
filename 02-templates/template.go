package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estructuras

type Usuarios struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

//Funciones

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una función"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	// fmt.Fprintln(rw, "Hola Mundo")
	template, err := template.New("index.html").Funcs(funciones).
		ParseFiles("index.html")

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
	}
}

func main() {

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//Servidor
	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	fmt.Println("El servidor esta corriendo en el puerto 5000")
	fmt.Println("Run server: http://localhost:5000/")

	log.Fatal(server.ListenAndServe())

}
