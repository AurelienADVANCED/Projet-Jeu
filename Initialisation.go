package main

import (
	items "GoProjet/Items"
	"math/rand"
)

//Initialisation

// Liste de toutes les armes
var toutesLesArmes []items.Arme

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

func InitialisationDesItems() {
	// Ajoute toutes les armes à la liste
	toutesLesArmes = []items.Arme{
		{Nom: "Bâton de Magie", Degats: 35.0, Classe: "Magicien", Default: true},
		{Nom: "Sceptre Ancien", Degats: 60.0, Classe: "Magicien", Default: false},
		{Nom: "Bâton de Glace", Degats: 45.0, Classe: "Magicien", Default: false},
		{Nom: "Bâton de Feu Enchanté", Degats: 70.0, Classe: "Magicien", Default: false},

		{Nom: "Épée Longue", Degats: 50.0, Classe: "Chevalier", Default: true},
		{Nom: "Épée du Chevalier Noir", Degats: 75.0, Classe: "Chevalier", Default: false},
		{Nom: "Lance de Justicier", Degats: 65.0, Classe: "Chevalier", Default: false},
		{Nom: "Épée à Deux Mains", Degats: 85.0, Classe: "Chevalier", Default: false},

		{Nom: "Arc Long", Degats: 40.0, Classe: "Elfe", Default: true},
		{Nom: "Arc de la Lune", Degats: 55.0, Classe: "Elfe", Default: false},
		{Nom: "Arc Court Agile", Degats: 35.0, Classe: "Elfe", Default: false},
		{Nom: "Arc Enchanté", Degats: 60.0, Classe: "Elfe", Default: false},

		{Nom: "Hache de Guerre", Degats: 60.0, Classe: "Nain", Default: true},
		{Nom: "Marteau des Profondeurs", Degats: 80.0, Classe: "Nain", Default: false},
		{Nom: "Hache à Deux Mains", Degats: 75.0, Classe: "Nain", Default: false},
		{Nom: "Pioche de Combat", Degats: 65.0, Classe: "Nain", Default: false},

		{Nom: "Fronde", Degats: 20.0, Classe: "Gobelin", Default: true},
		{Nom: "Dague Empoisonnée", Degats: 30.0, Classe: "Gobelin", Default: false},
		{Nom: "Arc Court de Gobelin", Degats: 25.0, Classe: "Gobelin", Default: false},
		{Nom: "Bâton Tordu", Degats: 20.0, Classe: "Gobelin", Default: false},

		{Nom: "Massue", Degats: 55.0, Classe: "Orc", Default: true},
		{Nom: "Hache du Chef Orc", Degats: 70.0, Classe: "Orc", Default: false},
		{Nom: "Épée Large d'Orc", Degats: 65.0, Classe: "Orc", Default: false},
		{Nom: "Marteau de Guerre Orc", Degats: 80.0, Classe: "Orc", Default: false},
	}

}

func AddViande(Quantite int) *items.Nourriture {
	return CreerNourriture("Morceau de viande", Quantite, 2, 20.0, "🥩")
}

func AddPotionDeSoin(Quantite int) *items.Nourriture {
	return CreerNourriture("Potion de Soin", Quantite, 2, 50.0, "🧪")
}

// Initialisation des personnages

func CreerMagicien() Magicien {
	vieAleatoire := rand.Float64() * 100
	return Magicien{
		Personnage: Personnage{
			Vie:              vieAleatoire,
			VieMax:           vieAleatoire,
			Arme:             GetDefaultWeaponByClass("Magicien"),
			Mana:             150.0,
			ManaMax:          150.0,
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
			Arme:             GetDefaultWeaponByClass("Chevalier"),
			Mana:             0.0,
			ManaMax:          0.0,
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
			Arme:             GetDefaultWeaponByClass("Elfe"),
			Mana:             rand.Float64() * 50,
			ManaMax:          50,
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
			Arme:             GetDefaultWeaponByClass("Nain"),
			Mana:             0, // Nains n'utilisent généralement pas de Mana
			ManaMax:          0,
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
			Arme:             GetDefaultWeaponByClass("Gobelin"),
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
			Arme:             GetDefaultWeaponByClass("Orc"),
			NiveauDeMenace:   rand.Float64() * 10,
			Classe:           "Orc",
			RayonDeplacement: 1,
		},
	}
}

func GetDefaultWeaponByClass(classe string) *items.Arme {
	for _, arme := range toutesLesArmes {
		if arme.Default && arme.Classe == classe {
			return &arme
		}
	}
	return nil
}
