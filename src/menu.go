package Hangman

import (
	"fmt"
	"os"
	"os/exec"
)

func Menu() (string, string) {

	c := exec.Command("clear") //A revoir https://github.com/Bhupesh-V/til/blob/master/Go/clear-terminal-screen-in-go.md
	c.Stdout = os.Stdout       //This part is for "clean" the cmd terminale
	c.Run()
	//We initialise colors and variables
	difficultyLevel := ""
	colorCyan := "\033[36m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorPurple := "\033[35m"
	scanResult := ""
	displayType := ""
	//----------------------
	file, _ := os.ReadFile("DB/WelcomeDB.txt") //We open the help file in DB
	println(string(file), "\n\n")
	for i := 0; i < 10; { //We made a infinit loop who ask for 1 or 2
		fmt.Print(colorGreen, "Start new Game[1]  ", colorRed, "Exit the game[2]", colorCyan, " : ")
		fmt.Scan(&scanResult)
		if scanResult == "help" || scanResult == "-h" { //We check if the flags for help are used
			DisplayHelp()
		} else if scanResult == "exit" || scanResult == "-e" || scanResult == "leave" { //We check if the flags for exit are used
			fmt.Println("End of the game")
			os.Exit(0)
		} else if scanResult == "1" { //If the value is 1 we break the infinit loop and don't do anything
			break
		} else if scanResult == "2" { //If value is 2 we stop the Programme
			fmt.Println(colorRed, "Goodby, End of the Game")
			os.Exit(0)
		} else {
			fmt.Println(colorRed, "You don't enter 1 or 2")
		}

	}
	fmt.Println(colorCyan, " \n Choose your display type : \n \n ")
	for i := 0; i < 10; { //We made a infinit loop we ask 1, 2, 3 or 4
		fmt.Print(colorGreen, "Normal[1]  ", colorYellow, "Standard[2]  ", colorRed, "Shadow[3]  ", colorPurple, "Thinkertoy[4]  ", colorCyan, ":  ")
		fmt.Scan(&difficultyLevel)
		if difficultyLevel == "help" || difficultyLevel == "-h" { //We check if the flags for help are used
			DisplayHelp()
		} else if difficultyLevel == "exit" || difficultyLevel == "-e" || difficultyLevel == "leave" { //We check if the flags for exit are used
			fmt.Println("End of the game")
			os.Exit(0)
		} else if difficultyLevel == "4" {
			displayType = "DB/AsciiArtThinkertoy.txt"
			break
		}
		if difficultyLevel == "1" {
			fmt.Println(colorGreen)
			displayType = "Normal"
			break
		} else if difficultyLevel == "2" {
			fmt.Println(colorYellow)
			displayType = "DB/AsciiArtStandard.txt"
			break
		} else if difficultyLevel == "3" {
			fmt.Println(colorRed)
			displayType = "DB/AsciiArtShadow.txt"
			break
		} else {
			fmt.Println(colorRed, "You don't enter 1,2 or 3")
		}
	}

	fmt.Println(colorCyan, " \n Choose your difficulty level or load your save : \n \n ")
	for i := 0; i < 10; { //We made a infinit loop we ask 1, 2, 3 or 4
		fmt.Print(colorGreen, "Easy[1]  ", colorYellow, "Medium[2]  ", colorRed, "Hard[3]  ", colorPurple, "Load the save[4]  ", colorCyan, ":  ")
		fmt.Scan(&difficultyLevel)
		if difficultyLevel == "help" || difficultyLevel == "-h" { //We check if the flags for help are used
			DisplayHelp()
		} else if difficultyLevel == "exit" || difficultyLevel == "-e" || difficultyLevel == "leave" { //We check if the flags for exit are used
			fmt.Println("End of the game")
			os.Exit(0)
		} else if difficultyLevel == "4" {
			return "saveOn", displayType
		}
		if difficultyLevel == "1" {
			fmt.Println(colorGreen)
			return "DB/wordsEasy.txt", displayType
		} else if difficultyLevel == "2" {
			fmt.Println(colorYellow)
			return "DB/wordsMedium.txt", displayType
		} else if difficultyLevel == "3" {
			fmt.Println(colorRed)
			return "DB/wordsHard.txt", displayType
		} else {
			fmt.Println(colorRed, "You don't enter 1,2 or 3")
		}
	}
	return "DB/wordsEasy.txt", displayType //This value is never used but the function don't work is this line missing
}
