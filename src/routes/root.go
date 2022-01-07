package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RootRouter - root subroutes
func RootRouter(r *gin.Engine) *gin.Engine {
	r.Static("/public", "./public")
	root := r.Group("/")
	v1 := r.Group("/v1")
	r.LoadHTMLGlob("templates/root/*.html")
	{
		root.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	return r
}
