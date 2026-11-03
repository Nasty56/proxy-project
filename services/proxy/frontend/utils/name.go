package utils

import "strings"

// Fullname формирует полное имя в формате "Фамилия Имя Отчество"
func Fullname(lastName, firstName, middleName string) string {
	return strings.TrimSpace(lastName + " " + firstName + " " + middleName)
}

// NameInitials формирует инициалы в формате "Фамилия И. О."
func NameInitials(fullName string) string {
	parts := strings.Fields(fullName)
	if len(parts) < 1 {
		return fullName
	}
	if len(parts) == 1 {
		return parts[0]
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			result += " " + string(parts[i][0]) + "."
		}
	}
	return result
}