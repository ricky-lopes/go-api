package controller

import (
	"fmt"
	"net/http"
	"teste/src/configuration/logger"
	"teste/src/configuration/validation"
	"teste/src/controller/user/model/request"
	"teste/src/controller/user/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar informações do usuário", err, zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}
