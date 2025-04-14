package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.MaxInt64)
	sort.Ints([]int{1, 2, 3})
	sort.SliceStable(nil, nil)
	sort.Slice(nil, nil)
	strconv.Itoa(1)
	strconv.Atoi("1")
	strconv.ParseInt("10", 10, 64)
	rand.Intn(1)
	rand.Intn(1)

	dp := make([]int, 10)
	for i := range dp {
		dp[i] = rand.Intn(10)
	}

	var s []byte
	fmt.Println(string(s))

}
