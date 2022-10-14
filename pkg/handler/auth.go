package handler

import (
	"net/http"

	"github.com/asadbek21coder/bookshelf"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input bookshelf.User

	if err := c.Bind(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{

		"data":    data,
		"isOk":    true,
		"message": "ok",
	})
}

func (h *Handler) getUser(c *gin.Context) {
	// var input bookshelf.User
	header := bookshelf.Header{}

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, err.Error())
		return
	}

	data, err := h.services.Authorization.GetUser(header)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "ok",
	})
}
