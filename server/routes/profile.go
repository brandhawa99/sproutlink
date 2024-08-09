package routes

import "net/http"

func ProfileRouter() *http.ServeMux {
	profileRouter := http.NewServeMux()

	return profileRouter
}
