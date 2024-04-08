package serv

import (
	"fmt"
	"github.com/KevinZonda/repo/repo_collection"
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"html/template"
	"net/http"
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

func Html(r *gin.Engine) {
	r.FuncMap = template.FuncMap{
		"notnil": func(a any) bool {
			return a != nil
		},
	}

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		ua := c.Request.UserAgent()
		ua_s := useragent.Parse(ua)

		platfm := c.Query("platform")
		switch platfm {
		case "win", "mac", "linux":
		default:
			if ua_s.IsWindows() || ua_s.IsAndroid() {
				platfm = "win"
			} else if ua_s.IsMacOS() || ua_s.IsIOS() {
				platfm = "mac"
			} else if ua_s.IsLinux() {
				platfm = "linux"
			} else {
				platfm = "win"
			}
		}
		fmt.Println(platfm)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":    "Main website",
			"UA":       ua,
			"seq":      seq().Template(),
			"platform": platfm,
		})
	})

}
