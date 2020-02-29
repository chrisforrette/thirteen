package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "0.0.0"
	}

	// Default endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Message: "Hello, world!"}
		data, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	})

	// Health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		resp := VersionResponse{Version: version}
		data, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	})

	log.Println(fmt.Sprintf("Server running on port: %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}