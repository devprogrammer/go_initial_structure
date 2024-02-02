package httphandler

import (
	"net/http"
	"strconv"
	"strings"
	"userlogin/errors"

	"github.com/go-chi/render"
)

var (
	headerKey  = "AR-Token"
	contextKey = "auth-data-key"
)

type SuccessResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, message string) {
	render.Status(r, http.StatusBadRequest)
	render.Respond(w, r, &ErrorResponse{Error: true, Message: message})
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, message string) {
	render.Status(r, http.StatusNotFound)
	render.Respond(w, r, &ErrorResponse{Error: true, Message: message})
}

func ForbiddenRequestResponse(w http.ResponseWriter, r *http.Request, message string) {
	render.Status(r, http.StatusForbidden)
	render.Respond(w, r, &ErrorResponse{Error: true, Message: message})
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, message string) {
	render.Status(r, http.StatusForbidden)
	render.Respond(w, r, &ErrorResponse{Error: true, Message: message})
}

func Respond(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := err.(*errors.Error); ok {
		render.Status(r, e.Code)
		render.Respond(w, r, &ErrorResponse{Error: true, Message: e.Message})
	} else {
		InternalServerErrorResponse(w, r, err.Error())
	}
}

func OK(w http.ResponseWriter, r *http.Request, message string, data interface{}) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, &SuccessResponse{
		Error:   false,
		Message: message,
		Data:    data,
	})
}

func parseParams(r *http.Request) (int, int, error) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 20
	}

	return page, limit, nil
}

func isMeRequest(r *http.Request) bool {
	return strings.Contains(r.URL.Path, "/me/")
}
