package main

import (
	"log"
	"net/http"
)

type Pets struct {
	Name  string `json:"name,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Age   int    `json:"age,omitempty"`
	Color string `json:"color,omitempty"`
}

func Initializer() []Pets {
	return []Pets{
		{"Juniper", "Cat", 5, "Orange"},
		{"Ashby", "Cat", 5, "Gray"},
		{"Bruce", "Dog", 8, "Golden"},
	}
}

var p = Initializer()

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /pets/", ShowAllPetsHandler)
	mux.Handle("POST /pets/", loggerMiddleware(authMiddleware(http.HandlerFunc(AddNewPetHandler))))

	mux.HandleFunc("GET /pets/{name}", GetPetHandler)
	mux.Handle("DELETE /pets/{name}", loggerMiddleware(authMiddleware(http.HandlerFunc(DeletePetHandler))))

	log.Print("server listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
