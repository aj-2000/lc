package dsa

import (
	"fmt"
	"strconv"
)

func AI2AS(a []int) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = strconv.Itoa(a[i])
	}
	return result
}

func AS2AI(a []string) ([]int, error) {
	result := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		num, err := strconv.Atoi(a[i])
		if err != nil {
			return nil, fmt.Errorf("failed to convert %s to int", a[i])
		} else {
			result[i] = num
		}
	}
	return result, nil
}

func AAI2AAS(a [][]int) [][]string {
	result := make([][]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = AI2AS(a[i])
	}
	return result
}

func AAS2AAI(a [][]string) ([][]int, error) {
	result := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		nums, err := AS2AI(a[i])
		if err != nil {
			return nil, err
		} else {
			result[i] = nums
		}
	}
	return result, nil
}
