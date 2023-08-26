package handler

import (
	"fibonacci-spiral-matrix/api/internal/fibonacci"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Fibonacci struct {
	service fibonacci.Service
}

func NewFibonacci(service fibonacci.Service) *Fibonacci {
	return &Fibonacci{service: service}
}

func (t *Fibonacci) NewSpiral() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rowsParam := ctx.DefaultQuery("rows", "5")
		colsParam := ctx.DefaultQuery("cols", "5")

		rws, err := strconv.Atoi(rowsParam)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		cls, err := strconv.Atoi(colsParam)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		result := t.service.GenerateSpiralFibonacci(ctx, rws, cls)
		ctx.JSON(http.StatusOK, result)
	}
}
