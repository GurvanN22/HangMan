package src

import (
	"fmt"
	"os"
	"os/exec"
)

func Display(db []string, array []string, value int, NameDB string, displayType string) {
	c := exec.Command("clear") //We use this to clean the terminal
	c.Stdout = os.Stdout
	c.Run()
	if value == 1 { //If the player win
		file, _ := os.ReadFile("DB/EndDB.txt") //We open the help file in DB
		println(string(file), "\n\n")
		fmt.Print("The word was :  ")
		if displayType == "Normal" {
			for _, letter := range array { //We print the answer
				fmt.Print(letter)
			}
		} else {
			fmt.Println("\n")
			DisplayAsciiArt(ArrayStrToStr(array), displayType)
		}
		AddNewWord(db, NameDB) //We call the addNewWord function
		fmt.Print("The word was add at the data base, thank you !")
	} else {
		file, _ := os.ReadFile("DB/LoseDB.txt") //We open the help file in DB
		println(string(file), "\n\n")
		fmt.Print("The word was :  ")
		if displayType == "Normal" {
			for _, letter := range array { //We print the answer
				fmt.Print(letter)
			}
		} else {
			fmt.Println("\n")
			DisplayAsciiArt(ArrayStrToStr(array), displayType)
		}
	}
}
