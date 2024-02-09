package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// Permet d'attaquer un autre personnage
func attaquer(joueur *Personnage, adversaire *Personnage) string {
	adversaire.Vie -= joueur.Arme.Degats
	message :=
		"Degats infligés : " + fmt.Sprintf("%.2f", joueur.Arme.Degats) +
			"\nVie de l'adversaire : " + fmt.Sprintf("%.2f", adversaire.Vie) +
			"\n" + joueur.Classe + " : " + fmt.Sprintf("%.2f", joueur.Vie) + "\n"

	if adversaire.Vie <= 0 {
		carte.retirerPersonnage(adversaire)
		message += "\n L'adversaire a été vaincu et retiré de la carte."

	}
	return message
}

// permet d'attaquer des monstres
func AttackMonster(joueur *Personnage, adversaire *Monstre) string {
	adversaire.Vie -= joueur.Arme.Degats
	message :=
		"Degats infligés : " + fmt.Sprintf("%.2f", joueur.Arme.Degats) +
			"\nVie de l'adversaire : " + fmt.Sprintf("%.2f", adversaire.Vie) +
			"\n" + joueur.Classe + " : " + fmt.Sprintf("%.2f", joueur.Vie) + "\n"

	if adversaire.Vie <= 0 {
		carte.Grille[0][0].Monstre = nil
		message += "\n L'adversaire a été vaincu et retiré de la carte."

	}
	return message
}

// permet de retirer un personnage de la carte, par exemple lorsqu'il est vaincu
func (c *Carte) retirerPersonnage(p *Personnage) {
	for y := range c.Grille {
		for x := range c.Grille[y] {
			if c.Grille[y][x].Personnage == p {
				c.Grille[y][x].Personnage = nil
				return
			}
		}
	}
}

// Fonction pour placer un personnage sur la carte
func (c *Carte) placerPersonnage(p *Personnage, x, y int) {
	if x < 0 || x >= c.Largeur || y < 0 || y >= c.Hauteur {
		fmt.Println("Position hors de la carte.")
		return
	}
	c.Grille[y][x].Personnage = p
}

func (c *Carte) peutDeplacer(x, y int) bool {
	return x >= 0 && x < c.Largeur && y >= 0 && y < c.Hauteur &&
		c.Grille[y][x].Personnage == nil
}

func (c *Carte) trouverPositionPersonnage(p *Personnage) (int, int, bool) {
	for y, ligne := range c.Grille {
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
		switch {
		case message == "Vous êtes ici":
			return "Déplacement impossible. Vous êtes déjà ici."
		case message == "Case vide":
			c.deplacerPersonnage(p, xDest, yDest)
			return "Déplacement effectué."
		case message == "Hors de la carte":
			return "Déplacement impossible. Hors de la carte."
		case strings.Contains(message, "Objet"):
			// Récupérer l'objet
			item := c.Grille[yDest][xDest].Contenu
			inventaire.AddItems(item, yDest, xDest, c)
			return "Vous avez récupé un object"
		case strings.Contains(message, "Monstre"):
			monstres := map[string]MonstreRef{
				"Gobelin": {&gobelin.Monstre, "Gobelin", func() bool { return gobelin.Vie > 0 }},
				"Orc":     {&orc.Monstre, "Orc", func() bool { return orc.Vie > 0 }},
			}

			for key, val := range monstres {
				if strings.Contains(message, key) {
					if val.estVivant() {
						message := AttackMonster(p, val.monstre)
						return "Vous avez attaqué un " + val.nom + ".\n" + message
					} else {
						return "Le " + val.nom + " est mort"
					}
				}
			}
			return "Déplacement impossible. " + message
		case strings.Contains(message, "Personnage"):
			personnages := map[string]PersonnageRef{
				"Elfe":      {&elf.Personnage, "Elfe", func() bool { return elf.Vie > 0 }},
				"Magicien":  {&mage.Personnage, "Magicien", func() bool { return mage.Vie > 0 }},
				"Chevalier": {&chevalier.Personnage, "Chevalier", func() bool { return chevalier.Vie > 0 }},
				"Nain":      {&nain.Personnage, "Nain", func() bool { return nain.Vie > 0 }},
			}

			for key, val := range personnages {
				if strings.Contains(message, key) && p.Classe != key {
					if val.estVivant() {
						message := attaquer(p, val.personnage)
						return "Vous avez attaqué un " + val.nom + ".\n" + message
					} else {

						return "Le " + val.nom + " est mort"
					}
				}
			}
			return "Déplacement impossible. " + message
		default:
			return "Déplacement impossible. " + message
		}
	} else {
		return "Déplacement impossible. Hors de portée."
	}

}

func (c *Carte) obtenirAutresPersonnages(p *Personnage) []*Personnage {
	var autresPersonnages []*Personnage
	for _, ligne := range c.Grille {
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
			c.Grille[yOrig][xOrig].Personnage = nil
			c.Grille[yDest][xDest].Personnage = personnage

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

func (c *Carte) deplacerPersonnage(p *Personnage, xDest, yDest int) {
	if xDest < 0 || xDest >= c.Largeur || yDest < 0 || yDest >= c.Hauteur {
		fmt.Println("Position de destination invalide.")
		return
	}

	// Trouver la position actuelle du personnage
	var xOrig, yOrig int
	trouve := false
	for y := range c.Grille {
		for x := range c.Grille[y] {
			if c.Grille[y][x].Personnage == p {
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
	c.Grille[yDest][xDest].Personnage = p
	c.Grille[yOrig][xOrig].Personnage = nil
}
