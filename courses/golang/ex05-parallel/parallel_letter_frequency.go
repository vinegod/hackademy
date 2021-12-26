package letter

func Frequency(letter string) map[rune]int {
	characters := make(map[rune]int)

	for _, r := range letter {
		characters[r]++
	}
	return characters
}

func frequency(letter string, mp chan map[rune]int) {
	mp <- Frequency(letter)
}

func ConcurrentFrequency(letters []string) map[rune]int {
	maps := make(chan map[rune]int)

	for i := range letters {
		go frequency(letters[i], maps)
	}

	characters := make(map[rune]int)

	for i := 0; i < len(letters); i++ {
		for r, n := range <-maps {
			characters[r] += n
		}
	}

	return characters
}
