package cipher

import "strings"

// Caesar шифрует или дешифрует текст по алгоритму Цезаря.
// shift > 0 — сдвиг вперёд (шифрование)
// shift < 0 — сдвиг назад (дешифровка)
func Caesar(input string, shift int) string {
	var result strings.Builder
	shift = ((shift % 26) + 26) % 26 // нормализуем сдвиг

	for _, r := range input {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune('a' + (r-'a'+rune(shift))%26)
		case r >= 'A' && r <= 'Z':
			result.WriteRune('A' + (r-'A'+rune(shift))%26)
		default:
			result.WriteRune(r)
		}
	}
	return result.String()
}
