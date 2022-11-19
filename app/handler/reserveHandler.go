package handler

import (
	"avito-test-task/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
