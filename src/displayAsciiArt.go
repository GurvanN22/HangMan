package src

import "fmt"

func DisplayAsciiArt(word string, types string) { //This function Display Ascii art
	arrayDB := ArrayInit(types)
	var displayArray [8]string
	for _, value := range word {
		if value == 95 { //We check if the caracter is a "_"
			for i := 0; i != 8; i++ {
				displayArray[i] = displayArray[i] + arrayDB[568+i]
			}
		} else {
			value = (value-97)*9 + 586
			valueInt := int(value)
			for i := 0; i != 8; i++ {
				displayArray[i] = displayArray[i] + arrayDB[valueInt+i]
			}
		}
	}
	for _, value := range displayArray {
		fmt.Println(value)
	}
}
