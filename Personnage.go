package main

import (
	"math/rand"
	"time"
)

// Déclaration des ensembles d'armes comme variables globales
// var armesMagicien = []string{"Bâton magique", "Sceptre", "Grimoire"}
// var armesChevalier = []string{"Épée", "Lance", "Bouclier"}
// var armesElfe = []string{"Arc", "Arbalète", "Dague elfique"}
// var armesNain = []string{"Hache", "Marteau", "Pioche de combat"}
// var armesGobelin = []string{"Dague rouillée", "Fronde", "Lance-pierre"}
// var armesOrc = []string{"Massue", "Épée à deux mains", "Hache de guerre"}

/*
Si l'armure est deux fois plus grande que les degats du personnage, alors les degats sont annulés
Exemple: armure = 20, degats = 10, alors degats = 0
Exemple: armure = 20, degats = 30, alors degats = 10
Sinon les degats sont réduits de la valeur de l'armure

Si l'agilité du personnage est supérieure à celle de l'ennemi, chaque niveau supérieur d'agilité donnera 5% de chance de plus
d'éviter le coup de l'ennemie.

Si le mana du personnage tombe a 0 il subira des dégats de 10 par tour.
*/

// Dérivés de Personnage
type Magicien struct {
	Personnage
	niveauDeMagie int
}

type Chevalier struct {
	Personnage
	armure        string
	codeDeHonneur string
}

type Elfe struct {
	Personnage
	affiniteNature float64
	longevite      int
}

type Nain struct {
	Personnage
	expertiseEnForge string
	resistanceAlcool float64
}

type Gobelin struct {
	Monstre
	fourberie float64
	vitesse   float64
}

type Orc struct {
	Monstre
	brutalite  float64
	leadership float64
}

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialise le générateur de nombres aléatoires
}

func AddVie(Personnage *Personnage, vie float64) bool {
	if Personnage.Vie == Personnage.VieMax {
		return true
	}
	if Personnage.Vie+vie > Personnage.VieMax {
		Personnage.Vie = Personnage.VieMax
	} else {
		Personnage.Vie = Personnage.Vie + vie
	}
	return false
}
