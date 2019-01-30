package endpoints

import (
	"net/http"
	"logger"
)

/*
 * Up function helps to check if web service is running or not
 */
func Up(w http.ResponseWriter, r *http.Request) {
	defer logger.RecoverFunc(w, r)
}