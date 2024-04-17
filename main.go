package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("home", "templates/base.html", "templates/index/subtitle.html", "templates/index/paragraph.html")
	r.AddFromFiles("article", "templates/base.html", "templates/article/subtitle.html", "templates/article/paragraph.html")
	return r
}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	r.HTMLRender = createMyRender()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/other", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "other page",
	// 	})
	// })

	// r.GET("/templates/app", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index", gin.H{"name": "app.tmpl"})
	// })

	// r.GET("/templates/base", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "base", gin.H{"name": "base.tmpl"})
	// })

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	r.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	r.Run("127.0.0.1:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
