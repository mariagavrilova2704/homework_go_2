package boolean

// Блок boolean закрепляет работу с bool:
// сравнения, логические операторы &&, ||, !,
// простые условия и функции, которые отвечают true или false.

// CanEnter проверяет, может ли человек войти.
//
// TODO: вход разрешён только совершеннолетнему человеку с билетом.
func CanEnter(age int, hasTicket bool) bool {
	if age >= 18 && hasTicket {
		return true
	}
	return false
}

// IsAdult проверяет, является ли человек совершеннолетним.
//
// TODO: определите совершеннолетие по возрасту.
func IsAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// CanBuyAlcohol проверяет, можно ли купить алкоголь.
//
// TODO: для учебной задачи используйте возрастное ограничение из тестов.
func CanBuyAlcohol(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// CanRest проверяет, можно ли отдыхать.
//
// TODO: отдых возможен в выходной или праздничный день.
func CanRest(isWeekend, isHoliday bool) bool {
	if isWeekend || isHoliday {
		return true
	}
	return false
}

// IsWorkingDay проверяет, является ли день рабочим.
//
// TODO: обработайте английские названия дней недели.
// Неизвестные значения считаются нерабочими.
func IsWorkingDay(day string) bool {
	if day == "monday" || day == "tuesday" || day == "wednesday" || day == "thursday" || day == "friday" {
		return true
	}
	return false
}

// HasAccess проверяет доступ пользователя.
//
// TODO: доступ есть у пользователя с административными правами или у владельца ресурса.
func HasAccess(isAdmin, isOwner bool) bool {
	if isOwner || isAdmin {
		return true
	}
	return false
}

// CanApplyDiscount проверяет, можно ли применить скидку.
//
// TODO: скидка доступна VIP-пользователю или при достаточной сумме заказа.
func CanApplyDiscount(isVIP bool, total int) bool {
	if isVIP || total >= 5000 {
		return true
	}
	return false
}

// ShouldNotify проверяет, нужно ли отправлять уведомление.
//
// TODO: уведомление отправляется только при выполнении обоих входных условий.
func ShouldNotify(emailVerified, notificationsEnabled bool) bool {
	if emailVerified && notificationsEnabled {
		return true
	}
	return false
}

// IsValidScore проверяет корректность оценки.
//
// TODO: оценка должна входить в допустимый диапазон.
func IsValidScore(score int) bool {
	if score >= 0 && score <= 100 {
		return true
	}
	return false
}

// IsInRange проверяет, что value находится в диапазоне [min, max].
//
// TODO: проверьте попадание значения в диапазон вместе с границами.
func IsInRange(value, min, max int) bool {
	if value >= min && value <= max {
		return true
	}
	return false
}

// IsLeapYear проверяет, является ли год високосным.
//
// TODO: реализуйте стандартное правило високосного года.
// Проверьте обычные годы, века и годы, кратные 400.
func IsLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// CanWithdraw проверяет, можно ли снять деньги.
//
// TODO: учитывайте блокировку аккаунта, сумму снятия и доступный баланс.
func CanWithdraw(balance, amount int, blocked bool) bool {
	if balance >= amount && !blocked && amount > 0 {
		return true
	}
	return false
}

// LoginAllowed проверяет, разрешён ли вход.
//
// TODO: вход разрешён только при успешной проверке пароля и второго фактора.
func LoginAllowed(passwordOK, otpOK bool) bool {
	if passwordOK && otpOK {
		return true
	}
	return false
}

// IsEmpty проверяет, является ли строка пустой.
//
// TODO: отличайте пустую строку от строки с пробелами или другими символами.
func IsEmpty(text string) bool {
	if len(text) == 0 {
		return true
	}
	return false
}

// Not возвращает противоположное bool-значение.
//
// TODO: реализуйте логическое отрицание.
func Not(flag bool) bool {
	if !flag {
		return true
	}
	return false
}
