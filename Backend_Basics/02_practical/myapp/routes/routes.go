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
	p := mux.Vars(r)
	course := p["course"]

	if course == "" {
		course = r.URL.Query().Get("course")
	}

	_, err := w.Write([]byte("Hello World\nThis course is  " + course))
	if err != nil {
		fmt.Println(err)
	}
}
