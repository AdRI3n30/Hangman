package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	mot := MotRandom()
	lettreEntrer := ""
	lettresDevinees := make([]bool, len(mot))
	tentatives := 10
	erreur := 0
	for tentatives > 0 {
		fmt.Println(mot)
		dessin(erreur)
		fmt.Println(motMasque(mot, lettresDevinees))
		fmt.Printf("Tentatives restantes : %d\n", tentatives)
		fmt.Println("Lettres déja utilisées : " + lettreEntrer)
		fmt.Print("Entrer une lettre: ")
		var lettre string
		fmt.Scanln(&lettre)

		if lettre == "" || len(lettre) != 1 || !IsAlpha(lettre) {
			fmt.Println("\nVeuillez entrer une seule lettre valide.")
			continue
		}
		if strings.Contains(lettreEntrer, lettre) {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Print("Veuillez entrer une lettre non utilisée.")
		}

		lettreEntrer += lettre

		lettre_fausse := true

		for i, c := range mot {
			if lettre[0] == byte(c) {
				lettresDevinees[i] = true
				lettre_fausse = false
			}
		}
		if lettre_fausse {
			tentatives--
			erreur++
		}

		if motMasque(mot, lettresDevinees) == mot {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Printf("Bien joué vous avez trouvé le mot " + mot)
			break
		}
		fmt.Printf("\x1bc")
		fmt.Printf("\x1b[2J")

	}
	if tentatives == 0 {
		fmt.Printf("Vous avez perdu -_-")
	}

}
func MotRandom() string {
	file, err := os.Open("words.txt")
	var words []string
	if err != nil {
		log.Fatal(err)
	}
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
	mot_rand := words[rand.Intn(len(words))]
	return mot_rand
}

func motMasque(mot string, lettresDevinees []bool) string {
	motMasque := ""
	for i, c := range mot {
		if lettresDevinees[i] {
			motMasque += string(c)
		} else {
			motMasque += "_"
		}
	}
	return motMasque
}

func dessin(erreur int) {
	f, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cpt := 0
	cptEnd := 0
	switch erreur {
	case 0:
		cpt = 0
		cptEnd = 0
	case 1:
		cpt = 0
		cptEnd = 8
	case 2:
		cpt = 9
		cptEnd = 16
	case 3:
		cpt = 16
		cptEnd = 24
	case 4:
		cpt = 24
		cptEnd = 32
	case 5:
		cpt = 32
		cptEnd = 40
	case 6:
		cpt = 40
		cptEnd = 48
	case 7:
		cpt = 48
		cptEnd = 56
	case 8:
		cpt = 56
		cptEnd = 64
	case 9:
		cpt = 64
		cptEnd = 72
	case 10:
		cpt = 72
		cptEnd = 80
	}
	i := 0
	for scanner.Scan() {
		if i >= cpt && i < cptEnd {
			fmt.Println(scanner.Text())
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func IsAlpha(s string) bool {
	length := len(s)
	compt := 0
	for index, i := range s {
		if i >= rune(65) && i < rune(91) || i >= rune(97) && i <= rune(122) {
			compt++
			index++
		}
	}
	if compt == length {
		return true
	} else {
		return false
	}
}
