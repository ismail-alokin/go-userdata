package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-userdata/api/users"
)

func main() {
	router := gin.Default()
	router.POST("/users", users.GetUserInformationList)

	router.Run(":8081")
}
