package endpoints

import (
	"net/http"
	"disk"
	"logger"
	"fmt"
)

/*
 * Exists function helps to check if the file exists or not
 */
func Exists(w http.ResponseWriter, r *http.Request) {
	defer logger.RecoverFunc(w, r)

	if r.Method != http.MethodHead {
		logger.Error(r,
			fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))

		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}

	fPath := r.URL.Path[len("/sample/exists/"):]

	if exists, err := disk.Exists(fPath); !exists {
		logger.Error(r, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	logger.Success(r, http.StatusOK)
}