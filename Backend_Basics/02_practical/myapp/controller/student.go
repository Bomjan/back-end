package controller

import (
	"encoding/json"
	"myapp/model"
	"net/http"
)

// This should add student to the database.
func AddStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&student)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	r.Body.Close()

	dbErr := student.Create()
	if dbErr != nil {
		w.Write([]byte(dbErr.Error()))
		return
	}
	w.Write([]byte("Student created successfully"))
}
