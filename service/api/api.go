package api

import (
	"net/http"
)

type HandlerFunc func(r *http.Request) (interface{}, int, error)
