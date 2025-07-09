package utils

import "net/http"

type AppErr struct {
	errMessage string
	statusCode int
}

func (e *AppErr) Error() string {
	return e.errMessage
}

func (e *AppErr) GetStatusCode() int {
	return e.statusCode
}

func (e *AppErr) SetMessage(msg string) {
	e.errMessage = msg
}

func (e *AppErr) SetStatusCode(code int) {
	e.statusCode = code
}

var (
	EncodingMessageError = &AppErr{
		errMessage: "error encoding json",
		statusCode: http.StatusInternalServerError}
	ConnectingToDatabaseError = &AppErr{
		errMessage: "error connecting to the database",
		statusCode: http.StatusInternalServerError}
	InvalidRequestPayload = &AppErr{
		errMessage: "invalid request payload",
		statusCode: http.StatusBadRequest}
	UnitNotFoundError = &AppErr{
		errMessage: "not found",
		statusCode: http.StatusNotFound}
	DatabaseQueryError = &AppErr{
		errMessage: "database query error",
		statusCode: http.StatusInternalServerError}
	UnknownInternalServerError = &AppErr{
		errMessage: "unknown internal server error",
		statusCode: http.StatusInternalServerError}
)
