package utils


import {
	"encoding/json"
	"io/ioutils"
	"net/http"
}

func ParseBody(r *http.Request, X interface{}) {
	if body, err:= ioutils.Readall(r.Body)
}