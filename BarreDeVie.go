package main

import (
	"fmt"
	"strings"
)

func genererBarre(valeurActuelle float64, valeurMaximale float64, couleurPleine, couleurVide string) string {
	if valeurMaximale <= 0 {
		return "Erreur: valeurMaximale doit être > 0"
	}
	proportionPleine := int((valeurActuelle / valeurMaximale) * 10)
	if proportionPleine < 0 {
		proportionPleine = 0 // Assurez-vous que proportionPleine n'est pas négative
	} else if proportionPleine > 10 {
		proportionPleine = 10 // Assurez-vous que proportionPleine ne dépasse pas 10
	}
	proportionVide := 10 - proportionPleine

	return strings.Repeat(couleurPleine, proportionPleine) + strings.Repeat(couleurVide, proportionVide)
}

func afficherStatutPersonnage(p Personnage) {

	barreVie := genererBarre(p.Vie, p.VieMax, "🟩", "🟥")
	affichageVie := fmt.Sprintf("%s %s %.2f/%.2f PV", p.Classe, barreVie, p.Vie, p.VieMax)
	fmt.Println(affichageVie)
	if p.ManaMax > 0 { // Affiche aussi la barre de mana si nécessaire
		barreMana := genererBarre(p.Mana, p.ManaMax, "🟦", "🟥")
		affichageMana := fmt.Sprintf("%s %s %.2f/%.2f Mana", p.Classe, barreMana, p.Mana, p.ManaMax)
		fmt.Println(affichageMana)
	}
}
