package endpoints

import (
	"net/http"
	"disk"
	"logger"
	"fmt"
)

/*
 * Download function helps to download a file using web api
 */
func Download(w http.ResponseWriter, r *http.Request) {
	defer logger.RecoverFunc(w, r)

	if r.Method != http.MethodGet {
		logger.Error(r,
			fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))

		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}

	fpath := r.URL.Path[len("/sample/download/"):]

	storagePath, err := disk.ConvertToStoragePath(fpath)
	if err != nil {
		logger.Error(r, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, storagePath)
	logger.Success(r, http.StatusOK)
}