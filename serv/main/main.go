package main

import (
	"github.com/KevinZonda/repo/repo_collection"
	"github.com/KevinZonda/repo/serv"
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
	g.Run("127.0.0.1:8080")
}
