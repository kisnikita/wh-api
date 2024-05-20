package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func castId(c *gin.Context) (int, error) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return -1, errors.New("invalid type of user id")
	}

	return idInt, nil
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
