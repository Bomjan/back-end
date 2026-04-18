package controller

import (
	"encoding/json"
	"myapp/model"
	"net/http"
)

// This shoudl add student to the database.
func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&stud)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	r.Body.Close()
	dbErr := stud.Create()
	if dbErr != nil {
		w.Write([]byte(dbErr.Error()))
		return
	}

	// if there is no error at all then
	w.Write([]byte("The response is sucessful"))
}
