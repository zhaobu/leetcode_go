package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getIndex() []int {
	ret := []int{0, 1, 2, 3, 4}
	for i := len(ret); i > 0; i-- {
		j := rand.Intn(i)
		ret[j], ret[i-1] = ret[i-1], ret[j]
	}

	return ret[len(ret)-3:]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	result := make([][]byte, 0, 1000000)
	arr := []byte{0, 0, 0, 0, 0}
	count := [2]byte{5, 0}
	for {
		indexs := getIndex()
		for i := 0; i < 3; i++ {
			j := indexs[i]
			if arr[j] == 0 {
				count[0]--
				count[1]++
				arr[j] = 1
			} else {
				count[0]++
				count[1]--
				arr[j] = 0
			}
		}
		ret := make([]byte, 5)
		copy(ret, arr)
		result = append(result, ret)
		if count[1] == 5 {
			for i := 0; i < len(result); i++ {
				fmt.Printf("res=%+v\n", result[i])
			}
			break
		}
		if len(result) > 1000000 {
			fmt.Printf("break")
			break
		}
	}

}
