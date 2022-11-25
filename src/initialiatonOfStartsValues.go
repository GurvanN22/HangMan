package src

import "fmt"

func InitialiatonOfStartsValues(nameDB string, displayType string) {
	var ( //We initialise variables
		remain     []string
		usedValues []string
		mistakes   int
		truevValue int
		alphabet   []string
	)
	array := ArrayInit(nameDB)             //We use ArrayInit fonction to made an array from a txt data base
	randomNb := Random(array)              // We initilise the random index in array for pick a word
	array = Clean(array)                   //We check if the array is correctly formated with words
	word := StringToArray(array[randomNb]) //We initilise the word to find
	lenghtWord := len(word)                //We inisilise the lenhgt of the word
	empty := InitToadd(word)
	for i := 'a'; i != '{'; i++ { //We make the alphabet in alphabet array
		alphabet = append(alphabet, string(i))
	}
	empty, truevValue, remain = AddRandomLtrs(empty, word) //We add random letters to the empty array and extrate the number of ShowValue and the letters Used
	for i := 0; i != len(remain); i++ {                    //We delete the used value in the alphabet
		alphabet = SupLetter(remain[i], alphabet)
	}
	fmt.Println(Base(array, word, empty, lenghtWord, truevValue, mistakes, usedValues, alphabet, nameDB, displayType)) //We call the base function
}
