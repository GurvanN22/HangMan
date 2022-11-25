package src

import (
	"fmt"
	"os"
	"os/exec"
)

func DisplayHangMan(remain []string, empty []string, usedValues []string, value int, displayType string) {
	c := exec.Command("clear") //We use this to clean the terminal
	c.Stdout = os.Stdout
	c.Run()
	colorRed := "\033[31m" //We import red color
	if value >= 7 {
		fmt.Println(colorRed) //We color the Hangman in red if the player make more than 7 mistakes
	}
	if value != 0 { //We will display the Hangman
		array := ArrayInit("DB/DBH.txt")              //We extracte the Hangman from the DB
		valueIndex := (value - 1) * 8                 //A part of the Hangman make 7 line so valueIndex is the index of the start of the new Hangman to display
		for i := valueIndex; i != valueIndex+7; i++ { //We display the next 7 new lines
			fmt.Println(array[i])
		}
	}
	if displayType == "Normal" { //We display informations on the game
		for _, letter := range empty { //We print the answer
			fmt.Print(letter)
		}
	} else {
		fmt.Println("\n")
		DisplayAsciiArt(ArrayStrToStr(empty), displayType)
	}
	fmt.Println("\n")
	fmt.Println("UsedValue :", usedValues)
	fmt.Println("Remainning letters : ", remain)
	fmt.Println("Remain try : ", 10-value)
}
