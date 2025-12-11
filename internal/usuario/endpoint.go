package usuario

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		Nombre   string `json:"nombre"`
		Apellido string `json:"apellido"`
		Correo   string `json:"correo"`
		Telefono string `json:"telefono"`
	}

	ErrorResp struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

func makeCreateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReq

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{"Request formato invalido"})
			return
		}

		if req.Nombre == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{"Nombre es requerido"})
			return
		}

		if req.Apellido == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorResp{"Apellido es requerido"})
			return
		}

		fmt.Println(req)
		json.NewEncoder(w).Encode(req)
	}
}

func makeGetEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Consultar usuario")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Consultar todo usuario")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeUpdateEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Actualizar usuario")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndpoint() Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Eliminar usuario")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
