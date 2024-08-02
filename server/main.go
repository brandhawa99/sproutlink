package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/brandhawa99/sproutlink/server/middleware"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	mux := http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
	)
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
	mux.HandleFunc("/ping", h1)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}
	fmt.Println("Server Listening on Port 8080")
	log.Fatal(server.ListenAndServe())
}
