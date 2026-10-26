package utils

// GeometryProgression - базовое вычесление н-ного (step) элемента геометрической прогрессии
func GeometryProgression(startNumber float32, acceleration float32, step float32) float32 {
	newNumber := startNumber * Exponentiation(acceleration, step)
	return newNumber
}

// GeometryProgressionSum - базовое вычесление суммы геометрической прогрессии до элемента (step)
func GeometryProgressionSum(startNumber float32, acceleration float32, step float32) float32 {
	switch {
	case acceleration == 1:
		return startNumber
	case (acceleration < 1) && (acceleration > -1):
		sum := startNumber / (1 - acceleration)
		return sum
	default:
		sum := (startNumber * (Exponentiation(acceleration, step) - 1)) / (acceleration - 1)
		return sum
	}
}

// Exponentiation - возведение числа (number) в степень (power)
func Exponentiation(number float32, power float32) float32 {
	switch {
	case power >= 1:
		result := number
		for i := float32(1); i < power; i++ {
			result = result * number
		}

		return result
	case power <= -1:
		result := number
		for i := float32(1); i < power; i++ {
			result = result * number
		}

		return 1/result
	case power < 1 && power > 0:
		
		return 1
	case power < 0 && power > -1:

		return 1
	default:
		return 1
	}
}

// NRootOfValue - функция нахождения корня числа по степени
// func NRootOfValue()

// ArithmeticProgression - базовое вычесление н-ного (step) элемента арифметической прогрессии
func ArithmeticProgression(startNumber float32, acceleration float32, step int) float32 {
	newNumber := startNumber + acceleration*float32(step-1)
	return newNumber
}

// ArithmeticProgressionSum - базовое вычесление суммы арифметической прогрессии до элемента (step)
func ArithmeticProgressionSum(startNumber float32, acceleration float32, step int) float32 {
	sum := (startNumber + (startNumber + acceleration*float32(step-1))) * float32(step) / 2
	return sum
}

func ArithmeticProgressionInfinite(startNumber float32, acceleration float32) (func() float32, float32) {
	startNumber = startNumber + acceleration
	return ArithmeticProgressionInfinite(startNumber, acceleration)
}


// GeometryProgressionInfinite - бесконечное рекурсивное вычисление прогрессии с определённым шагом и последующий вывод значений
func GeometryProgressionInfinite(startNumber float32, acceleration float32) <-chan ProgressionInfo {
	infochannel := make(chan ProgressionInfo)
	step := 1
	
	go func() {
		infochannel <- ProgressionInfo{startNumber, step}
		stepGeometryInfinite(startNumber, acceleration, step, infochannel)
	}()

	return infochannel
}
// stepGeometryInfinite - шаг в рекурсию для GeometryProgressionInfinite
func stepGeometryInfinite(startNumber float32, acceleration float32, step int, infochannel chan ProgressionInfo) {
	newSN := startNumber * acceleration
	step += 1
	infochannel <- ProgressionInfo{newSN, step}
	stepGeometryInfinite(newSN, acceleration, step, infochannel)
}
type ProgressionInfo struct {
	CurrentValue float32 // текущее значение
	StepCount    int // например, номер шага (или любое другое второе число)
}