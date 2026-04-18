package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	var port = 8080
	// router := mux.NewRouter()
	router := mux.NewRouter()

	router.HandleFunc("/home", homeHandler)
	router.HandleFunc("/home/{course}", homeHandler)
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// this takes the anyvalue that is passed after the /home url.
	// That is how request works. understood for the first time.
	// In our case we have {course}. so everything passed after the /home is stored in a map with course key
	p := mux.Vars(r) // r is for all the requests, there is no filtration
	course := p["course"]
	fmt.Println(p)

	_, err := w.Write([]byte("Hello World\nThis course is " + course))
	if err != nil {
		fmt.Println(err)
	}
}
