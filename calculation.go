package main

import (
	"fmt"
	"math"
)

// 求和
func Sum(array []float64) float64 {
	var sum float64
	for _, value := range array {
		sum += value
	}
	return sum
}

// 求平均值
func Average(array []float64) float64 {
	var avg float64
	if len(array) == 0 {
		return avg
	}
	sum := Sum(array)
	avg = sum / float64(len(array))
	return avg
}

// 求方差
func Variance(array []float64) float64 {
	var vari float64
	if len(array) == 0 {
		return vari
	}
	avg := Average(array)
	var variSum float64
	for _, value := range array {
		variSum += math.Pow(value-avg, 2)
	}
	vari = variSum / float64(len(array))

	return vari
}

// 求标准差
func StandardDeviation(array []float64) float64 {
	var stan float64
	if len(array) == 0 {
		return stan
	}
	vari := Variance(array)
	stan = math.Sqrt(vari)
	return stan
}

func main() {
	array := []float64{2, 2, 2, 2}
	fmt.Println(Sum(array))
	fmt.Println(Average(array))
	fmt.Println(Variance(array))
	fmt.Println(StandardDeviation(array))
}
