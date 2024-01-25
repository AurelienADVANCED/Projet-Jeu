package main

import (
	items "GoProjet/Items"
	"math/rand"
)

//Initialisation

// Initialisation des armes pour différentes classes

// Armes Magicien
func ArmeDeBasePourMagicien() items.Arme {
	return items.Arme{Nom: "Bâton de Magie", Degats: 35.0, StackMax: 1, Classe: "Magicien"}
}

func SceptreAncien() items.Arme {
	return items.Arme{Nom: "Sceptre Ancien", Degats: 60.0, StackMax: 1, Classe: "Magicien"}
}

func BatonDeGlace() items.Arme {
	return items.Arme{Nom: "Bâton de Glace", Degats: 45.0, StackMax: 1, Classe: "Magicien"}
}

func BatonDeFeuEnchante() items.Arme {
	return items.Arme{Nom: "Bâton de Feu Enchanté", Degats: 70.0, StackMax: 1, Classe: "Magicien"}
}

// Chevalier
func ArmeDeBasePourChevalier() items.Arme {
	return items.Arme{Nom: "Épée Longue", Degats: 50.0, StackMax: 1, Classe: "Chevalier"}
}

func EpeeDeChevalierNoir() items.Arme {
	return items.Arme{Nom: "Épée du Chevalier Noir", Degats: 75.0, StackMax: 1, Classe: "Chevalier"}
}

func LanceDeJusticier() items.Arme {
	return items.Arme{Nom: "Lance de Justicier", Degats: 65.0, StackMax: 1, Classe: "Chevalier"}
}

func EpeeADeuxMains() items.Arme {
	return items.Arme{Nom: "Épée à Deux Mains", Degats: 85.0, StackMax: 1, Classe: "Chevalier"}
}

// Elfe
func ArmeDeBasePourElfe() items.Arme {
	return items.Arme{Nom: "Arc Long", Degats: 40.0, StackMax: 1, Classe: "Elfe"}
}

func ArcDeLaLune() items.Arme {
	return items.Arme{Nom: "Arc de la Lune", Degats: 55.0, StackMax: 1, Classe: "Elfe"}
}

func ArcCourtAgile() items.Arme {
	return items.Arme{Nom: "Arc Court Agile", Degats: 35.0, StackMax: 1, Classe: "Elfe"}
}

func ArcEnchante() items.Arme {
	return items.Arme{Nom: "Arc Enchanté", Degats: 60.0, StackMax: 1, Classe: "Elfe"}
}

// Nain
func ArmeDeBasePourNain() items.Arme {
	return items.Arme{Nom: "Hache de Guerre", Degats: 60.0, StackMax: 1, Classe: "Nain"}
}

func MarteauDesProfondeurs() items.Arme {
	return items.Arme{Nom: "Marteau des Profondeurs", Degats: 80.0, StackMax: 1, Classe: "Nain"}
}

func HacheADeuxMains() items.Arme {
	return items.Arme{Nom: "Hache à Deux Mains", Degats: 75.0, StackMax: 1, Classe: "Nain"}
}

func PiocheDeCombat() items.Arme {
	return items.Arme{Nom: "Pioche de Combat", Degats: 65.0, StackMax: 1, Classe: "Nain"}
}

// Gobelin
func ArmeDeBasePourGobelin() items.Arme {
	return items.Arme{Nom: "Fronde", Degats: 20.0, StackMax: 1, Classe: "Gobelin"}
}

func DagueEmpoisonnee() items.Arme {
	return items.Arme{Nom: "Dague Empoisonnée", Degats: 30.0, StackMax: 1, Classe: "Gobelin"}
}

func ArcCourtDeGobelin() items.Arme {
	return items.Arme{Nom: "Arc Court de Gobelin", Degats: 25.0, StackMax: 1, Classe: "Gobelin"}
}

func BatonTordu() items.Arme {
	return items.Arme{Nom: "Bâton Tordu", Degats: 20.0, StackMax: 1, Classe: "Gobelin"}
}

// Orc
func ArmeDeBasePourOrc() items.Arme {
	return items.Arme{Nom: "Massue", Degats: 55.0, StackMax: 1, Classe: "Orc"}
}

func HacheDuChefOrc() items.Arme {
	return items.Arme{Nom: "Hache du Chef Orc", Degats: 70.0, StackMax: 1, Classe: "Orc"}
}

func EpeeLargeDOrc() items.Arme {
	return items.Arme{Nom: "Épée Large d'Orc", Degats: 65.0, StackMax: 1, Classe: "Orc"}
}

func MarteauDeGuerreOrc() items.Arme {
	return items.Arme{Nom: "Marteau de Guerre Orc", Degats: 80.0, StackMax: 1, Classe: "Orc"}
}

func AddCoteDePorc(Quantite int) items.Nourriture {
	return items.Nourriture{Nom: "Côte de Porc", Stack: Quantite, Stack_max: 100, VieRecup: 20.0, Symbole: "C"}
}

func AddPotionDeSoin(Quantite int) items.Nourriture {
	return items.Nourriture{Nom: "Potion de Soin", Stack: Quantite, Stack_max: 100, VieRecup: 50.0, Symbole: "P"}
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
