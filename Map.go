package main

import (
	items "GoProjet/Items"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func AffichageChoixInventaire() {

	fmt.Println("1. Afficher l'inventaire")
	fmt.Println("2. Utiliser un Item")
	fmt.Println("3. Quitter l'inventaire")
	fmt.Print("Entrez votre choix (1-3) : \n\n")
}

func UtiliserInventaire() {
	for {
		clearConsole()
		var choix int
		AffichageChoixInventaire()
		fmt.Scan(&choix)
		switch choix {
		case 1:
			clearConsole()
			AfficherInventaire()
		case 2:
			fmt.Println("Entrez le nom de l'item : ")
		case 3:
			clearConsole()
			carte.afficher()
			return
		default:
			fmt.Println("Choix non valide. Veuillez entrer un numéro entre 1 et 3.")
		}
	}
}

func (c *Carte) obtenirContenuCase(x, y int) (contenu string) {
	if x < 0 || x >= c.Largeur || y < 0 || y >= c.Hauteur {
		return "Hors de la carte"
	}

	caseCourante := c.Grille[y][x]

	if caseCourante.Personnage == PersonnageSelected {
		contenu = "Vous êtes ici"
	}
	if caseCourante.Personnage != nil && caseCourante.Personnage != PersonnageSelected {
		contenu = caseCourante.Personnage.Classe
	} else if caseCourante.Monstre != nil {
		contenu = "Monstre: " + caseCourante.Monstre.Classe
	} else if caseCourante.Contenu != nil {
		contenu = "Objet: " + caseCourante.Contenu.GetNom()
	} else {
		contenu = "Case vide"
	}

	return
}

// Méthode pour afficher la carte
func (c Carte) afficher() {
	fmt.Print("   ") // Espacement pour les numéros de ligne
	for x := 0; x < c.Largeur; x++ {
		fmt.Printf("%2d ", x)
	}
	fmt.Println()
	for y, ligne := range c.Grille {
		fmt.Printf("%2d ", y)
		var ligneAffichage []string
		for _, caseCourante := range ligne {
			symbole := " . " // Symbole pour une case vide
			if caseCourante.Personnage != nil {
				if PersonnageSelected != nil && caseCourante.Personnage == PersonnageSelected {
					symbole = "\033[31m " + string(caseCourante.Personnage.Classe[0]) + " \033[0m" // Personnage sélectionné en rouge
				} else {
					symbole = "\033[32m " + string(caseCourante.Personnage.Classe[0]) + " \033[0m" // Autres personnages
				}
			} else if caseCourante.Monstre != nil {
				symbole = "\033[34m " + string(caseCourante.Monstre.Classe[0]) + " \033[0m" // Symbole pour un monstre
			} else if caseCourante.Contenu != nil {
				if caseCourante.Contenu.TypeItem() == 2 {
					symbole = "\033[33m " + string([]rune(caseCourante.Contenu.GetNom())[0]) + " \033[0m" // Symbole pour une arme
				} else if caseCourante.Contenu.TypeItem() == 1 {
					symbole = "\033[33m A \033[0m" // Symbole pour un autre item
				}

			}
			ligneAffichage = append(ligneAffichage, symbole)
		}

		fmt.Print(strings.Join(ligneAffichage, ""))
		if y < len(OtherPersonnage) && OtherPersonnage[y] != nil {
			afficherStatutPersonnage(*OtherPersonnage[y], false)
		} else {
			fmt.Println()
		}
	}
	fmt.Println()
}

func nouvelleCarte(largeur, hauteur int) Carte {
	grille := make([][]Case, hauteur)
	for i := range grille {
		grille[i] = make([]Case, largeur)
	}
	return Carte{Grille: grille, Largeur: largeur, Hauteur: hauteur}
}

func PlacerObjetsSurCarte(carte *Carte) {
	var classArmes []items.Arme // Liste pour stocker les armes de la classe du personnage

	// Filtrer les armes par classe
	for _, arme := range toutesLesArmes {
		if arme.Classe == PersonnageSelected.Classe && arme.Nom != PersonnageSelected.Arme.Nom {
			classArmes = append(classArmes, arme)
		}
	}

	// Assure-toi que la graine aléatoire est initialisée correctement
	rand.Seed(time.Now().UnixNano())

	// Placer potions de soin et viande
	for x := 0; x < carte.Largeur; x++ {
		for y := 0; y < carte.Hauteur; y++ {
			if rand.Float64() < Probabilite {
				if rand.Intn(2) == 0 {
					carte.Grille[x][y].Contenu = AddPotionDeSoin(1)
				} else {
					carte.Grille[x][y].Contenu = AddViande(1)
				}
			}
		}
	}

	// Vérifier si la liste d'armes de classe n'est pas vide
	if len(classArmes) > 0 {
		// Sélectionner une arme aléatoire
		armeAleatoire := classArmes[rand.Intn(len(classArmes))]

		// Sélectionner une position aléatoire sur la carte
		xAleatoire := rand.Intn(carte.Largeur)
		yAleatoire := rand.Intn(carte.Hauteur)

		// Placer l'arme à la position aléatoire
		carte.Grille[xAleatoire][yAleatoire].Contenu = &armeAleatoire
	}
}

func (c *Carte) placerMonstre(m *Monstre, x, y int) {
	if x < 0 || x >= c.Largeur || y < 0 || y >= c.Hauteur {
		fmt.Println("Position hors de la carte pour le monstre.")
		return
	}
	c.Grille[y][x].Monstre = m
}
