package infrastructure

import "net/http"

const (
	ItemTypeArticle = 1
	ItemTypeProduct = 2
)

func GetMessageForHTTPStatus(statusCode int) string {
	switch statusCode {
	case http.StatusUnprocessableEntity:
		return "unprocessable entity, likely due to an invalid payload"
	case http.StatusInternalServerError:
		return "error ocurred while processing the request"
	case http.StatusBadRequest:
		return "the provided data in the request was not valid, please try again"
	default:
		return "ok"
	}
}
