package conversion

import (
	"fmt"
	"strconv"
	"strings"
)

// Блок conversion закрепляет:
// явные преобразования типов, named types,
// strconv, парсинг строк, форматирование значений
// и обработку ошибок при преобразовании.

// Rub — пользовательский тип для рублей.
type Rub int

// Kop — пользовательский тип для копеек.
type Kop int

// UserID — пользовательский тип для идентификатора пользователя.
type UserID int64

// IntToInt64 явно преобразует int в int64.
//
// TODO: выполните явное преобразование типа.
func IntToInt64(n int) int64 {
	return int64(n)
}

// Int64ToInt явно преобразует int64 в int.
//
// TODO: выполните явное преобразование типа.
// Для этой учебной задачи переполнение не обрабатываем.
func Int64ToInt(n int64) int {
	return int(n)
}

// RubToKop переводит рубли в копейки.
//
// TODO: переведите значение из рублей в копейки и сохраните тип результата.
func RubToKop(rub Rub) Kop {
	return Kop(rub * 100)
}

// KopToRub переводит копейки в рубли через целочисленное деление.
//
// TODO: переведите значение из копеек в рубли и сохраните тип результата.
func KopToRub(kop Kop) Rub {
	return Rub(kop / 100)
}

// UserIDToString преобразует UserID в string.
//
// TODO: преобразуйте пользовательский идентификатор в строку.
func UserIDToString(id UserID) string {
	return strconv.FormatInt(int64(id), 10)
}

// ParseInt преобразует строку в int.
//
// TODO: распарсите строку как int и корректно верните ошибку.
func ParseInt(text string) (int, error) {
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// ParseAndDouble парсит строку в int и умножает результат на 2.
//
// TODO: распарсите строку, обработайте ошибку и верните удвоенное число.
func ParseAndDouble(text string) (int, error) {
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return num * 2, nil
}

// IntToString преобразует int в string.
//
// TODO: преобразуйте целое число в строку.
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// FloatToString форматирует float64 с двумя знаками после точки.
//
// TODO: преобразуйте дробное число в строку с фиксированным количеством знаков после точки.
func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64)
}

// ParseBoolText преобразует строку в bool.
//
// TODO: распарсите строковое bool-значение и корректно верните ошибку.
func ParseBoolText(text string) (bool, error) {
	text = strings.TrimSpace(text)
	b, err := strconv.ParseBool(text)
	if err != nil {
		return false, err
	}
	return b, nil
}

// BoolToText преобразует bool в string.
//
// TODO: преобразуйте bool в его строковое представление.
func BoolToText(value bool) string {
	return strconv.FormatBool(value)
}

// SumIntAndInt64 складывает int и int64.
//
// TODO: сложите значения разных целочисленных типов через явное преобразование.
func SumIntAndInt64(a int, b int64) int64 {
	return int64(a) + b
}

// PriceRubStringToKop преобразует строку с рублями в копейки.
//
// TODO: распарсите цену в рублях, проверьте корректность и верните цену в копейках.
func PriceRubStringToKop(text string) (Kop, error) {
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	if num < 0 {
		return 0, fmt.Errorf("цена не может быть отрицательной: %d", num)
	}
	rub := Rub(num)
	return Kop(rub * 100), nil
}

// SafeParsePositive безопасно парсит положительное число.
//
// TODO: верните положительное число из строки или 0 при некорректном значении.
func SafeParsePositive(text string) int {
	text = strings.TrimSpace(text)
	num, err := strconv.Atoi(text)
	if err != nil || num <= 0 {
		return 0
	}
	return num
}

// FormatUser форматирует пользователя в строку.
//
// TODO: соберите строковое представление пользователя по формату из тестов.
func FormatUser(id UserID, name string) string {
	return "user:" + strconv.FormatInt(int64(id), 10) + ":" + name
}
