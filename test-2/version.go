package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"io/ioutil"

	"github.com/gorilla/mux"
)

const (
	versionFile = "./resource/version.json"
)

func main() {
	fmt.Println("starting http server ")
	r := mux.NewRouter()
	r.HandleFunc("/version", getversion)

	s := &http.Server{
		Handler:      r,
		Addr:         ":8000", // To make sure this is accessible outside docker network
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func getversion(w http.ResponseWriter, r *http.Request) {

	absPathVersion, _ := filepath.Abs(versionFile)
	version, err := ioutil.ReadFile(absPathVersion)
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something went wrong on the server, couldn't determine version"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(version)
}
