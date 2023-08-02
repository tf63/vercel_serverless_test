package main

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Vercel Serverless Functions Go Example")
}

func main() {
	port := "80" 
	fmt.Printf("Starting server on port %s...\n", port)
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}