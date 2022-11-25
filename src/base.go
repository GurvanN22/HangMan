package src

//This function is the core of the whole Hangman who link all functions and transmit informations , the Hangman end want the end condition is true
func Base(db []string, array []string, toAdd []string, lenght int, trueValue int, falseValue int, usedValues []string, remain []string, NameDB string, displayType string) string {
	DisplayHangMan(remain, toAdd, usedValues, falseValue, displayType)
	if lenght == trueValue || falseValue == 10 { //,We made the end condition , if trueValue = lenght the player win, if falseValue = 10 the player loose
		if falseValue == 10 { //We check if the player win or loose
			Display(db, array, 0, NameDB, displayType) //If the index 2 argument is 0 it say to the Display function the player loose
		} else {
			Display(db, array, 1, NameDB, displayType) //Else is the index 2 argument is 1 it say the player win
		}
		return ""
	}
	entry := ValueTest(array, toAdd, lenght, trueValue, falseValue, usedValues, remain, NameDB, displayType) //We use ValueTest to ask a letter in the terminale
	if IsIn(array, entry) != 0 && IsIn(remain, entry) != 0 {                                                 //If letter is in the word and is letter is not already use
		toAdd = Switch(array, toAdd, entry)                                                                   //We add the find letter in the display word (with the "_")
		trueValue += IsIn(array, entry)                                                                       //We add the to true value the occurance of the letter in the word
		remain = SupLetter(entry, remain)                                                                     //We delete the letter give by the player in the remain array of unuse letters
		return Base(db, array, toAdd, lenght, trueValue, falseValue, usedValues, remain, NameDB, displayType) //We make a new loop with the new informations
	} else {
		falseValue++                                                                                          //We add one to error
		usedValues = append(usedValues, entry)                                                                //We add the letter to usedValue who contain the False answers of the player
		remain = SupLetter(entry, remain)                                                                     //We delete the letter give by the player in the remain array of unuse letters
		return Base(db, array, toAdd, lenght, trueValue, falseValue, usedValues, remain, NameDB, displayType) //We make a new loop with the new informations
	}
}
