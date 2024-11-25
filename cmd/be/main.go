package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	port := ":8081"
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from Backend Server")
	fmt.Println("Replied with a hello message")
}
