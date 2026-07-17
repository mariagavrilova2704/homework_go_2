package float

import (
	"fmt"
	"math"
)

// Блок float закрепляет работу с float64:
// дробные числа, проценты, округление, форматирование,
// преобразование float64 в int и простые вычисления.

// Add складывает два дробных числа.
//
// TODO: реализуйте сложение двух входных значений.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract вычитает второе дробное число из первого.
//
// TODO: реализуйте вычитание значения b из значения a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply умножает два дробных числа.
//
// TODO: реализуйте умножение двух входных значений.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide делит первое дробное число на второе.
//
// TODO: реализуйте безопасное деление дробных чисел.
// При невозможности деления функция должна вернуть 0.
func Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// DiscountPrice считает цену после скидки.
//
// TODO: примените процент скидки к цене.
// Некорректная или слишком большая скидка должна обрабатываться безопасно.
func DiscountPrice(price, percent float64) float64 {
	if percent <= 0 || (price*percent/100) > price {
		return price
	}
	return price - (price * percent / 100)
}

// AddTax считает цену с налогом.
//
// TODO: примените процент налога к цене.
// Некорректный налог не должен увеличивать цену.
func AddTax(price, taxPercent float64) float64 {
	if taxPercent <= 0 || price*(1+(taxPercent/100)) < price {
		return price
	}
	return price * (1 + (taxPercent / 100))
}

// CelsiusToFahrenheit переводит градусы Цельсия в Фаренгейты.
//
// TODO: реализуйте перевод температуры по стандартной формуле.
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

// FahrenheitToCelsius переводит градусы Фаренгейта в Цельсии.
//
// TODO: реализуйте обратный перевод температуры по стандартной формуле.
func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

// Average считает среднее арифметическое двух дробных чисел.
//
// TODO: посчитайте среднее значение для двух входных чисел.
func Average(a, b float64) float64 {
	return (a + b) / 2
}

// Round2 округляет число до 2 знаков после точки.
//
// TODO: округлите value до двух знаков после точки.
func Round2(value float64) float64 {
	return math.Round(value*100) / 100
}

// FormatPrice форматирует цену с двумя знаками после точки.
//
// TODO: верните строковое представление цены с двумя цифрами после точки.
func FormatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

// PercentOf считает процент от числа.
//
// TODO: посчитайте указанную процентную часть от total.
func PercentOf(total, percent float64) float64 {
	return total * percent / 100
}

// GrowthPercent считает рост в процентах между старым и новым значением.
//
// TODO: посчитайте процентное изменение между oldValue и newValue.
// Нулевое старое значение должно обрабатываться безопасно.
func GrowthPercent(oldValue, newValue float64) float64 {
	if oldValue == 0 {
		return 0
	}
	return (newValue - oldValue) / oldValue * 100
}

// IsPositive проверяет, что число строго больше нуля.
//
// TODO: определите, является ли значение положительным.
func IsPositive(value float64) bool {
	if value > 0 {
		return true
	}
	return false
}

// FloatToInt преобразует float64 в int.
//
// TODO: выполните явное преобразование к int и проверьте поведение на дробных числах.
func FloatToInt(value float64) int {
	return int(value)
}
