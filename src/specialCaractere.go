package src

func specialCaractere(letter string) string {
	if letter == "é" || letter == "è" || letter == "ê" {
		return "e"
	} else if letter == "ç" {
		return "c"
	} else if letter == "à" || letter == "â" {
		return "a"
	} else if letter == "ù" {
		return "u"
	} else {
		return letter
	}

}
