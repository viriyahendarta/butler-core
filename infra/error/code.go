package error

import "net/http"

const minCustomErrorCode Code = 1000

//Error code initialization
const (
	CodeBadRequest          Code = http.StatusBadRequest
	CodeUnauthorized        Code = http.StatusUnauthorized
	CodeForbidden           Code = http.StatusForbidden
	CodeInternalServerError Code = http.StatusInternalServerError
	CodeNotImplemented      Code = http.StatusNotImplemented

	//Custom error code must be 1000 and above
	//below that will be treated as http code

	CodeUnknown      Code = 1000
	CodeQueryGeneral Code = 2000
	CodeQueryUser    Code = 2001

	CodeDatabaseGeneral Code = 3000

	CodeRenderResponse Code = 5000
)

var mappingCodeValue = map[Code]string{
	CodeUnknown: "Unknown Error",

	CodeBadRequest:          "Bad Request",
	CodeUnauthorized:        "Unauthorized",
	CodeForbidden:           "Forbidden",
	CodeInternalServerError: "Internal Server Error",
	CodeNotImplemented:      "Not Implemented",

	CodeQueryGeneral: "Failed to run query",
	CodeQueryUser:    "Failed to run user query",

	CodeDatabaseGeneral: "Something wrong happened in database",

	CodeRenderResponse: "Failed to render response",
}
