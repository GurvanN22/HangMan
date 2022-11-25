package Hangman

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

//This function take all importants value of the game and save them in a DB line by line
func SaveProgration(array []string, toAdd []string, lenght int, trueValue int, falseValue int, usedValues []string, remain []string, NameDB string, dispalyType string) {
	//We make a array with all informations
	arrayToDB := []string{ArrayStrToStr(array), ArrayStrToStr(toAdd), strconv.Itoa(lenght), strconv.Itoa(trueValue), strconv.Itoa(falseValue), ArrayStrToStr(usedValues), ArrayStrToStr(remain), NameDB, dispalyType}
	dbString := ""
	for _, val := range arrayToDB {
		dbString += val + "\n"
	}
	bytes := []byte(dbString)
	ioutil.WriteFile("DB/DBSave.txt", bytes, 0664) //We save informations in the DB
}

//This function extracte the save from the DB and run the Hangman herself
func ExtractSave() {
	arraySave := ArrayInit("DB/DBSave.txt")
	word := StringToArray(arraySave[0])
	toAdd := StringToArray(arraySave[1])
	lenght, _ := strconv.Atoi(arraySave[2])
	trueValue, _ := strconv.Atoi(arraySave[3])
	falseValue, _ := strconv.Atoi(arraySave[4])
	usedValues := StringToArray(arraySave[5])
	remain := StringToArray(arraySave[6])
	NameDB := arraySave[7]
	displayType := arraySave[8]
	array := ArrayInit(NameDB)
	colorPurple := "\033[35m" // We display the Hangman in purple for say that is a save
	fmt.Print(colorPurple)
	fmt.Println(Base(array, word, toAdd, lenght, trueValue, falseValue, usedValues, remain, NameDB, displayType)) //We call the base function
}
