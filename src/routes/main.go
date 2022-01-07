package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/controllers"
)

//SetupRouter - setup routes for api
func SetupRouter(env *controllers.Env) *gin.Engine {
	r := gin.Default()
	{
		//RootRouter(r)
		FibonacciRouter(r, env)
	}
	return r
}
