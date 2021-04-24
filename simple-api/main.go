package main

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrorNoSuchKey = errors.New("no such key")

	store = make(map[string]string)
)

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)

	return nil
}

func Put(key string, value string) error {
	store[key] = value

	return nil
}

// keyValuePutHandler expects to be called with a PUT request for
// the "/v1/key/{key}" resource.
func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
