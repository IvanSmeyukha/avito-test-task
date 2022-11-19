package handler

import (
	"avito-test-task/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) reserveMoneyFromUserBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var trans models.Transaction
	err := c.Bind(&trans)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	trans.UserId = id
	err = h.services.Reserve.ReserveMoney(trans)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}

func (h *Handler) writeOffRevenue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var trans models.Transaction
	err := c.Bind(&trans)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	trans.UserId = id
	err = h.services.Reserve.WriteOffRevenue(trans)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}
