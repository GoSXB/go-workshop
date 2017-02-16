package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Profile is a platform user.
type Profile struct {
	Username string
	Name     string
	Email    string
}

// Thing is a thing. YEAH.
type Thing struct {
	ID    uint
	Owner string
	Name  string
}

var (
	users  = make(map[string]Profile)
	things = make(map[int]Thing)
)

func init() {
	users["dolanor"] = Profile{Username: "dolanor", Name: "Tanguy Herrmann", Email: "tanguy@tuxago.com"}
	users["soulou"] = Profile{Username: "soulou", Name: "Léo Unbekandt", Email: "leo@scalingo.com"}

	things[1] = Thing{ID: 1, Owner: "dolanor", Name: "Mon raspberry"}
	things[2] = Thing{ID: 2, Owner: "dolanor", Name: "Mon XPS 13"}
	things[3] = Thing{ID: 3, Owner: "soulou", Name: "Mon XPS 13"}
}

func toJSON(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func displayThings(w http.ResponseWriter, r *http.Request) {
	b, err := toJSON(things)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func displayThing(w http.ResponseWriter, r *http.Request) {
	thingID := mux.Vars(r)["thingId"]
	id, err := strconv.Atoi(thingID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	thing, ok := things[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := toJSON(thing)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func displayUsers(w http.ResponseWriter, r *http.Request) {
	b, err := toJSON(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	user, ok := users[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := toJSON(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", displayUsers)
	r.HandleFunc("/users/{username}", displayUser)
	r.HandleFunc("/things", displayThings)
	r.HandleFunc("/things/{thingId}", displayThing)
	//http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3333", r))
}
