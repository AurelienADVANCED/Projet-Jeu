package main

import (
	items "GoProjet/Items"
	"math/rand"
)

//Initialisation

// Initialisation des armes pour différentes classes

func CreerNourriture(nom string, stack, stackMax int, vieRecup float64, symbole string) *items.Nourriture {
	return &items.Nourriture{
		Nom:       nom,
		Stack:     stack,
		Stack_max: stackMax,
		VieRecup:  vieRecup,
		Symbole:   symbole,
	}
}

func CreerArme(nom string, degats float64, classe string) items.Arme {
	return items.Arme{Nom: nom, Degats: degats, StackMax: 1, Classe: classe}
}

// Armes Magicien
func ArmeDeBasePourMagicien() items.Arme {
	return CreerArme("Bâton de Magie", 35.0, "Magicien")
}

func SceptreAncien() items.Arme {
	return CreerArme("Sceptre Ancien", 60.0, "Magicien")
}

func BatonDeGlace() items.Arme {
	return CreerArme("Bâton de Glace", 45.0, "Magicien")
}

func BatonDeFeuEnchante() items.Arme {
	return CreerArme("Bâton de Feu Enchanté", 70.0, "Magicien")
}

// Chevalier
func ArmeDeBasePourChevalier() items.Arme {
	return CreerArme("Épée Longue", 50.0, "Chevalier")
}

func EpeeDeChevalierNoir() items.Arme {
	return CreerArme("Épée du Chevalier Noir", 75.0, "Chevalier")
}

func LanceDeJusticier() items.Arme {
	return CreerArme("Lance de Justicier", 65.0, "Chevalier")
}

func EpeeADeuxMains() items.Arme {
	return CreerArme("Épée à Deux Mains", 85.0, "Chevalier")
}

// Elfe
func ArmeDeBasePourElfe() items.Arme {
	return CreerArme("Arc Long", 40.0, "Elfe")
}

func ArcDeLaLune() items.Arme {
	return CreerArme("Arc de la Lune", 55.0, "Elfe")
}

func ArcCourtAgile() items.Arme {
	return CreerArme("Arc Court Agile", 35.0, "Elfe")
}

func ArcEnchante() items.Arme {
	return CreerArme("Arc Enchanté", 60.0, "Elfe")
}

// Nain
func ArmeDeBasePourNain() items.Arme {
	return CreerArme("Hache de Guerre", 60.0, "Nain")
}

func MarteauDesProfondeurs() items.Arme {
	return CreerArme("Marteau des Profondeurs", 80.0, "Nain")
}

func HacheADeuxMains() items.Arme {
	return CreerArme("Hache à Deux Mains", 75.0, "Nain")
}

func PiocheDeCombat() items.Arme {
	return CreerArme("Pioche de Combat", 65.0, "Nain")
}

// Gobelin
func ArmeDeBasePourGobelin() items.Arme {
	return CreerArme("Fronde", 20.0, "Gobelin")
}

func DagueEmpoisonnee() items.Arme {
	return CreerArme("Dague Empoisonnée", 30.0, "Gobelin")
}

func ArcCourtDeGobelin() items.Arme {
	return CreerArme("Arc Court de Gobelin", 25.0, "Gobelin")
}

func BatonTordu() items.Arme {
	return CreerArme("Bâton Tordu", 20.0, "Gobelin")
}

// Orc
func ArmeDeBasePourOrc() items.Arme {
	return CreerArme("Massue", 55.0, "Orc")
}

func HacheDuChefOrc() items.Arme {
	return CreerArme("Hache du Chef Orc", 70.0, "Orc")
}

func EpeeLargeDOrc() items.Arme {
	return CreerArme("Épée Large d'Orc", 65.0, "Orc")
}

func MarteauDeGuerreOrc() items.Arme {
	return CreerArme("Marteau de Guerre Orc", 80.0, "Orc")
}

func AddCoteDePorc(Quantite int) *items.Nourriture {
	return CreerNourriture("Côte de Porc", Quantite, 100, 20.0, "C")
}

func AddPotionDeSoin(Quantite int) *items.Nourriture {
	return CreerNourriture("Potion de Soin", Quantite, 100, 50.0, "P")
}

// Initialisation des personnages

func CreerMagicien() Magicien {
	vieAleatoire := rand.Float64() * 100
	return Magicien{
		Personnage: Personnage{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourMagicien(),
			Mana:             150.0,
			Force:            20.0,
			Agilite:          30.0,
			Armure:           10.0,
			Classe:           "Magicien",
			RayonDeplacement: 3,
		},
		niveauDeMagie: 5,
	}
}

func CreerChevalier() Chevalier {
	vieAleatoire := rand.Float64() * 150
	return Chevalier{
		Personnage: Personnage{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourChevalier(),
			Mana:             0.0,
			Force:            60.0,
			Agilite:          40.0,
			Armure:           50.0,
			Classe:           "Chevalier",
			RayonDeplacement: 2,
		},
		codeDeHonneur: "Protéger les Faibles",
	}
}

func nouveauElfe() Elfe {
	vieAleatoire := rand.Float64() * 80
	return Elfe{
		Personnage: Personnage{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourElfe(),
			Mana:             rand.Float64() * 50, // Ajuster selon la nature de l'Elfe
			Force:            rand.Float64() * 10,
			Agilite:          rand.Float64() * 20,
			Classe:           "Elfe",
			RayonDeplacement: 2,
		},
	}
}

func nouveauNain() Nain {
	vieAleatoire := rand.Float64() * 100
	return Nain{
		Personnage: Personnage{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourNain(),
			Mana:             0, // Nains n'utilisent généralement pas de Mana
			Force:            rand.Float64() * 15,
			Agilite:          rand.Float64() * 5,
			Classe:           "Nain",
			RayonDeplacement: 1,
		},
	}
}

func nouveauGobelin() Gobelin {
	vieAleatoire := rand.Float64() * 50
	return Gobelin{
		Monstre: Monstre{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourGobelin(),
			NiveauDeMenace:   rand.Float64() * 5,
			Classe:           "Gobelin",
			RayonDeplacement: 3,
		},
	}
}

func nouveauOrc() Orc {
	vieAleatoire := rand.Float64() * 150
	return Orc{
		Monstre: Monstre{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             ArmeDeBasePourOrc(),
			NiveauDeMenace:   rand.Float64() * 10,
			Classe:           "Orc",
			RayonDeplacement: 1,
		},
	}
}
