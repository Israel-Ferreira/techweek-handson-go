package utils

import (
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/exceptions"
)

func GetStatusCodeFromError(err error) int {

	switch err {
	case exceptions.ErrorNotFoundItem:
		return http.StatusNotFound
	case exceptions.ErrorInvalidQty:
	case exceptions.ErrorInvalidSKU:
	case exceptions.ErrorInvalidParam:
	case exceptions.ErrorInvalidTitle:
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError

}
