package serv

import (
	"github.com/KevinZonda/repo/repo_collection"
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/gin-gonic/gin"
)

func Gin() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func seq() repo_collection.Collection {
	return repo_collection.R
}

func API(r gin.IRouter) {
	r.GET("/package", func(c *gin.Context) {
		c.JSON(200, seq().FullRepository())
	})
	r.GET("/package/:name/:version/:platform/:arch", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(404, gin.H{
					"error": "not found",
				})
			}
		}()
		name := c.Param("name")
		platform := c.Param("platform")
		arch := c.Param("arch")
		version := c.Param("version")
		repo := seq().FullRepository()
		pkg := repo.Packages[name].Versions[version].Urls[repo_standard.Platform(platform)][repo_standard.Arch(arch)]
		c.Redirect(302, pkg)
	})
}
