package tools

import (
	"math/rand"
	"time"
)

func RandArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		// if rand.Intn(3) < 1 {
		// 	arr[i] = -rand.Intn(100)
		// 	continue
		// }
		arr[i] = rand.Intn(100)
	}
	return arr
}
