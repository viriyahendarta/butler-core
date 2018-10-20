package api

import (
	"net/http"
)

//HandlerFunc is a function for handling api request
type HandlerFunc func(r *http.Request) (interface{}, int, error)
