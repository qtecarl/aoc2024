package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	dat, err := os.ReadFile("./input.txt")
    check(err)

	l1 := make([]int, 1000)
	l2 := make([]int, 1000)

	rows := strings.Split(string(dat), "\n")
	for _, row := range rows {
		nums := strings.Split(row, "   ")

		v1, err := strconv.Atoi(nums[0])
		check(err)
		l1 = append(l1, v1)

		v2, err := strconv.Atoi(nums[1])
		check(err)
		l2 = append(l2, v2)
	}
	sort.Slice(l1, func(i, j int) bool { return l1[i] < l1[j] })
	sort.Slice(l2, func(i, j int) bool { return l2[i] < l2[j] })
	
	res := []float64{}
	for i, _ := range l1 {
		res = append(res, math.Abs(float64(l1[i] - l2[i])))
	}
	
	final_value := 0.0
	for _, v := range res {
		final_value += v
	}
	println(int(final_value))
}

