package profile

import (
	"fmt"
	"strings"
	"unicode"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
//
// TODO: верните строку "name=<name> age=<age> group=<group> status=<status>". Имя очистите по краям, приведите к нижнему регистру и сделайте первый символ заглавным; пустое имя замените на "Unknown". Возраст от 18 — adult, иначе minor; active=true — active, иначе inactive.
func BuildUserCard(name string, age int, active bool) string {
	name = strings.TrimSpace(strings.ToLower(name))
	if name == "" {
		name = "Unknown"
	} else {
		r := []rune(name)
		r[0] = unicode.ToUpper(r[0])
		name = string(r)
	}
	group := "minor"
	if age >= 18 {
		group = "adult"
	}
	status := "inactive"
	if active {
		status = "active"
	}
	return fmt.Sprintf("name=%s age=%d group=%s status=%s", name, age, group, status)
}
