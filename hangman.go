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
	/**On va d'abord crée une fonction Main pour réunir les autres fonctions qu'on va crée en dessous pour réaliser le jeu du pendu **/
	mot := MotRandom()
	lettreEntrer := ""
	lettresDevinees := make([]bool, len(mot))
	tentatives := 10
	erreur := 0
	//Cette boucle permet de continuer le jeu tant qu'il lui reste des tentatives
	for tentatives > 0 {
		fmt.Println(mot)
		dessin(erreur)
		fmt.Println(motMasque(mot, lettresDevinees))
		fmt.Printf("Tentatives restantes : %d\n", tentatives)
		fmt.Println("Lettres déja utilisées : " + lettreEntrer)
		fmt.Print("Entrer une lettre: ")
		var lettre string
		fmt.Scanln(&lettre)

		//Cette condition permet de vérifier si le joueur entre bien une lettre en miniscule et non un chiffre
		if lettre == "" || len(lettre) != 1 || !IsAlpha(lettre) {
			fmt.Println("\nVeuillez entrer une seule lettre valide.")
			continue
		}
		//Cette condition permet de vérifier  si le joueur entre bien une lettre déjà pas utiliser
		if strings.Contains(lettreEntrer, lettre) {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Print("Veuillez entrer une lettre non utilisée.")
		}

		lettreEntrer += lettre

		lettre_fausse := true

		//Boucle permettant de changer les valeurs dans le slice lettreDeveinees
		for i, c := range mot {
			if lettre[0] == byte(c) {
				lettresDevinees[i] = true
				lettre_fausse = false
			}
		}
		//Teste si lettre_fausse == True et donc diminue de 1 les tentatives et augmente de 1 les erreurs
		if lettre_fausse {
			tentatives--
			erreur++
		}

		//Cette condition permet de vérifier si le mot à été trouver et si c'est le cas arréte le jeu
		if motMasque(mot, lettresDevinees) == mot {
			fmt.Printf("\x1bc")
			fmt.Printf("\x1b[2J")
			fmt.Printf("Bien joué vous avez trouvé le mot " + mot)
			break
		}
		fmt.Printf("\x1bc")
		fmt.Printf("\x1b[2J")

	}
	//Condition permettant l'arrét de jeu si le joueur n'a plus de tentative
	if tentatives == 0 {
		fmt.Printf("Vous avez perdu -_-")
	}

}

// Fonction qui va prendre un mot random du Fichier "words.txt" et le return
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

// Fonction qui créée l'interface du mot qui est masqué
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

// Fonction permettant l'affichage du pendu
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

// Fonction qui teste si le string passé en paramètre est bien une lettre et nonautre chose
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
