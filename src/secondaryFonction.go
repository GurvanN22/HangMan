package src

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func ArrayInit(DB string) []string { //This function take nothing and return an array form word.txt
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	file, err := os.Open(DB)
	if err != nil { //We check if the data base is correct
		fmt.Println(colorRed, "ERROR no data base named : ", DB, colorReset)
		os.Exit(1) //We stop the programme
	}
	var array []string //We create the array who will contain the entire words data base

	scanner := bufio.NewScanner(file) //We take make a array with all the lines of the DB

	for scanner.Scan() { //We explore the scanner
		array = append(array, scanner.Text()) //We append the DB line by line in the array variable
	}
	return array
}

func IsIn(array []string, letter string) int { //We check if the letter is in the word and return the number of occurance the of letter in array
	result := 0
	for _, value := range array {
		if value == letter {
			result++
		}
	}
	return result
}

func Random(array []string) int {
	if len(array) == 1 { //				If the len of the array is equal to 1 then we know there is only one thing in it
		return 0
	}
	rand.Seed(time.Now().UnixNano()) //	to get different result every time
	return rand.Intn(len(array))     // 		return a random int to pick a random word in the database
}

func StringToArray(word string) []string { //We passed from a string to an array of string
	var result []string
	for _, value := range word {
		result = append(result, string(value))
	}
	return result
}

func SupLetter(letter string, alphabet []string) (result []string) { //We delete a letter from a alphabet
	result = alphabet
	for index, value := range result {
		if value == letter {
			transposition := result[index+1:]
			result = result[:index]
			index--
			for _, v := range transposition {
				result = append(result, v)
			}
		}
	}
	return
}

func Switch(array []string, empty []string, letter string) []string { //We add the letter to the empty
	for index, value := range array {
		if value == letter {
			empty[index] = letter
		}
	}
	return empty
}

func Clean(array []string) []string {
	/*function to turn uppercase characters into lowercase, if present, from the file words.txt
	and to stop the program if one character is not a valid letter the program can process*/
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	var result []string = array
	var invalids []string
	var sb string
	var hasError bool
	invalids = append(invalids, "\nInvalid character(s) in input file:\n")
	for i, s := range array {
		for _, r := range s {
			if unicode.IsUpper(r) {
				r = unicode.ToLower(r)
				sb = sb + string(r)
				result[i] = sb
			}
			if !unicode.IsLetter(r) {
				fmt.Print(colorRed)
				invalids = append(invalids, fmt.Sprintf("'%s' at word '%s', at line %v\n", string(r), s, i+1))
				hasError = true
			}
		}
	}
	if hasError {
		fmt.Println(invalids, colorReset)
		os.Exit(1)
	}
	return result
}

func ArrayStrToStr(array []string) string { //We tranform a array into a string
	str := ""
	for _, value := range array {
		str += value
	}
	return str
}
