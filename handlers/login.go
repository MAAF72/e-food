package handlers

import (
	"e-food/models"
	"e-food/restapi/operations/user"
	"e-food/utils"
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type userLoginImpl struct{}

// NewUserLoginHandler create new menu.LoginHandler
func NewUserLoginHandler() user.LoginHandler {
	return &userLoginImpl{}
}

func (impl *userLoginImpl) Handle(params user.LoginParams, principal interface{}) middleware.Responder {
	expectedEmail := "admin@live.com"
	expectedPassword := "admin123"
	expectedName := "Administrator"

	if !strings.EqualFold(*params.Login.Email, expectedEmail) {
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}

	if strings.Compare(*params.Login.Password, expectedPassword) != 0 {
		return user.NewRegisterNotFound()
	}

	token, err := utils.GenerateJWT(expectedEmail, expectedName)
	if err != nil {
		fmt.Println(err)
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}

	responseVal := &models.LoginSuccess{
		Success: true,
		Token:   token,
	}

	return user.NewLoginOK().WithPayload(responseVal)
}
