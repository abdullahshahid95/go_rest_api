package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json: "firstName"`
	Lastname  string `json: "lastName"`
	Email     string `json: "email"`
}

type Profile struct {
	Department  string `json: "department"`
	Designation string `json: "designation"`
	Employee    User   `json: "employee"`
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)
	w.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)
	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("id couldn't be converted to integer"))
		return
	}

	if id >= len(profiles) || id < 0 {
		w.WriteHeader(404)
		w.Write([]byte("No profiles found with specified id"))

		return
	}

	profile := profiles[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("id couldn't be converted to integer"))
		return
	}

	if id >= len(profiles) || id < 0 {
		w.WriteHeader(404)
		w.Write([]byte("No profiles found with specified id"))

		return
	}

	var updatedProfile Profile
	json.NewDecoder(r.Body).Decode(&updatedProfile)
	profiles[id] = updatedProfile

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProfile)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("id couldn't be converted to integer"))
		return
	}

	if id >= len(profiles) || id < 0 {
		w.WriteHeader(404)
		w.Write([]byte("No profiles found with specified id"))

		return
	}

	profiles = append(profiles[:id], profiles[id+1:]...)

	w.WriteHeader(200)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")
	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}
