package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ttallskog/sample-golang/gin-gonic/echo"
)

func main() {

	router := gin.New()

	echo.CreateEchoService(router)

	router.Run()
}
