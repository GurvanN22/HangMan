package Hangman

import (
	"fmt"
	"io/ioutil"
)

//This function ask to the player a new word to add in the DB
func AddNewWord(db []string, NameDB string) {
	fmt.Println()
	fmt.Print("You can add a new word in the Data Base : ")
	newWord := ""
	fmt.Scan(&newWord)        //We ask the word
	for i := 0; i == 1; i-- { //We check if the word is already in the DB
		if IsIn(db, newWord) == 0 {
			break
		} else {
			fmt.Println("The word is already in the data base please choose another word : ")
			fmt.Scan(&newWord)
		}
	}
	db = append(db, newWord)
	db = Clean(db) //We check if the word is correct
	dbString := ""
	for _, val := range db {
		dbString += val + "\n"
	}
	bytes := []byte(dbString)
	ioutil.WriteFile(NameDB, bytes, 0664) //We add the word in the DB
}
