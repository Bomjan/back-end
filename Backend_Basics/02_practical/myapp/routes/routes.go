package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up all the HTTP routes and starts the server.
// This is where the application actually starts listening for requests.
func InitializeRoutes() {
	var port = 8080
	router := mux.NewRouter()

	// Register route handlers - the order matters in mux.
	// /home and /home/{course} are catch-all routes, /student/add is the API endpoint.
	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/home/{course}", homeHandler)

	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	log.Println("Application running on port", port)
	// ListenAndServe blocks forever - that's why it's wrapped with log.Fatal.
	// If the server stops, the application exits.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// homeHandler is a simple test endpoint - just to verify the server is working.
// Not meant for production, but useful during development to check if routing works.
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// this takes the anyvalue that is passed after the /home url.
	// That is how request works. understood for the first time.
	// In our case we have {course}. so everything passed after the /home is stored in a map with course key
	p := mux.Vars(r) // r is for all the requests, there is no filtration
	course := p["course"]
	fmt.Println(p)

	// Write the response directly - no JSON here, just plain text for quick testing.
	_, err := w.Write([]byte("Hello World\nThis course is " + course))
	if err != nil {
		fmt.Println(err)
	}
}
