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
	reader := bufio.NewReader(os.Stdin)
	mot := MotRandom()
	lettreEntrer := ""
	lettresDevinees := make([]bool, len(mot))
	tentatives := 10
	for tentatives > 0 {
		fmt.Println(motMasque(mot, lettresDevinees))
		fmt.Printf("Tentatives restantes : %d\n", tentatives)
		fmt.Println("Lettres déja utilisées : " + lettreEntrer)
		fmt.Print("Entrer une lettre: ")
		lettre, _ := reader.ReadString('\n')

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
		}

		if motMasque(mot, lettresDevinees) == mot {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Printf("Bien joué vous avez trouvé le mot" + mot)
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
