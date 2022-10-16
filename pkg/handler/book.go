package handler

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/bookshelf"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBook(c *gin.Context) {

	var isbn bookshelf.Isbn
	header := bookshelf.Header{}

	if err := c.Bind(&isbn); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, err.Error())
		return
	}

	data, err := h.services.Book.CreateBook(isbn, header)
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

func (h *Handler) getAllBooks(c *gin.Context) {
	header := bookshelf.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, err.Error())
		return
	}

	data, err := h.services.Book.GetAllBooks(header)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// fmt.Println(data)

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"isOk":    true,
		"message": "ok",
	})

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {
	id := c.Param("id")
	header := bookshelf.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, err.Error())
		return
	}
	fmt.Println(header)
	fmt.Println(id)

	c.JSON(http.StatusOK, map[string]interface{}{
		// "data":    data,
		"isOk":    true,
		"message": "ok",
	})
}
