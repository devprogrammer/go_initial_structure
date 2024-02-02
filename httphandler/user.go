package httphandler

import (
	"net/http"
	"userlogin/services"
	"userlogin/types"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService services.UserService
	logger      *logrus.Logger
}

func NewUserHandler(userService services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := &types.CreateUserPayload{}

	response := make([]*types.User, 0)
	opErrors := make([]string, 0)

	for _, next := range payload.Payload {
		newUser, err := handler.userService.CreateUser(next)
		if err != nil {
			opErrors = append(opErrors, err.Error())
			continue
		}
		response = append(response, newUser)
	}

	// newUser, err := handler.userService.CreateUser(*types.CreateUserRequest)
	// if err != nil {
	// 	opErrors = append(opErrors, err.Error())
	// }
	// response = append(response, newUser)

	type createUserResponse struct {
		Users  []*types.User `json:"users"`
		Errors []string      `json:"errors"`
	}

	OK(w, r, "user created", &createUserResponse{
		Users:  response,
		Errors: opErrors,
	})
}

func (handler *UserHandler) TestAPI(w http.ResponseWriter, r *http.Request) {
	type testResponse struct {
		res string `json:"resp"`
		err string `json:"errors"`
	}
	res := "response"
	err := "no err"
	OK(w, r, "test api", &testResponse{
		res: res,
		err: err,
	})
}
