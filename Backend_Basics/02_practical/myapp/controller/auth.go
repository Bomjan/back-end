package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) { // this is a post method
	var admin model.Admin

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

func Login(w http.ResponseWriter, r *http.Request) {
	// get the data, read the data, compare
	var data model.Login

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		httpresp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	defer r.Body.Close()

	fmt.Println(data)

	returndedData := model.Admin{Email: data.Email}

	if getErr := returndedData.Read(); getErr != nil {

		switch getErr {
		case sql.ErrNoRows:
			httpresp.ResponseWithError(w, http.StatusUnauthorized, "The email has not been registered")
		default:
			httpresp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())

		}
		return
	}
	fmt.Println("from database ", returndedData)

	if data.Password != returndedData.Password {
		httpresp.ResponseWithError(w, http.StatusUnauthorized, "Login failed because password did not match")
		return
	} else {
		loginResponse := model.LoginResponse{
			FirstName: returndedData.FirstName,
			LastName:  returndedData.LastName,
			Email:     returndedData.Email,
		}

		fmt.Println("passing the data", loginResponse)

		cookie := http.Cookie{
			Name:    "my-cookie",
			Value:   "my-value",
			Expires: time.Now().Add(30 * time.Minute),
			Secure:  true,
		}
		http.SetCookie(w, &cookie)
		httpresp.ResponseWithJSON(w, http.StatusOK, loginResponse)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})

	httpresp.ResponseWithJSON(w, http.StatusOK, map[string]string{"status": "User logged out sucessfully"})
}
