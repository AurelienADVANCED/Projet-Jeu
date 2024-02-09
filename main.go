package main

import (
	items "GoProjet/Items"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	Probabilite = 0.03 // Probabilité d'apparition d'un objet sur une case
)

var (
	carte              Carte
	mage               Magicien
	chevalier          Chevalier
	nain               Nain
	gobelin            Gobelin
	orc                Orc
	elf                Elfe
	PersonnageSelected *Personnage
)

func clearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func choisirClasse() {
	var choix int
	clearConsole()
	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Magicien")
	fmt.Println("2. Chevalier")
	fmt.Println("3. Elfe")
	fmt.Println("4. Nain")

	fmt.Print("Entrez votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		PersonnageSelected = &mage.Personnage
	case 2:
		PersonnageSelected = &chevalier.Personnage
	case 3:
		PersonnageSelected = &elf.Personnage
	case 4:
		PersonnageSelected = &nain.Personnage
	// Ajoutez d'autres cas pour d'autres classes
	default:
		fmt.Println("Choix non valide. Sélectionnez une classe valide.")
		choisirClasse() // Réessayer
	}
}

// Fonction pour générer des réponses aléatoires prédéfinies
func genererReponseAleatoire() string {
	reponses := []string{
		"Bonjour !",
		"Que puis-je faire pour vous ?",
		"Je suis en train de me reposer.",
		"Laissez-moi tranquille.",
	}
	indice := rand.Intn(len(reponses))
	return reponses[indice]
}

func viderCacheClavier() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

var OtherPersonnage []*Personnage

// Vérifier que tous les ennemies sont mort
func VerifAllEnemy(carte Carte) bool {

	// Vérifier si tous les autres personnages sont morts
	for _, character := range OtherPersonnage {
		if character.Vie > 0 {
			return false // Il y a au moins un autre personnage en vie
		}
	}

	return true // Tous les ennemis et autres personnages sont morts, à l'exception du personnage sélectionné
}

func displayMenu() {
	fmt.Println("Choisissez une option :")
	fmt.Println("1. Action")
	fmt.Println("2. Changer la position du personnage")
	fmt.Println("3. Afficher les statistiques des personnages")
	fmt.Println("4. Afficher la carte")
	fmt.Println("5. Informations carte")
	fmt.Println("6. Mes statistiques")
	fmt.Println("7. Inventaire")
	fmt.Println("8. Quitter l'application")
	fmt.Print("Entrez votre choix (1-6) : ")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	InitilisationAPI()

	inventaire = Inventaire{
		Items:      make([]items.Item, 0),
		Taille:     0,
		Taille_max: 10,
	}

	AddPotionDeSoin(1)
	InitialisationDesItems()

	carte = nouvelleCarte(15, 15)

	mage = CreerMagicien()
	chevalier = CreerChevalier()
	nain = nouveauNain()
	gobelin = nouveauGobelin()
	orc = nouveauOrc()
	elf = nouveauElfe()

	carte.placerPersonnage(&mage.Personnage, 5, 5)
	carte.placerPersonnage(&chevalier.Personnage, 10, 10)
	carte.placerPersonnage(&nain.Personnage, 0, 0)
	carte.placerPersonnage(&elf.Personnage, 2, 2)
	carte.placerMonstre(&orc.Monstre, 3, 3)
	carte.placerMonstre(&gobelin.Monstre, 4, 4)

	choisirClasse()

	OtherPersonnage = carte.obtenirAutresPersonnages(PersonnageSelected)
	PlacerObjetsSurCarte(&carte)

	clearConsole()

	fmt.Println("Vous avez choisi : " + PersonnageSelected.Classe)

	carte.afficher()

	var choix int

	for {
		displayMenu()
		fmt.Scan(&choix)

		switch choix {
		case 1:
			var actionChoice int
			clearConsole()
			fmt.Println("2. Communiquer")
			fmt.Print("Entrez votre choix (1-3) : ")
			fmt.Scan(&actionChoice)

			switch actionChoice {
			case 2:
				clearConsole()
				viderCacheClavier()
				fmt.Print("Entrez votre message : ")
				var message string
				fmt.Scanln(&message)
				fmt.Println()

				// Faites en sorte que les autres personnages répondent avec des réponses aléatoires
				autresPersonnages := carte.obtenirAutresPersonnages(PersonnageSelected)
				for _, personnage := range autresPersonnages {
					fmt.Printf("%s répond : \"%s\"\n", personnage.Classe, genererReponseAleatoire())
				}
				fmt.Println()
			case 4:
				clearConsole()
				carte.afficher()
			default:
				fmt.Println("Choix non valide. Veuillez entrer un numéro entre 1 et 3.")
			}

		case 2:
			clearConsole()
			carte.afficher()
			var x, y int
			fmt.Println("Entrez la position de destination (x, y): ")
			fmt.Scan(&x, &y)
			var message string = carte.deplacerSiPossible(PersonnageSelected, x, y)
			fmt.Println(message)
			carte.deplacerAutresPersonnages(PersonnageSelected)
			carte.afficher()

			// Verifie si la partie est terminé
			Finish := VerifAllEnemy(carte)
			if Finish {
				fmt.Print("La partie est terminé")
			}

		case 3:
			afficherStatistiques()
		case 4:
			clearConsole()
			carte.afficher()
		case 5:
			clearConsole()
			carte.afficher()
			var x, y int
			fmt.Println("Entrez la position de destination (x, y): ")
			fmt.Scan(&x, &y)
			fmt.Println(carte.obtenirContenuCase(x, y))
		case 6:
			clearConsole()
			MesStatistiques()

		case 7:
			clearConsole()
			UtiliserInventaire()
		case 8:
			return
		default:
			fmt.Println("Choix non valide. Veuillez entrer un numéro entre 1 et 6.")
		}
	}
}
