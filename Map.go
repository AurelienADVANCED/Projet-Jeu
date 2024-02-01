package main

import (
	items "GoProjet/Items"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	Probabilite = 0.03 // Probabilité d'apparition d'un objet sur une case
)

var (
	carte            Carte
	mage             Magicien
	chevalier        Chevalier
	nain             Nain
	gobelin          Gobelin
	orc              Orc
	elf              Elfe
	PersonnageSelect *Personnage
)

type Personnage struct {
	Vie              float64
	VieMax           float64
	Arme             items.Arme
	Mana             float64 // Pour les personnages pouvant utiliser de la magie
	Force            float64 // Pour les personnages physiques
	Agilite          float64 // Pour les personnages physiques
	Armure           float64
	Classe           string
	RayonDeplacement int
}

type Monstre struct {
	Vie              float64
	VieMax           float64
	Arme             items.Arme
	NiveauDeMenace   float64
	Classe           string
	RayonDeplacement int
}

func attaquer(joueur *Personnage, adversaire *Personnage) string {
	adversaire.Vie -= joueur.Arme.Degats
	message := "Degats infligés : " + fmt.Sprintf("%.2f", joueur.Arme.Degats) + "\nVie de l'adversaire : " + fmt.Sprintf("%.2f", adversaire.Vie) + "\n" + joueur.Classe + " : " + fmt.Sprintf("%.2f", joueur.Vie) + "\n"

	if adversaire.Vie <= 0 {
		carte.retirerPersonnage(adversaire)
		message += "L'adversaire a été vaincu et retiré de la carte."
	}

	return message
}

func (c *Carte) retirerPersonnage(p *Personnage) {
	for y := range c.grille {
		for x := range c.grille[y] {
			if c.grille[y][x].Personnage == p {
				c.grille[y][x].Personnage = nil
				return
			}
		}
	}
}

type PersonnageRef struct {
	personnage *Personnage
	nom        string
	estVivant  func() bool
}

type Case struct {
	Personnage *Personnage
	Monstre    *Monstre
	Contenu    items.Item
}

type Carte struct {
	grille           [][]Case
	largeur, hauteur int
}

func nouvelleCarte(largeur, hauteur int) Carte {
	grille := make([][]Case, hauteur)
	for i := range grille {
		grille[i] = make([]Case, largeur)
	}
	return Carte{grille: grille, largeur: largeur, hauteur: hauteur}
}

// Fonction pour placer un personnage sur la carte
func (c *Carte) placerPersonnage(p *Personnage, x, y int) {
	if x < 0 || x >= c.largeur || y < 0 || y >= c.hauteur {
		fmt.Println("Position hors de la carte.")
		return
	}
	c.grille[y][x].Personnage = p
}

// Méthode pour afficher la carte
func (c Carte) afficher() {
	fmt.Print("   ") // Espacement pour les numéros de ligne
	for x := 0; x < c.largeur; x++ {
		fmt.Printf("%2d ", x)
	}
	fmt.Println()
	for y, ligne := range c.grille {
		fmt.Printf("%2d ", y)
		var ligneAffichage []string
		for _, caseCourante := range ligne {
			symbole := " . " // Symbole pour une case vide
			if caseCourante.Personnage != nil {
				if PersonnageSelect != nil && caseCourante.Personnage == PersonnageSelect {
					symbole = "\033[31m " + string(caseCourante.Personnage.Classe[0]) + " \033[0m" // Personnage sélectionné en rouge
				} else {
					symbole = "\033[32m " + string(caseCourante.Personnage.Classe[0]) + " \033[0m" // Autres personnages
				}
			} else if caseCourante.Monstre != nil {
				symbole = "\033[34m " + string(caseCourante.Monstre.Classe[0]) + " \033[0m" // Symbole pour un monstre
			} else if caseCourante.Contenu != nil {
				symbole = "\033[33m " + string(caseCourante.Contenu.GetNom()[0]) + " \033[0m" // Symbole pour un objet
			}
			ligneAffichage = append(ligneAffichage, symbole)
		}
		fmt.Println(strings.Join(ligneAffichage, ""))
	}
	fmt.Println()
}

func (c *Carte) obtenirPersonnage(x, y int) *Personnage {
	if x < 0 || x >= c.largeur || y < 0 || y >= c.hauteur {
		return nil
	}
	return c.grille[y][x].Personnage
}

func (c *Carte) deplacerPersonnage(p *Personnage, xDest, yDest int) {
	if xDest < 0 || xDest >= c.largeur || yDest < 0 || yDest >= c.hauteur {
		fmt.Println("Position de destination invalide.")
		return
	}

	// Trouver la position actuelle du personnage
	var xOrig, yOrig int
	trouve := false
	for y := range c.grille {
		for x := range c.grille[y] {
			if c.grille[y][x].Personnage == p {
				xOrig, yOrig = x, y
				trouve = true
				break
			}
		}
		if trouve {
			break
		}
	}

	// Déplacer le personnage
	c.grille[yDest][xDest].Personnage = p
	c.grille[yOrig][xOrig].Personnage = nil
}

func (c *Carte) peutDeplacer(x, y int) bool {
	return x >= 0 && x < c.largeur && y >= 0 && y < c.hauteur &&
		c.grille[y][x].Personnage == nil
}

func (c *Carte) trouverPositionPersonnage(p *Personnage) (int, int, bool) {
	for y, ligne := range c.grille {
		for x, caseCourante := range ligne {
			if caseCourante.Personnage == p {
				return x, y, true
			}
		}
	}
	return -1, -1, false
}

func (c *Carte) deplacerSiPossible(p *Personnage, xDest, yDest int) string {
	clearConsole()

	xOrig, yOrig, trouve := c.trouverPositionPersonnage(p)
	if !trouve {
		return "Personnage non trouvé sur la carte."
	}

	// Calculez la distance entre la position actuelle et la destination
	distance := math.Sqrt(math.Pow(float64(xDest-xOrig), 2) + math.Pow(float64(yDest-yOrig), 2))
	if distance == 0 {
		return "Déplacement impossible. Vous êtes déjà ici."
	}

	// Vérifiez si la destination est dans le rayon autorisé
	if distance <= float64(p.RayonDeplacement) {
		var message string = carte.obtenirContenuCase(xDest, yDest)
		switch message {
		case "Vous êtes ici":
			return "Déplacement impossible. Vous êtes déjà ici."
		case "Case vide":
			c.deplacerPersonnage(p, xDest, yDest)
			return "Déplacement effectué."
		case "Hors de la carte":
			return "Déplacement impossible. Hors de la carte."
		default:
			personnages := map[string]PersonnageRef{
				"Elfe":      {&elf.Personnage, "Elfe", func() bool { return elf.Vie > 0 }},
				"Magicien":  {&mage.Personnage, "Magicien", func() bool { return mage.Vie > 0 }},
				"Chevalier": {&chevalier.Personnage, "Chevalier", func() bool { return chevalier.Vie > 0 }},
				"Nain":      {&nain.Personnage, "Nain", func() bool { return nain.Vie > 0 }},
			}

			for key, val := range personnages {
				if strings.Contains(message, key) && p.Classe != key {
					if val.estVivant() {
						message := attaquer(val.personnage, p)
						return "Vous avez attaqué un " + val.nom + ".\n" + message
					} else {
						return "Le " + val.nom + " est mort"
					}
				}
			}
			return "Déplacement impossible. " + message
		}
	} else {
		return "Déplacement impossible. Hors de portée."
	}

}

func (c *Carte) obtenirAutresPersonnages(p *Personnage) []*Personnage {
	var autresPersonnages []*Personnage
	for _, ligne := range c.grille {
		for _, caseCourante := range ligne {
			if caseCourante.Personnage != nil && caseCourante.Personnage != p {
				autresPersonnages = append(autresPersonnages, caseCourante.Personnage)
			}
		}
	}
	return autresPersonnages
}

func (c *Carte) deplacerAutresPersonnages(p *Personnage) {
	autresPersonnages := c.obtenirAutresPersonnages(p)
	for _, personnage := range autresPersonnages {
		xOrig, yOrig, _ := c.trouverPositionPersonnage(personnage)
		xDest, yDest := xOrig, yOrig

		// Générer une direction aléatoire
		direction := rand.Intn(4) // 0: haut, 1: bas, 2: gauche, 3: droite
		switch direction {
		case 0:
			yDest--
		case 1:
			yDest++
		case 2:
			xDest--
		case 3:
			xDest++
		}

		if c.peutDeplacer(xDest, yDest) {
			c.grille[yOrig][xOrig].Personnage = nil
			c.grille[yDest][xDest].Personnage = personnage

			xOrig, yOrig, _ := c.trouverPositionPersonnage(p)

			// Calculez la distance entre la position actuelle et la destination
			distance := math.Sqrt(math.Pow(float64(xDest-xOrig), 2) + math.Pow(float64(yDest-yOrig), 2))

			if distance <= float64(personnage.RayonDeplacement) {
				fmt.Print("Le personnage " + personnage.Classe + " vous a vu\n")
				message := attaquer(personnage, p)
				fmt.Print(message)

			}
		}
	}
}

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

func AfficherStats_Mage(perso Personnage) {
	fmt.Printf("	Mage\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", mage.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", mage.VieMax) + "\n")
	fmt.Print("Mana : " + fmt.Sprintf("%.2f", mage.Mana) + "\n")
	fmt.Print("Force : " + fmt.Sprintf("%.2f", mage.Force) + "\n")
	fmt.Print("Agilité : " + fmt.Sprintf("%.2f", mage.Agilite) + "\n")
	fmt.Print("Armure : " + fmt.Sprintf("%.2f", mage.Armure) + "\n")
	fmt.Print("Classe : " + mage.Classe + "\n")
	fmt.Print("Niveau de magie : " + fmt.Sprintf("%d", mage.niveauDeMagie) + "\n\n")
	fmt.Print("Armes : " + perso.Arme.Nom + "\n")
	fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
}

func AfficherStats_Chevalier(perso Personnage) {
	fmt.Printf("	Chevalier\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", chevalier.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", chevalier.VieMax) + "\n")
	fmt.Print("Mana : " + fmt.Sprintf("%.2f", chevalier.Mana) + "\n")
	fmt.Print("Force : " + fmt.Sprintf("%.2f", chevalier.Force) + "\n")
	fmt.Print("Agilité : " + fmt.Sprintf("%.2f", chevalier.Agilite) + "\n")
	//fmt.Print("Armure : " + fmt.Sprintf("%.2f", chevalier.armure) + "\n")
	fmt.Print("Classe : " + chevalier.Classe + "\n\n")
	//fmt.Print("Armure : " + chevalier.armure + "\n")
	fmt.Print("Armes : " + perso.Arme.Nom + "\n")
	fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
}

func AfficherStats_Nain(perso Personnage) {
	fmt.Printf("	Nain\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", nain.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", nain.VieMax) + "\n")
	fmt.Print("Mana : " + fmt.Sprintf("%.2f", nain.Mana) + "\n")
	fmt.Print("Force : " + fmt.Sprintf("%.2f", nain.Force) + "\n")
	fmt.Print("Agilité : " + fmt.Sprintf("%.2f", nain.Agilite) + "\n")
	fmt.Print("Armure : " + fmt.Sprintf("%.2f", nain.Armure) + "\n")
	fmt.Print("Classe : " + nain.Classe + "\n")
	fmt.Print("Expertise en forge : " + nain.expertiseEnForge + "\n")
	fmt.Print("Résistance à l'alcool : " + fmt.Sprintf("%.2f", nain.resistanceAlcool) + "\n\n")
	fmt.Print("Armes : " + perso.Arme.Nom + "\n")
	fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
}

func afficherStatistiques() {
	clearConsole()
	AfficherStats_Mage(mage.Personnage)
	AfficherStats_Chevalier(chevalier.Personnage)
	AfficherStats_Nain(nain.Personnage)
	ElfeInfo(elf.Personnage)
}

func GobelinInfo(gob Gobelin) {
	fmt.Printf("	Gobelin\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", gob.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", gob.VieMax) + "\n")
	fmt.Print("Niveau de menace : " + fmt.Sprintf("%.2f", gob.NiveauDeMenace) + "\n")
	//fmt.Print("Armes : " + gob.arme.Nom + "\n")
	//fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.arme.Degats) + "\n\n")
}

func ElfeInfo(perso Personnage) {
	fmt.Printf("	Elfe\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", elf.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", elf.VieMax) + "\n")
	fmt.Print("Affinité avec la nature : " + fmt.Sprintf("%.2f", elf.affiniteNature) + "\n")
	fmt.Print("Longévité : " + fmt.Sprintf("%d", elf.longevite) + "\n")
	fmt.Print("Armes : " + perso.Arme.Nom + "\n")
	fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
}

func (c *Carte) obtenirContenuCase(x, y int) (contenu string) {
	if x < 0 || x >= c.largeur || y < 0 || y >= c.hauteur {
		return "Hors de la carte"
	}

	caseCourante := c.grille[y][x]

	if caseCourante.Personnage == PersonnageSelect {
		contenu = "Vous êtes ici"
	}
	if caseCourante.Personnage != nil && caseCourante.Personnage != PersonnageSelect {
		contenu = caseCourante.Personnage.Classe
	} else if caseCourante.Monstre != nil {
		contenu = "Monstre: " + caseCourante.Monstre.Classe
	} else {
		contenu = "Case vide"
	}

	return
}

func MesStatistiques() {
	switch PersonnageSelect.Classe {
	case "Magicien":
		AfficherStats_Mage(mage.Personnage)
	case "Chevalier":
		AfficherStats_Chevalier(chevalier.Personnage)
	case "Nain":
		AfficherStats_Nain(nain.Personnage)
	case "Elfe":
		ElfeInfo(elf.Personnage)
	default:
		fmt.Println("Pas de statistiques pour cette classe")
	}
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
		PersonnageSelect = &mage.Personnage
	case 2:
		PersonnageSelect = &chevalier.Personnage
	case 3:
		PersonnageSelect = &elf.Personnage
	case 4:
		PersonnageSelect = &nain.Personnage
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

func (c *Carte) placerMonstre(m *Monstre, x, y int) {
	if x < 0 || x >= c.largeur || y < 0 || y >= c.hauteur {
		fmt.Println("Position hors de la carte pour le monstre.")
		return
	}
	c.grille[y][x].Monstre = m
}

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

func PlacerObjetsSurCarte(carte *Carte) {
	rand.Seed(42) // Utilisez une graine fixe pour la génération aléatoire
	for x := 0; x < carte.largeur; x++ {
		for y := 0; y < carte.hauteur; y++ {
			if rand.Float64() < Probabilite {
				if rand.Intn(2) == 0 {
					carte.grille[x][y].Contenu = AddPotionDeSoin(1)
				} else {
					carte.grille[x][y].Contenu = AddCoteDePorc(1)
				}
			}
		}
	}
}

var test []*Personnage

// Fonctions similaires pour placer des monstres et des items
func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/joueur", getJoueur)

	go func() {
		err := http.ListenAndServe(":8002", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	inventaire = Inventaire{
		items:      make([]items.Item, 0),
		taille:     0,
		taille_max: 15,
	}

	inventaire.AddItems(AddCoteDePorc(3))
	inventaire.AddItems(AddPotionDeSoin(20))
	inventaire.AddItems(AddPotionDeSoin(200))

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

	PlacerObjetsSurCarte(&carte)

	var choix int

	choisirClasse()
	test = carte.obtenirAutresPersonnages(PersonnageSelect)
	clearConsole()

	fmt.Println("Vous avez choisi : " + PersonnageSelect.Classe)
	carte.afficher()

	for {
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
				autresPersonnages := carte.obtenirAutresPersonnages(PersonnageSelect)
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
			var message string = carte.deplacerSiPossible(PersonnageSelect, x, y)
			fmt.Println(message)
			carte.deplacerAutresPersonnages(PersonnageSelect)
			carte.afficher()

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

func getJoueur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&test)
}
