package dto

import validation "github.com/go-ozzo/ozzo-validation"

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (request UserRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}
