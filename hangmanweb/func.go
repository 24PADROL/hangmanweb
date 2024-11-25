package hangmanweb

func FindingLetter() {
	for _, i := range Data.Word {
		if string(i) == "LettreARecuperer" {
			ishere = true
		}
	}
}

func RevealLetter() {
	if ishere == true {
		Data.TabHidden[count] = "LettreARecuperer"
	}
}
//test
