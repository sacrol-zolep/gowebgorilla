package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	port := "3333"
	router.HandleFunc("/usuarios", getUsuarios).Methods("GET")
	router.HandleFunc("/cursos", getCursos).Methods("GET")

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*3, "Server Timeout error!"),
		Addr:         "127.0.0.1:" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func getUsuarios(w http.ResponseWriter, r *http.Request) {
	log.Println("got /usuarios")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCursos(w http.ResponseWriter, r *http.Request) {
	log.Println("got /cursos")

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
