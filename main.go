package main

import (
	"github.com/gin-gonic/gin"

	"mockAPI/flow"
)

func main() {
	router := gin.Default()
	router.GET("/flows", flow.ListFlowsHandler)
	router.Run()
}
