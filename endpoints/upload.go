package endpoints

import (
	"net/http"
	"disk"
	"logger"
	"fmt"
)

/*
 * Upload function helps to save a file using web api
 */
func Upload(w http.ResponseWriter, r *http.Request) {
	defer logger.RecoverFunc(w, r)

	if r.Method != http.MethodPost {
		logger.Error(r,
			fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))

		http.Error(w,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}

	fpath := r.URL.Path[len("/sample/upload/"):]

	fin, handler, err := r.FormFile("file")
	if err != nil || handler.Size == 0 {
		logger.Error(r, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	inProgressExt := disk.GetRandomExt()

	/*
	 * Write file with the extension to keep failures in the case of network
	 * problems
	 */
	err = disk.WriteToStorage(fin, fpath + inProgressExt, handler.Size)
	if err != nil {
		logger.Error(r, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fin.Close()

	/*
	 * Rename file as its original file name to indicate that it successfully
	 * uploaded
	 */
	err = disk.Rename(fpath + inProgressExt, fpath)
	if err != nil {
		logger.Error(r, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.Success(r, http.StatusCreated)
}
