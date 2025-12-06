package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"
	http.HandleFunc("/usuarios", getUsuarios)
	http.HandleFunc("/cursos", getCursos)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /usuarios")

	io.WriteString(w, "Este es mi endpoint de usuarios! \n")
}

func getCursos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /cursos")

	io.WriteString(w, "Este es mi endpoint de cursos! \n")
}
