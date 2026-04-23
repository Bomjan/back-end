package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/date"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"strings"
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
