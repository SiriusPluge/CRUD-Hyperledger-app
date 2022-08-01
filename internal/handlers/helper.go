package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func respError(c *gin.Context, err error) {
	logrus.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func respOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
