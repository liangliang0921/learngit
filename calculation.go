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
	array1 := []float64{0, 575, 575, 0, 575, 575}
	// fmt.Println(Sum(array))
	fmt.Println(array1)
	fmt.Println("平均值:", Average(array1))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array1))

	array2 := []float64{0, 0, 0, 0, 0, 10, 20, 30, 90, 200}
	// fmt.Println(Sum(array))
	fmt.Println(array2)
	fmt.Println("平均值:", Average(array2))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array2))

	array3 := []float64{0, 0, 0, 0, 0, 0, 293, 293, 293, 20}
	// fmt.Println(Sum(array))
	fmt.Println(array3)
	fmt.Println("平均值:", Average(array3))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array3))

	array4 := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 200}
	fmt.Println(array4)
	// fmt.Println(Sum(array))
	fmt.Println("平均值:", Average(array4))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array4))

	array5 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array5)
	// fmt.Println(Sum(array))
	fmt.Println("平均值:", Average(array5))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array5))

	array6 := []float64{20, 30, 40, 50, 20, 33, 27, 28, 19, 33}
	fmt.Println(array6)
	// fmt.Println(Sum(array))
	fmt.Println("平均值:", Average(array6))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array6))

	array7 := []float64{0, 0, 0, 0, 0, 0, 0, 0, 9, 10}
	fmt.Println(array7)
	// fmt.Println(Sum(array))
	fmt.Println("平均值:", Average(array7))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array7))

	array8 := []float64{1, 1, 3, 4}
	fmt.Println(array8)
	// fmt.Println(Sum(array))
	fmt.Println("平均值:", Average(array8))
	// fmt.Println(Variance(array))
	fmt.Println("标准差:", StandardDeviation(array8))

}
