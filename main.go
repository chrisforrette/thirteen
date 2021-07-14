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
	DefaultName    = "Thirteen"
	DefaultVersion = "0.0.0"
)

type service struct {
	Name    string
	Version string
}

type Response struct {
	Message string `json:"message"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type HealthResponse struct {
	Ok bool `json:"ok"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

func CreateServerMux(svc service) *http.ServeMux {
	mux := http.NewServeMux()

	// Default endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			Message: "Hello world!",
			Service: svc.Name,
			Version: svc.Version,
		}
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
		resp := VersionResponse{Version: svc.Version}
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

	name := os.Getenv("NAME")
	if name == "" {
		name = DefaultName
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = DefaultVersion
	}

	svc := service{
		Name:    name,
		Version: version,
	}

	mux := CreateServerMux(svc)

	log.Println(fmt.Sprintf("Server running on port: %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
