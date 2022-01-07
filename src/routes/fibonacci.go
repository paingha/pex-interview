package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/controllers"
)

const (
	FibonacciRootRoute          = "/v1/fibonacci"
	CurrentFibonacciNumberPath  = "/current"
	NextFibonacciNumberPath     = "/next"
	PreviousFibonacciNumberPath = "/previous"
)

//FibonacciRouter - fibonacci subroutes
func FibonacciRouter(r *gin.Engine, env *controllers.Env) *gin.Engine {
	v1 := r.Group(FibonacciRootRoute)
	{
		v1.GET(CurrentFibonacciNumberPath, env.Current)
		v1.POST(NextFibonacciNumberPath, env.Next)
		v1.POST(PreviousFibonacciNumberPath, env.Previous)
	}
	return r
}
