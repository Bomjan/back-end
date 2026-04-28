package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/date"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Enroll(w http.ResponseWriter, r *http.Request) {
	var en model.Enroll

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&en); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "Invalid Json")
		return
	}

	en.DateEnrolled = date.GetDate()
	defer r.Body.Close()

	if saveErr := en.EnrollStud(); saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate_key") {
			httpresp.ResponseWithError(w, http.StatusForbidden, "Duplicate keys")
			return
		} else {
			httpresp.ResponseWithError(w, http.StatusInternalServerError, saveErr.Error())
			return
		}
	}

	// No error
	httpresp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled successfully"})

}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	std_id, _ := strconv.ParseInt(sid, 10, 64)
	cid_id, _ := strconv.ParseInt(cid, 10, 64)

	e := model.Enroll{StdId: std_id, CourseId: cid_id}
	if getErr := e.Get(); getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpresp.ResponseWithError(w, http.StatusNotFound, "No such enrollments")
			return
		default:
			httpresp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
			return
		}
	}

	httpresp.ResponseWithJSON(w, http.StatusOK, e)

}
