package Hangman

func AddRandomLtrs(entry []string, word []string) ([]string, int, []string) {
	/*function to add a few random letters to the word according to its length*/
	var remain []string
	len := len(entry)
	numberOfDisplayLetters := len/2 - 1
	numberOfTrueValues := 0
	var randomIndex int
	for i := 0; i < numberOfDisplayLetters; i++ {
		randomIndex = Random(entry)
		letter := word[randomIndex]
		if IsIn(entry, letter) == 0 {
			numberOfOccurences := IsIn(word, letter)
			numberOfTrueValues += numberOfOccurences
			entry = Switch(word, entry, letter)
			remain = append(remain, letter)
		}
	}
	return entry, numberOfTrueValues, remain
}
