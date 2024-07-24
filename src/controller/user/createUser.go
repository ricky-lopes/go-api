package controller

import (
	"fmt"
	"teste/src/configuration/validation"
	"teste/src/controller/user/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
