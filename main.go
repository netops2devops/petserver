package main

import (
	"encoding/json"
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
	http.HandleFunc("GET /pets", RedirectGetPetHandler)
	http.HandleFunc("POST /pets", RedirectPostPetHandler)

	http.HandleFunc("GET /pets/", ShowAllPetsHandler)
	http.HandleFunc("POST /pets/", NewPetHandler)

	http.HandleFunc("GET /pets/{name}", GetPetHandler)
	http.HandleFunc("DELETE /pets/{name}", DeletePetHandler)

	http.ListenAndServe(":8080", nil)
}

func RedirectGetPetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/pets/", http.StatusMovedPermanently)
}

func RedirectPostPetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/pets/", http.StatusMovedPermanently)
}

func ShowAllPetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetPetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	for _, v := range p {
		if v.Name == name {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func NewPetHandler(w http.ResponseWriter, r *http.Request) {
	var newpet Pets
	err := json.NewDecoder(r.Body).Decode(&newpet)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	p = append(p, newpet)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func DeletePetHandler(w http.ResponseWriter, r *http.Request) {
	petName := r.PathValue("name")

	for i, v := range p {
		if v.Name == petName {
			p = append(p[:i], p[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string][]Pets{
				"Pets": p,
			})

			return
		}
	}

	http.Error(w, "Pet not found", http.StatusNotFound)
}
