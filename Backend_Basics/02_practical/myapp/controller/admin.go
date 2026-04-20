package controller

import (
	httpresp "myapp/utils/httpResp"
	"net/http"
)

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpresp.ResponseWithError(w, http.StatusSeeOther, "no cookies on the plate")
			return false
		}

		httpresp.ResponseWithError(w, http.StatusInternalServerError, "Internal server error")
		return false
	}

	if cookie.Value != "my-value" {
		httpresp.ResponseWithError(w, http.StatusUnauthorized, "cookies does not match")
		return false
	}
	return true
}
