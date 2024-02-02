package store

import (
	"context"
	"userlogin/types"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	// methods to access to DB
	CreateUser(request *types.CreateUserRequest) (*types.User, error)
}

type userStore struct {
	client *mongo.Client
	ctx    context.Context
	logger *logrus.Logger
}

func NewUserStore(client *mongo.Client, ctx context.Context, logger *logrus.Logger) UserStore {
	return &userStore{client: client, ctx: ctx, logger: logger}
}

func (p *userStore) CreateUser(request *types.CreateUserRequest) (*types.User, error) {
	user := types.NewUser(request)
	// tx := p.db.Begin()

	// if err := tx.Error; err != nil {
	// 	p.logger.WithError(err).Error("failed to obtain bd txn")
	// }
	return user, nil
}

// func ModifyUser(userId string, user *types.User) (*types.User, error) {
// 	// return nil, error
// }
