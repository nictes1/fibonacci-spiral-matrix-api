package routes

import (
	"fibonacci-spiral-matrix/api/handler"
	"fibonacci-spiral-matrix/api/internal/fibonacci"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}
type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	//datastore
}

func NewRouter(r *gin.Engine) Router {
	return &router{r: r}
}

func (r *router) MapRoutes() {
	r.buildRoutes()
}

func (r *router) buildRoutes() {
	fibonacciSrv := fibonacci.NewService()
	handler := handler.NewFibonacci(fibonacciSrv)

	r.r.GET("/spiral", handler.NewSpiral())
}
