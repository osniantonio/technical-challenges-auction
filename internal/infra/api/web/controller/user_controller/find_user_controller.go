package user_controller

import (
	"context"
	"fullcycle-auction_go/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
