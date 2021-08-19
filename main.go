package main

import (
	"github.com/gin-gonic/gin"

	"mockAPI/flow"
)

func main() {
	router := gin.Default()
	router.POST("/flows", flow.ListFlowsHandler)
	router.Run()
}
