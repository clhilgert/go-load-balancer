package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const backendURL = "http://localhost:8081"

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	// Forward request to the backend
	req, err := http.NewRequest(r.Method, backendURL+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request to backend", http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to contact backend server", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	fmt.Printf("Response from server: %s\n\n", resp.Status)
}
