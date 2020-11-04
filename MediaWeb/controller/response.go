package controller

import (
	"MediaWeb/api/defs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ErrResponse(w *http.ResponseWriter, err defs.HttpErr) {
	(*w).WriteHeader(err.StatusCode)
	errJson, e := json.Marshal(err.Err)
	if e != nil {
		fmt.Println(e)
	} else {
		io.WriteString(*w, string(errJson))
	}
}

func SuccessResponse(w *http.ResponseWriter, info string) {
	(*w).WriteHeader(200)
	io.WriteString(*w, info)
}
