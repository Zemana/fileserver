package logger

import (
	"net/http"
	"fmt"
	"log"
)

/*
 * RecoverFunc helps to recover the critical operations and log the problem.
 * Then the program may work continuously
 */
func RecoverFunc(w http.ResponseWriter, r *http.Request) {
	if e := recover(); e != nil {
		log.SetPrefix("Panic: ")
		log.Println(e, r)

		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
}

/*
 * Error helps to print logs in error situation
 */
func Error(r *http.Request, err error) {
	output := format(r)
	output += fmt.Sprintf("\n\t\t\t\t\tMsg: %s", err)

	log.SetPrefix("Error: ")
	log.Println(output)
}

/*
 * Success helps to print logs when operation successfully ends
 */
func Success(r *http.Request, status int) {
	output := format(r)
	output += fmt.Sprintf("\n\t\t\t\t\tStatus: %d", status)

	log.SetPrefix("Success: ")
	log.Println(output)
}

/*
 * format helps to prepare human-readable formatted output of http request
 */
func format(r *http.Request) string {
	if r == nil {
		return ""
	}

	infoMap := map[string] interface{} {
		"From": r.RemoteAddr,
		"Method": r.Method,
	}

	output := r.URL.Path
	for k, v := range infoMap {
		output += fmt.Sprintf("\n\t\t\t\t\t%s: %s", k, v)
	}

	return output
}