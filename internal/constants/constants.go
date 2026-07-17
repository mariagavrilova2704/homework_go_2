package constants

// Блок constants закрепляет:
// const, iota, enum-like значения, switch,
// статусы, роли, приоритеты, типы событий и простые текстовые маппинги.

// TODO: задайте имя приложения как константу.
// Ожидаемое значение проверяется unit- и integration-тестами.
const appName = "go-homework-2"

// TODO: задайте максимальное количество попыток как константу.
const maxAttempts = 3

// Статусы оплаты.
//
// TODO: оформите связанные статусы через const-блок.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// Роли пользователя.
//
// TODO: оформите связанные роли через const-блок.
const (
	RoleGuest = iota
	RoleUser
	RoleAdmin
)

// Приоритеты задачи.
//
// TODO: оформите приоритеты так, чтобы их значения шли по возрастанию важности.

const (
	PriorityLow    = 1
	PriorityMedium = 2
	PriorityHigh   = 3
)

// Типы событий.
//
// TODO: оформите связанные типы событий через const-блок.
const (
	EventCreated = iota
	EventUpdated
	EventDeleted
)

// AppName возвращает имя приложения.
//
// TODO: верните значение соответствующей константы.
func AppName() string {
	return appName
}

// MaxAttempts возвращает максимальное количество попыток.
//
// TODO: верните значение соответствующей константы.
func MaxAttempts() int {
	return maxAttempts
}

// StatusText возвращает текстовое представление статуса.
//
// TODO: преобразуйте известные статусы в текст, неизвестные — в "unknown".
func StatusText(status int) string {
	switch status {
	case int(StatusNew):
		return "new"
	case int(StatusPaid):
		return "paid"
	case int(StatusCanceled):
		return "canceled"
	default:
		return "unknown"
	}
}

// IsFinalStatus проверяет, является ли статус финальным.
//
// TODO: определите, какие статусы завершают жизненный цикл оплаты.
func IsFinalStatus(status int) bool {
	if status == int(StatusPaid) || status == int(StatusCanceled) {
		return true
	}
	return false
}

// NextStatus возвращает следующий статус.
//
// TODO: реализуйте переход из нового статуса в оплаченный.
// Финальные и неизвестные статусы обработайте безопасно.
func NextStatus(status int) int {
	if status == int(StatusPaid) || status == int(StatusCanceled) {
		return status
	}
	if status == int(StatusNew) {
		return int(status) + 1
	}
	return 0
}

// RoleText возвращает текстовое представление роли.
//
// TODO: преобразуйте известные роли в текст, неизвестные — в "unknown".
func RoleText(role int) string {
	switch role {
	case int(RoleGuest):
		return "guest"
	case int(RoleUser):
		return "user"
	case int(RoleAdmin):
		return "admin"
	default:
		return "unknown"
	}
}

// CanEdit проверяет, может ли пользователь редактировать данные.
//
// TODO: разрешите редактирование только роли с максимальными правами.
func CanEdit(role int) bool {
	if role == int(RoleAdmin) {
		return true
	}
	return false
}

// HTTPStatusText возвращает текст HTTP-статуса.
//
// TODO: обработайте основные HTTP-коды из тестов и неизвестный код.
func HTTPStatusText(code int) string {
	switch code {
	case 100:
		return "Continue"
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 300:
		return "Multiple choices"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal server error"
	default:
		return "Unknown"
	}
}

// DayType возвращает тип дня недели.
//
// TODO: определите рабочие и выходные дни по номеру дня.
// Для этой задачи считаем: 1 — понедельник, 7 — воскресенье.
func DayType(day int) string {
	if day == 0 {
		return "unknown"
	}
	if day == 6 || day == 7 {
		return "weekend"
	}
	return "working"
}

// PriorityText возвращает текстовое представление приоритета.
//
// TODO: преобразуйте известные приоритеты в текст, неизвестные — в "unknown".
func PriorityText(priority int) string {
	switch priority {
	case int(PriorityLow):
		return "low"
	case int(PriorityMedium):
		return "medium"
	case int(PriorityHigh):
		return "high"
	default:
		return "unknown"
	}
}

// IsKnownStatus проверяет, известен ли статус.
//
// TODO: проверьте, входит ли статус в набор объявленных статусов.
func IsKnownStatus(status int) bool {
	switch status {
	case int(StatusNew), int(StatusPaid), int(StatusCanceled):
		return true
	default:
		return false
	}
}

// PaymentStateText возвращает текст состояния оплаты.
//
// TODO: определите состояние оплаты по двум флагам.
// Отмена должна иметь приоритет над оплатой.
func PaymentStateText(paid, canceled bool) string {
	if paid && canceled {
		return "canceled"
	}
	if paid {
		return "paid"
	}
	if canceled {
		return "canceled"
	}
	return "pending"
}

// TrafficLightAction возвращает действие по цвету светофора.
//
// TODO: преобразуйте известные цвета светофора в действие.
func TrafficLightAction(color string) string {
	switch color {
	case "red":
		return "stop"
	case "yellow":
		return "wait"
	case "green":
		return "go"
	default:
		return "unknown"
	}
}

// GradeText возвращает текстовую оценку по score.
//
// TODO: преобразуйте score в текстовую оценку по диапазонам из тестов.
func GradeText(score int) string {
	switch {
	case score > 0 && score <= 30:
		return "retry"
	case score > 30 && score <= 60:
		return "passed"
	case score > 60 && score <= 80:
		return "good"
	case score > 80 && score <= 100:
		return "excellent"
	default:
		return "invalid"
	}
}

// EventTypeText возвращает текстовое представление типа события.
//
// TODO: преобразуйте известные типы событий в текст, неизвестные — в "unknown".
func EventTypeText(eventType int) string {
	switch eventType {
	case int(EventCreated):
		return "created"
	case int(EventUpdated):
		return "updated"
	case int(EventDeleted):
		return "deleted"
	default:
		return "unknown"
	}
}
