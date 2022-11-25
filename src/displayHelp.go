package src

import (
	"fmt"
	"os"
)

func DisplayHelp() {
	file, _ := os.ReadFile("DB/commandHelp.txt") //We open the help file in DB
	fmt.Print("\n", "\n", "\n", "\n", string(file), "\n", "\n", "\n")

}
