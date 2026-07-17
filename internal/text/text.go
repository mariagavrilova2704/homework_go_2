package text

import "strings"

// Блок text закрепляет работу со строками:
// string, byte, rune, immutable string, Unicode,
// пакет strings и простые операции над текстом.

// ByteLen возвращает длину строки в байтах.
//
// TODO: посчитайте размер строки в байтах.
func ByteLen(s string) int {
	return len(s)
}

// RuneLen возвращает количество Unicode-символов в строке.
//
// TODO: посчитайте количество rune в строке.
func RuneLen(s string) int {
	return len([]rune(s))
}

// FirstRune возвращает первый Unicode-символ строки.
//
// TODO: безопасно верните первый символ строки с учётом Unicode.
func FirstRune(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	return string(r[0])
}

// LastRune возвращает последний Unicode-символ строки.
//
// TODO: безопасно верните последний символ строки с учётом Unicode.
func LastRune(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	return string(r[len(r)-1])
}

// Trim убирает пробелы по краям строки.
//
// TODO: очистите строку от внешних пробельных символов.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToLower переводит строку в нижний регистр.
//
// TODO: приведите строку к нижнему регистру.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper переводит строку в верхний регистр.
//
// TODO: приведите строку к верхнему регистру.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// NormalizeEmail нормализует email.
//
// TODO: очистите email от внешних пробелов и приведите к единому регистру.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// ContainsWord проверяет, содержит ли text подстроку word.
//
// TODO: проверьте наличие word внутри text.
func ContainsWord(text, word string) bool {
	if strings.Contains(text, word) {
		return true
	}
	return false
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
//
// TODO: верните новую строку, где первый символ заменён на r.
// Пустая строка должна обрабатываться безопасно.
func ReplaceFirstRune(s string, r rune) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	rWords := []rune(s)
	rWords[0] = r
	return string(rWords)
}

// ReverseRunes разворачивает строку по Unicode-символам.
//
// TODO: разверните строку без поломки Unicode-символов.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Initials возвращает инициалы имени и фамилии.
//
// TODO: соберите инициалы из имени и фамилии.
// Пробелы по краям не должны влиять на результат.
func Initials(firstName, lastName string) string {
	firstName = strings.ToUpper(strings.TrimSpace(firstName))
	lastName = strings.ToUpper(strings.TrimSpace(lastName))

	if firstName == "" && lastName == "" {
		return ""
	}

	rFirstName := []rune(firstName)
	rLastName := []rune(lastName)

	if lastName == "" {
		return string(rFirstName[0])
	}
	if firstName == "" {
		return string(rLastName[0])
	}
	return string(rFirstName[0]) + string(rLastName[0])
}

// RepeatWord повторяет слово count раз.
//
// TODO: повторите слово нужное количество раз.
// Неположительное количество повторений должно давать пустую строку.
func RepeatWord(word string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(word, count)
}

// JoinWithComma объединяет строки через запятую.
//
// TODO: объедините элементы с разделителем-запятой.
func JoinWithComma(values []string) string {
	return strings.Join(values, ",")
}

// IsPalindrome проверяет, является ли строка палиндромом.
//
// TODO: сравните строку с её развёрнутой версией.
// Регистр и пробелы не должны мешать проверке.
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	reversed := string(r)
	if reversed == s {
		return true
	}
	return false
}
