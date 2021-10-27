package infrastructure

import "net/http"

// GetMessageForHTTPStatus will return an appropriate string based on the given
// HTTP status code.
func GetMessageForHTTPStatus(statusCode int) string {
	switch statusCode {
	case http.StatusUnprocessableEntity:
		return "unprocessable entity, likely due to an invalid payload"
	case http.StatusInternalServerError:
		return "error occurred while processing the request"
	case http.StatusBadRequest:
		return "the provided data in the request was not valid, please try again"
	case http.StatusUnauthorized:
		return "unauthorized"
	default:
		return "ok"
	}
}
