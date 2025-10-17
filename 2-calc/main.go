package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operations = getOperations()
	var selectedOperation string
	var numbers string

	for {
		fmt.Printf("Выберите операцию: %v\n", operations)
		_, err := fmt.Scan(&selectedOperation)
		if err != nil {
			continue
		}

		var isTrue = false
		for _, value := range operations {
			if selectedOperation == value {
				isTrue = true
				break
			}
		}
		if isTrue {
			break
		}
	}

	for {
		fmt.Print("Введите числа: ")
		_, err := fmt.Scan(&numbers)
		if err != nil {
			continue
		}
		break
	}
	nums := parseFloatArray(numbers)
	if selectedOperation == operations[0] {
		fmt.Printf("AVG: %.2f\n", getAVG(nums))
	} else if selectedOperation == operations[1] {
		fmt.Printf("SUM: %.2f\n", getSum(nums))
	} else if selectedOperation == operations[2] {
		fmt.Printf("MED: %.2f\n", getMed(nums))
	}
}

func parseFloatArray(s string) []float64 {
	var numStr = strings.Split(strings.TrimSpace(s), ",")
	var nums []float64
	for _, num := range numStr {
		val, _ := strconv.ParseFloat(num, 64)
		nums = append(nums, val)
	}
	return nums
}

func getSum(num []float64) float64 {
	var sum = 0.0
	for _, value := range num {
		sum += value
	}
	return sum
}

func getAVG(nums []float64) float64 {
	var sum = getSum(nums)
	return sum / float64(len(nums))

}

func getMed(nums []float64) float64 {
	sort.Float64s(nums)

	if len(nums)%2 == 1 {
		return nums[len(nums)/2]
	} else {
		return 0.5 * (nums[len(nums)/2-1] + nums[len(nums)/2])
	}
}

func getOperations() []string {
	return []string{
		"AVG",
		"SUM",
		"MED",
	}
}
