package server

import (
	"fmt"
	"github.com/codewithsirohi/hisab-kitab/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// Create a new instance of the mux router
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/", controllers.HomeController).Methods("GET")
	router.HandleFunc("/about", controllers.AboutController).Methods("GET")
	router.HandleFunc("/status", controllers.AboutController).Methods("GET")

	// Specify the port and start the server
	port := 8080

	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
