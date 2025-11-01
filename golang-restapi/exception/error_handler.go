package exception

import (
	"golang-restapi/helper"
	"golang-restapi/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writter, request, err) {
		return
	}
	if validationError(writter, request, err) {
		return
	}
	internalServerError(writter, request, err)
}

func validationError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writter.Header().Add("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResonse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST ERROR",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writter, webResonse)
		return true
	} else {
		return false
	}
}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Add("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResonse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND ERROR",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writter, webResonse)
		return true
	} else {
		return false
	}
}

func internalServerError(writter http.ResponseWriter, request *http.Request, err interface{}) {
	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResonse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writter, webResonse)
}
