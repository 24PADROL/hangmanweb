package hangmanweb

import "fmt"

func CheckWin() {
	for _, i := range Data.TabHidden {
		if i != "-" {
			win = true
		}
	}
	fmt.Println("gg")
}
//test
