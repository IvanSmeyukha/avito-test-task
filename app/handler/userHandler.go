package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type amountOfMoney struct {
	Amount float32 `json:"amount"`
}

func (a *amountOfMoney) get() float32 {
	return a.Amount
}

func (h *Handler) addMoneyToUserBalance(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input amountOfMoney
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.User.AddMoney(id, input.get())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusOK)
	return
}

func (h *Handler) getUserBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	balance, err := h.services.User.GetBalance(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, balance)
	return
}
