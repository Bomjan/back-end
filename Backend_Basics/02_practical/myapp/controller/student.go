package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"net/http"
	"strconv"

	httpresp "myapp/utils/httpResp"

	"github.com/gorilla/mux"
)

// AddStudent handles POST requests to create a new student record.
// It decodes the JSON request body, validates the input, and persists the student to the database.
// Returns 201 Created on success, 400 Bad Request on validation or decode errors.
//
// Why this matters:
// - Proper HTTP status codes enable clients to handle success/failure cases correctly.
// - Returning structured JSON errors helps frontend developers debug issues quickly.
// - Closing the request body prevents resource leaks under high load.

func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student
	decoder := json.NewDecoder(r.Body)

	// Decode the incoming JSON request body into the student struct.
	// If the JSON is malformed or missing required fields, return a 400 error with details.
	// This validation is important: rejecting malformed input early prevents invalid data from reaching the database.
	if err := decoder.Decode(&stud); err != nil {
		// response, _ := json.Marshal(map[string]string{"error": err.Error()})
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(response)

		// Now from utils
		httpresp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	r.Body.Close()

	// Persist the student record to the database.
	// Create() returns an error if validation fails or the database operation fails.
	// Database-level validation acts as the final safety net before data is stored.
	saveErr := stud.Create()
	if saveErr != nil {
		// response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(response)

		// Now from the utils, you get
		httpresp.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	// Successfully created the student - return 201 Created with a confirmation message.
	// Using 201 (not 200) is important: it signals resource creation, enabling correct client-side behavior (e.g., redirects, cache invalidation).

	// response, _ := json.Marshal(map[string]string{"status": "Student Added"})
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(response)

	// From utils
	httpresp.ResponseWithJSON(w, http.StatusCreated, "Student Added")
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	// get student of the id stdId
	s := model.Student{StdId: stdId}
	getErr := s.Read() // run the query to find details of the student of the id stdId

	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpresp.ResponseWithError(w, http.StatusNotFound, "Student Not Found")
		default:
			httpresp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpresp.ResponseWithJSON(w, http.StatusOK, s)

}

func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}

	return userId, nil
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	old_sid := mux.Vars(r)["sid"]
	old_stdId, idErr := getUserId(old_sid)

	if idErr != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	defer r.Body.Close()

	if updateErr := stud.Update(old_stdId); updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpresp.ResponseWithError(w, http.StatusNotFound, "Student Not Found")
		default:
			httpresp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpresp.ResponseWithJSON(w, http.StatusOK, stud)
	}
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]

	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "Student Deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, err := model.GetAllStudents()
	if err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.ResponseWithJSON(w, http.StatusOK, students)
}
