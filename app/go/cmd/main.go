package main

import (
	"fmt"
	"net/http"

	"github.com/tf63/vercel_serverless_test/handler"
)

func main() {
	port := "80"
	fmt.Printf("Starting server on port %s...\n", port)
	http.HandleFunc("/", handler.Handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
