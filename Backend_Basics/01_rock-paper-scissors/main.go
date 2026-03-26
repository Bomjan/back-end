package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rockpaperscs/rps"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("index.html")

	if err != nil {
		log.Println(err)
		return
	}
	// templ.Execute(writer, the data that we want to inject along with the templ)
	err = templ.Execute(w, nil) // second nil is the data that we want to inject.

	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-type", "text/html")

}

func playHandler(w http.ResponseWriter, r *http.Request) {
	result := rps.PlayRound(1)

	// fmt.Println("winner", result.Winner)
	// fmt.Println("computer choice", result.ComputerChoice)
	// fmt.Println("round result", result.RoundResule)

	obj, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(obj)
}
func main() {
	var port = 8080
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playHandler)
	log.Println("Starting web server on port: ", port)
	http.ListenAndServe(":8080", nil)
}

// localhost:8080/
