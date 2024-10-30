package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//Shows the message of commands entered, launches or closes the game depending on the command entered
func Menu() {
	if len(os.Args)<2{
		nameFill = "data/motsimple.txt"
	}else if os.Args[1] == "-h" {
		fmt.Println("les arguments sont -s pour les mots simples\n et -d pour les mots durs les mots sont par default")
		return
	} else if os.Args[1] == "-s" {
		nameFill = "data/motsimple.txt"
	} else if os.Args[1] == "-d" {
		nameFill = "data/motdur.txt"
	} else {
		fmt.Println("arg invalide")
		return
	}
	WelcomePlayer()
	fmt.Println("Bienvenue")
	for {
		fmt.Println("\nPour jouer, écrire [play]\nPour quitter, écrire [quit]")
		fmt.Scan(&play)
		if play == "play" {
			for {
				fmt.Println("\nChoisissez la difficulté [s/d]")
				fmt.Scan(&diff)
				if diff != "s" && diff != "d" {
					fmt.Println("difficulté non valide")
					continue
				}
				break
			}
			Game()
			continue
		}
		if play == "quit" {
			break
		}
		fmt.Println("commande invalide")
	}
}

//Run the entire flow of the game
func Game() {
	hangman = nil
	printHangman()
	randomWord()
	//defined the elements of the game
	word = strings.ToLower(words[rand.Intn(200)])
	wordfind = ""
	wordcomp = ""
	tab = nil
	pv = 10
	letterTestedFalse = nil
	letterTestedTrue = nil
	wordletter := []rune(word)
	WordSelect()
}

func WordSelect() {
	for i := 0; i < len([]rune(word)); i++ {
		tab = append(tab, "_")
		tab = append(tab, " ")
		wordfind += "_ "
		if wordletter[i] != 'é' && wordletter[i] != 'è' {
			wordcomp += string(wordletter[i])
		}
		if wordletter[i] == 'é' || wordletter[i] == 'è' {
			wordcomp += "e"
		}
	}
}
	if diff == "s" {
		fmt.Println(wordfind)
	}
	for {
		stupid = false
		win = true
		lose = false
		fmt.Scan(&input)
		if input == "é" || input == "è" {
			input = "e"
		}
		if input == "à" {
			input = "a"
		}
		for _, i := range letterTestedFalse {
			if i == input {
				fmt.Println("t'as déjà testé la lettre(espèce de débile)")
				stupid = true
			}
		}
		for _, i := range letterTestedTrue {
			if i == input {
				fmt.Println("t'as déjà testé la lettre(espèce de débile)")
				stupid = true
			}
		}
		if stupid {
			continue
		}
		if len(input) != 1 && input != word {
			if input == "quit" {
				break
			}
			pv -= 2
			PrintNext(pv, input)
			PrintTab()
		}
		if len(input) > 1 && input == word {
			PrintWin(word)
			break
		}
		if !(input >= "A" && input <= "Z" || input >= "a" && input <= "z") {
			fmt.Println("La lettre n'est pas valide")
			continue
		}
		for _, i := range wordcomp {
			if string(i) == input {
				tab[count] = input
				ishere = true

			}
			count += 2
		}
		if ishere {
			letterTestedTrue = append(letterTestedTrue, input)
		} else {
			letterTestedFalse = append(letterTestedFalse, input)
		}
		count = 0
		if !ishere && len(input) == 1 {
			pv--
			PrintNext(pv, input)
			PrintTab()
		}
		//If there are no more attempts left, the defeat message is displayed.
		if pv <= 0 {
			PrintLose(word)
			lose = true
		}
		ishere = false
		wordfind = ""
		for i := 0; i < len(tab); i++ {
			wordfind += tab[i]
		}
		if pv != 0 {
			fmt.Print("lettre fausse")
			fmt.Println(letterTestedFalse)
			if diff == "s" {
				fmt.Println(wordfind)
			} else {
				fmt.Print("lettre bonne")
				fmt.Println(letterTestedTrue)
			}
		}
		for i := 0; i < len(tab); i++ {
			if tab[i] == "_" {
				win = false
			}
		}
		//Displays the Hangman according to the number of tries we have
		if pv != 10 {
			fmt.Printf(hangman[9-pv])
		}
		if win {
			PrintWin(word)
			break
		}
		if lose {
			break
		}
	}
}
