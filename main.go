package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DefaultPort    = "8080"
	DefaultVersion = "0.0.0"
)

type Response struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Ok bool `json:"ok"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

func CreateServerMux(version string) *http.ServeMux {
	mux := http.NewServeMux()

	// Default endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Message: "Hello world!"}
		data, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	})

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		resp := HealthResponse{Ok: true}
		data, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	})

	// Version endpoint
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		resp := VersionResponse{Version: version}
		data, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	})

	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = DefaultVersion
	}

	mux := CreateServerMux(version)

	log.Println(fmt.Sprintf("Server running on port: %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
