package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessJSONResponse(jsonResposeData map[string]interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": jsonResposeData})
}

func CheckHttpBadRequest(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func HandleServerError(err error, c *gin.Context) {
	fmt.Println("errrrrrrrrr", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
