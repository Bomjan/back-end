package controller

import (
	"encoding/json"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) { // this is a post method
	var admin model.Signup

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admin); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	r.Body.Close()

	if saveErr := admin.Create(); saveErr != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "Error while saving")
		return
	}

	httpresp.ResponseWithJSON(w, http.StatusOK, "Signup sucessful")
}
