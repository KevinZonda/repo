package serv

import (
	"github.com/KevinZonda/repo/repo_collection"
	"github.com/KevinZonda/repo/repo_standard"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
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
		repo := seq().FullRepository()

		if c.Query("detail") != "true" {
			copyRepo := repo
			for k, v := range repo.Packages {
				v.History = nil
				copyRepo.Packages[k] = v
			}
			repo = copyRepo
		}
		c.JSON(200, repo)
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
		pkg := repo.Packages[name].History[version].Urls[repo_standard.Platform(platform)][repo_standard.Arch(arch)]
		if pkg == "" {
			c.JSON(404, gin.H{
				"error": "not found",
			})
			return
		}

		proxy := c.Query("proxy")
		if proxy != "true" {
			c.Redirect(302, pkg)
			return
		}
		remote, _ := url.Parse(pkg)
		px := httputil.NewSingleHostReverseProxy(remote)
		px.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL = remote
		}

		px.ServeHTTP(c.Writer, c.Request)
		return

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
		platformName := ""
		switch platfm {
		case "win":
			platformName = "Windows"
		case "mac":
			platformName = "macOS (Darwin)"
		case "linux":
			platformName = "Linux"
		default:
			platformName = "Unknown"
		}
		proxy := c.Query("proxy") == "true"
		full := c.Query("full") == "true"
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"UA":           ua,
			"seq":          seq().Template(),
			"platform":     platfm,
			"platformName": platformName,
			"full":         full,
			"proxy":        proxy,
		})
	})

}
