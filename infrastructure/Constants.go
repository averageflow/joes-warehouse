package infrastructure

import "net/http"

func GetMessageForHTTPStatus(statusCode int) string {
	switch statusCode {
	case http.StatusUnprocessableEntity:
		return "unprocessable entity, likely due to an invalid payload"
	case http.StatusInternalServerError:
		return "error ocurred while processing the request"
	default:
		return "ok"
	}
}
