package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/paingha/pex/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	SwaggerRootRoute = "/swagger"
)

//SwaggerRouter - swagger subroutes
func SwaggerRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group(SwaggerRootRoute)
	{
		v1.GET("*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return r
}
