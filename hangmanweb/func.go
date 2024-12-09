package hangmanweb

import "fmt"

func CheckWin() {
	for _, i := range Data.TabHidden {
		win = true
		if i == "_" { // If any element is "_", the game is not won
			win = false
			break
		}
	}
	if win {
		fmt.Println("GG, vous avez gagné !")
	}
}

func Win() {
	
}
