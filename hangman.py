import random
import sys

def main():
    r = '\033[91m'             
    g = '\033[92m'          
    n = '\033[0m'
    b = '\033[96m'
    v = '\033[95m'
    mot = MotRandom()
    lettreEntrer = ""
    lettreDevine = [*range(1, len(mot))]
    for i in range(len(lettreDevine)):
        lettreDevine[i] = False
    tentatives = 10
    erreur = 0
    lettre_reveal = (len(mot) // 2) - 1
    for _ in range(lettre_reveal):
        random_index = random.randint(0, (len(mot) - 1))
        if not lettreDevine[random_index]:
            lettreDevine[random_index] = True
    while tentatives > 0:
        print(b +dessin(erreur) + n)
        print(motMasquee(mot, lettreDevine), "\n")
        print(v +"Tentatives restantes : "+ n  ,tentatives )
        print(v +"Lettres déja utilisées : "+ n , lettreEntrer )
        print(v +"Entrer une lettre: " + n)
        lettre = input()

        if lettre == "" or len(lettre) != 1:
             print("\x1bc")
             print("\x1b[2J")
             print("Veuillez entrer une seule lettre valide.")
             continue
        
        if lettre == type(int):
            print("\x1bc")
            print("\x1b[2J")
            print("Veuillez ne mettre que des lettre et rien d'autre")
            continue

        if IsAlpha(lettre) == False:
            print("\x1bc")
            print("\x1b[2J")
            print("Veuillez mettre une lettre en miniscule")
            continue

        if lettre in lettreEntrer:
            print("\x1bc")
            print("\x1b[2J")
            print("Veuillez entrer une lettre non utilisée.") 



        lettreEntrer += lettre 
        lettre_fausse = True

        for i in range(len(mot) - 1):
            if lettre[0] == mot[i]:
                lettreDevine[i] = True
                lettre_fausse = False

        if lettre_fausse:
            tentatives -= 1
            erreur += 1

        if motcomplet(motMasquee(mot, lettreDevine)):
            print("\x1bc")
            print("\x1b[2J")
            print(g + "Bien joué vous avez trouvé le mot " + mot + n)
            exit()    

        print("\x1bc")
        print("\x1b[2J")

    if tentatives == 0:
        print(r +"Vous avez perdu -_-" + n)



def MotRandom():
    with open("words.txt", "r") as f:
        contenu = f.readlines()
        length = len(contenu)
        num = random.randrange(1, length)
        mot = contenu[num]
    return mot 
    

def motMasquee(mot , lettreDevine):
    motmasque = ""
    motcomplet = False
    for i in range(len(mot) - 1):
        if lettreDevine[i]:
            motmasque += mot[i]
        else:
            motmasque += "_"
    return motmasque 


def motcomplet(mot1):
    if "_" in mot1:
        motcomplet = False
    else:
        motcomplet = True
    return motcomplet

def IsAlpha(lettre):
    validation = True
    lettre_alpha = "abcdefghijklmnopqrstuvwxyz"
    if lettre in lettre_alpha:
        validation = True
    else:
        validation = False
    return validation

def dessin(erreur):
    with open("hangman.txt", "r") as f:
        contenu = f.readlines()
    cpt = 0
    cptEnd = 0
    if erreur == 0:
        cpt = 0
        cptEnd = 0
    elif erreur == 1:
        cpt = 0
        cptEnd = 8
    elif erreur == 2:
        cpt = 9
        cptEnd = 16
    elif erreur == 3:
        cpt = 16
        cptEnd = 24
    elif erreur == 4:
        cpt = 24
        cptEnd = 32
    elif erreur == 5:
        cpt = 32
        cptEnd = 40
    elif erreur == 6:
        cpt = 40
        cptEnd = 48
    elif erreur == 7:
        cpt = 48
        cptEnd = 56
    elif erreur == 8:
        cpt = 56
        cptEnd = 64
    elif erreur == 9:
        cpt = 64
        cptEnd = 72
    elif erreur == 10:
        cpt = 72
        cptEnd = 80
    dessin = ""
    for i in range(cpt , cptEnd):
        dessin += contenu[i]
    return dessin
        
        

print(main())
