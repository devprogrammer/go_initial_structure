package services

import (
	"context"
	"userlogin/store"
	"userlogin/types"
	"userlogin/utils"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	// methods to access a function of store.
	CreateUser(request *types.CreateUserRequest) (*types.User, error)
}

type userService struct {
	client *mongo.Client
	ctx    context.Context
	store  store.UserStore
	logger *logrus.Logger
}

func NewUserService(client *mongo.Client, ctx context.Context, userStore store.UserStore, logger *logrus.Logger) UserService {
	return &userService{
		client: client,
		ctx:    ctx,
		store:  userStore,
		logger: logger,
	}
}

func (p *userService) CreateUser(request *types.CreateUserRequest) (*types.User, error) {
	if err := utils.ValidateEmail(request.Email); err != nil {
		p.logger.WithField("email", request.Email).WithError(err).Error("invalid email address")
	}

	// check whether this user is existed or not
	// request.Password, err := password.Generate(14, 1, 1, false, false)
	// if err != nil {
	// 	p.logger.WithError(err).Error("failed to super org")
	// 	return nil, errors.New(http.StatusInternalServerError, "Couldn't generate password")
	// }
	user, _ := p.store.CreateUser(request)
	return user, nil
}
