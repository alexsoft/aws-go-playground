package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("aws-go-playground")
	fmt.Printf("port = %s\n", port)
	fmt.Printf("app-server = %s\n", getAppServer())

	http.HandleFunc("/", index)
	http.HandleFunc("/healthcheck", healthcheck)

	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "AWS Go Playground\n")
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	server := getAppServer()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"ok": true, "app-server": server})
}

func getAppServer() string {
	server := os.Getenv("APP_SERVER")
	if server == "" {
		server = "NA"
	}
	return server
}
