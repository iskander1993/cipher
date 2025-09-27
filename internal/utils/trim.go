package utils

// Убираем пробелы в начале и конце строки
func TrimSpace(s string) string {
	start := 0
	end := len(s) - 1

	//Двигаемся с начала строки, пока встречаем пробелы
	for start <= end && s[start] == ' ' {
		start++
	}
	//Двигаемся с конца строки, пока встречаем пробелы
	for end >= start && s[end] == ' ' {
		end--
	}

	//Если строка пустая - вернем ""
	if start > end {
		return ""
	}
	return s[start : end+1]
}
