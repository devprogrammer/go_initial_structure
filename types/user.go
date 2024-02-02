package types

import (
	"time"
	"userlogin/utils"

	"github.com/google/uuid"
)

type TokenType string

type User struct {
	ID        uuid.UUID `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"-" gorm:"password"`
	Phone     string    `json:"phone" gorm:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

type Token struct {
	ID        uuid.UUID `json:"id"`
	UserId    string    `json:"user_id"`
	Token     string    `json:"token" gorm:"token"`
	TokenType TokenType `json:"token_type" gorm:"token_type"`
	User      *User     `json:"-" gorm:"-"`
	CreatedAt time.Time `json:"-" gorm:"create_at"`
}

func NewToken(user *User, token string, tokentype TokenType) *Token {
	if token == "" {
		token = utils.RandomString(64)
	}
	return &Token{
		ID:        uuid.New(),
		UserId:    user.ID.String(),
		Token:     token,
		TokenType: tokentype,
		User:      user,
		CreatedAt: time.Now(),
	}
}

func NewUser(req *CreateUserRequest) *User {
	password := req.Password
	if password != "" {
		password = utils.HashPassword(req.Password)
	}
	return &User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}
}

// type UserPreference struct {
// 	ID                     uuid.UUID `json:"id"`
// 	UserId                 string    `json:"user_id"`
// 	NotifyOnNewUsers       bool      `json:"notify_on_new_users" gorm:"notify_on_new_users"`
// 	NotifyOnNewOperations  bool      `json:"notify_on_new_operations" gorm:"notify_on_new_operations"`
// 	NotifyOnTransactions   bool      `json:"notify_on_transactions" gorm:"notify_on_transactions"`
// 	NotifyOnNewPortfolio   bool      `json:"notify_on_new_portfolio" gorm:"notify_on_new_portfolio"`
// 	NotifyOnLatestActivity bool      `json:"notify_on_latest_activity" gorm:"notify_on_latest_activity"`
// }

// type Preference struct {
// 	NotifyOnNewUsers       bool `json:"notify_on_new_users"`
// 	NotifyOnNewOperations  bool `json:"notify_on_new_operations"`
// 	NotifyOnTransactions   bool `json:"notify_on_transactions"`
// 	NotifyOnNewPortfolio   bool `json:"notify_on_new_portfolio"`
// 	NotifyOnLatestActivity bool `json:"notify_on_latest_activity"`
// }

// func NewUserPreference(userId string, pref *Preference) *UserPreference {
// 	return &UserPreference{
// 		UserId:                 userId,
// 		NotifyOnNewUsers:       pref.NotifyOnNewUsers,
// 		NotifyOnNewOperations:  pref.NotifyOnNewOperations,
// 		NotifyOnTransactions:   pref.NotifyOnTransactions,
// 		NotifyOnNewPortfolio:   pref.NotifyOnNewPortfolio,
// 		NotifyOnLatestActivity: pref.NotifyOnLatestActivity,
// 	}
// }

// func (*UserPreference) TableName() string {
// 	return "preferences"
// }

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	u.ID = uuid.New()
// 	return nil
// }
