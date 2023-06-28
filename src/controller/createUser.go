package controller

import (
	"fmt"

	"github.com/EmesDev/GoLang-CRUD/src/configuration/rest_err"
	"github.com/EmesDev/GoLang-CRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error\n", err.Error()))

		c.JSON(restErr.Code, restErr)

		return
	}

	fmt.Println(userRequest)
}
