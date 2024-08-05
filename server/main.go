package main

import (
	"fmt"
	"net/http"

	"github.com/brandhawa99/sproutlink/server/middleware"
	"github.com/brandhawa99/sproutlink/server/routes"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	// created main server mux
	mux := http.NewServeMux()

	//middleware
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
	)

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	// sub routers
	userRouter := routes.UserRouter()
	mux.Handle("/users/", http.StripPrefix("/users", userRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	fmt.Println("Server Listening on Port 8080")
	log.Fatal(server.ListenAndServe())

}
