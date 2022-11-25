package Hangman

import (
	"fmt"
	"os"
)

func ValueTest(array []string, toAdd []string, lenght int, trueValue int, falseValue int, usedValues []string, remain []string, NameDB string, displayType string) string { //We ask a letter and ckeck if the letter is a correct input, else if the input is incorrect it make a infinit loop while the input is correct
	entry := ""
	for i := 0; i <= 1; i-- { // We make a inifinit loop while entry is not Letter
		fmt.Print("Enter a letter : ")
		fmt.Scan(&entry)
		entry = specialCaractere(entry)
		if entry == "help" || entry == "-h" { //We check if the flags for help are used
			DisplayHelp()
		} else if entry == "exit" || entry == "-e" || entry == "leave" { //We check if the flags for exit are used
			fmt.Println("End of the game")
			os.Exit(0)
		} else if entry == "save" || "-s" == entry { //We check if the flags for save are used
			fmt.Println("Save and End of the game")
			SaveProgration(array, toAdd, lenght, trueValue, falseValue, usedValues, remain, NameDB, displayType)
			os.Exit(0)
		} else if len(entry) == 1 { //We check if the entry letter is a min letter
			if entry >= "A" && entry <= "Z" { //If the entry is a capital Letter we transform the entry in a min Letter
				bytes := []byte(entry)
				entry = string(bytes[0] + 32)
			}
			if entry >= "a" && entry <= "z" {
				if IsIn(remain, entry) == 0 && len(entry) == 1 { // We check if the entry is in the remain array
					fmt.Println("The value was already use")
				} else {
					break
				}
			} else {
				fmt.Println("ERROR : the caracter is not a letter")
			}
		} else if len(entry) > 1 {
			fmt.Println("ERROR : too many values ")
		}

	}
	return entry
}
