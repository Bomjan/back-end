package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {

	var course model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "Shit man, your json is like an ass hole")
		return
	}
	fmt.Printf("DEBUG: Received course - ID: %d, Name: %s\n", course.CourseID, course.CourseName)
	r.Body.Close()
	saveErr := course.Create()
	if saveErr != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	httpresp.ResponseWithJSON(w, http.StatusCreated, "Course Created")
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseID, cError := getUserId(cid)
	if cError != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, cError.Error())
		return
	}

	var course model.Course
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&course); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "Your json has big ass error")
		return
	}
	defer r.Body.Close()
	if err := course.Update(courseID); err != nil {
		switch err {
		case sql.ErrNoRows:
			httpresp.ResponseWithError(w, http.StatusNotFound, "Course doesn't exisit, just like your brain")
		default:
			httpresp.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		}
	} else {
		httpresp.ResponseWithJSON(w, http.StatusOK, course)
		fmt.Println(course)
	}
}
func DeleteCourse(w http.ResponseWriter, r *http.Request)  {}
func GetCourse(w http.ResponseWriter, r *http.Request)     {}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {}
