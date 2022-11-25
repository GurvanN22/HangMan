package Hangman

func InitToadd(array []string) []string { //Here we create the "empty" array with "_" we replace letters
	var result []string
	for range array {
		result = append(result, "_")
	}
	return result
}
