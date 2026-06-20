package main

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for row := 0; row < 8; row++ {
		for _, ch := range text {
			result[row] += banner[ch][row]
		}
	}
	return result
}
