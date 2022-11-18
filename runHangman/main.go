package main

import (
	"Hangman"
)

func main() {
	nameDB, displayType := Hangman.Menu()
	if nameDB != "saveOn" { //We check if a save is load or not
		Hangman.InitialiatonOfStartsValues(nameDB, displayType)
	} else {
		Hangman.ExtractSave()
	}
}
