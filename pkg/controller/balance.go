package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Controllers) GetById(c *gin.Context) {
	var userid, err = getUserId(c)
	if err != nil {
		return
	}

	balanceId, err := h.services.GetById(userid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balanceId.Balance,
	})

}
