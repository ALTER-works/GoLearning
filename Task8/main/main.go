package main

import (
	"Task8/utils"
	"fmt"
)

func main() {
	fmt.Println("Добро пожаловать в calc app!")

	// Возведение в степень
	var a float32 = 2
	fmt.Println("Результат возведение 2 в 5-ую степень:", utils.Exponentiation(a, 5))

	// Арифметическая прогрессия
	var b float32 = 3
	fmt.Println("Результат вычисления 7 числа в арифм прогрессии, начинающийся с 3 и с шагом в 5:",utils.ArithmeticProgression(b, 5, 7))

	var c float32 = 4
	fmt.Println("Результат вычисления суммы арифм прогрессии, начинающийся с 4 и с шагом в 5:",utils.ArithmeticProgressionSum(c, 5, 7))

	// Геометрическая прогрессия
	var d float32 = 5
	fmt.Println("Результат вычисления 7 числа в геометрической прогрессии, начинающийся с 5 и с шагом в 5:",utils.GeometryProgression(d, 5, 7))

	var e float32 = 6
	fmt.Println("Результат вычисления суммы геометрической прогрессии, начинающийся с 6 и с шагом в 5:",utils.GeometryProgression(e, 5, 7))

	// Бесконечные прогрессии через го рутины
	var firstValue float32 = 7
	var steps float32 = 5
	go utils.ArithmeticProgressionInfinite(firstValue, steps)

	
}
