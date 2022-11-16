package handlers

import (
	"e-food/models"
	"e-food/restapi/operations/user"

	"github.com/go-openapi/runtime/middleware"
)

type userRegisterImpl struct{}

// NewUserRegisterHandler create new menu.RegisterHandler
func NewUserRegisterHandler() user.RegisterHandler {
	return &userRegisterImpl{}
}

func (impl *userRegisterImpl) Handle(params user.RegisterParams, principal interface{}) middleware.Responder {
	responseVal := &models.RegisterSuccess{
		Success: true,
		Message: "OK",
	}

	return user.NewRegisterOK().WithPayload(responseVal)
}
