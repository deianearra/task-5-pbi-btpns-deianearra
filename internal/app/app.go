package app

import (
	"task-5-pbi-btpns-deianearra/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := handlers.SetupRouter()
	gin.SetMode(gin.ReleaseMode)
	r.Run(":8080")
}
