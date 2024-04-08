package main

import (
	"github.com/KevinZonda/repo/repo_collection"
	"github.com/KevinZonda/repo/serv"
	"github.com/gin-gonic/gin"
	"time"
)
import "github.com/gin-contrib/cors"

func main() {
	repo_collection.R.StartDaemon(2 * time.Hour)
	g := serv.Gin()
	cors_cfg := cors.DefaultConfig()
	cors_cfg.AllowAllOrigins = true
	g.Use(cors.New(cors_cfg))
	serv.API(g)
	serv.Html(g)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "not found",
		})
		return
	})
	g.Run("127.0.0.1:18439")
}
