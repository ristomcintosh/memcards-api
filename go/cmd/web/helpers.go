package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int, errorMsg ...string) {
	var msg string

	if errorMsg[0] != "" {
		msg = errorMsg[0]
	} else {
		msg = http.StatusText(status)
	}

	http.Error(w, msg, status)
}
