package types

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	// Role             string                     `json:"role"`
	CreatedBy string `json:"created_by"`
	// TwoFactorEnabled bool                       `json:"two_factor_enabled"`
	// Organization     *CreateOrganizationRequest `json:"organization"`
}

type CreateUserPayload struct {
	Payload []*CreateUserRequest `json:"payload"`
}
