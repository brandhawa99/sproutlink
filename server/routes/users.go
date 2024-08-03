package routes

import (
	"fmt"
	"net/http"
)

func UserRouter() *http.ServeMux {
	userRouter := http.NewServeMux()

	userRouter.HandleFunc("GET /user", UserPong)

	return userRouter
}

func UserPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getting user\n")
}
