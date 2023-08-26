package fibonacci

import (
	"fibonacci-spiral-matrix/api/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GenerateSpiralFibonacci(ctx *gin.Context, rows int, cols int) model.Spiral
}

type service struct {
}

func NewService() Service {
	return &service{}
}

//generate spiral matrix
func (s *service) GenerateSpiralFibonacci(ctx *gin.Context, rows int, cols int) model.Spiral {
	matrix := initializeMatrix(rows, cols)

	top, bottom, left, right := 0, rows-1, 0, cols-1
	direction := 0
	fib := fibonacci(rows * cols)
	idx := 0

	for top <= bottom && left <= right {
		if direction == 0 {
			for i := left; i <= right; i++ {
				matrix[top][i] = fib[idx]
				idx++
			}
			top++
		} else if direction == 1 {
			for i := top; i <= bottom; i++ {
				matrix[i][right] = fib[idx]
				idx++
			}
			right--
		} else if direction == 2 {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = fib[idx]
				idx++
			}
			bottom--
		} else if direction == 3 {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = fib[idx]
				idx++
			}
			left++
		}
		direction = (direction + 1) % 4
	}

	return model.Spiral{
		Timestamp: time.Now().Unix(),
		Rows:      matrix,
	}
}

func initializeMatrix(rows, cols int) [][]int {
	spiral := make([][]int, rows)
	for i := range spiral {
		spiral[i] = make([]int, cols)
	}
	return spiral
}

func fibonacci(n int) []int {
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}
