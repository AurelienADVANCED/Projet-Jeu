package main

import (
	"fmt"
)

func AfficherStats_Mage(perso Personnage) {
	fmt.Printf("	Mage\n")
	if mage.Vie <= 0 {
		fmt.Print("Personnage est mort \n\n")
	} else {
		fmt.Print("Vie : " + fmt.Sprintf("%.2f", mage.Vie) + "\n")
		fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", mage.VieMax) + "\n")
		fmt.Print("Mana : " + fmt.Sprintf("%.2f", mage.Mana) + "\n")
		fmt.Print("Mana Max : " + fmt.Sprintf("%.2f", mage.ManaMax) + "\n")
		fmt.Print("Force : " + fmt.Sprintf("%.2f", mage.Force) + "\n")
		fmt.Print("Agilité : " + fmt.Sprintf("%.2f", mage.Agilite) + "\n")
		fmt.Print("Armure : " + fmt.Sprintf("%.2f", mage.Armure) + "\n")
		fmt.Print("Classe : " + mage.Classe + "\n")
		fmt.Print("Niveau de magie : " + fmt.Sprintf("%d", mage.niveauDeMagie) + "\n\n")
		fmt.Print("Armes : " + perso.Arme.Nom + "\n")
		fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
		afficherStatutPersonnage(mage.Personnage, true)
	}
}

func AfficherStats_Chevalier(perso Personnage) {
	fmt.Printf("	Chevalier\n")
	if chevalier.Vie <= 0 {
		fmt.Print("Personnage est mort \n\n")
	} else {
		fmt.Print("Vie : " + fmt.Sprintf("%.2f", chevalier.Vie) + "\n")
		fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", chevalier.VieMax) + "\n")
		fmt.Print("Mana : " + fmt.Sprintf("%.2f", chevalier.Mana) + "\n")
		fmt.Print("Mana Max : " + fmt.Sprintf("%.2f", chevalier.ManaMax) + "\n")
		fmt.Print("Force : " + fmt.Sprintf("%.2f", chevalier.Force) + "\n")
		fmt.Print("Agilité : " + fmt.Sprintf("%.2f", chevalier.Agilite) + "\n")
		fmt.Print("Classe : " + chevalier.Classe + "\n\n")
		fmt.Print("Armes : " + perso.Arme.Nom + "\n")
		fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
		afficherStatutPersonnage(chevalier.Personnage, true)
	}
}

func AfficherStats_Nain(perso Personnage) {
	fmt.Printf("	Nain\n")
	if nain.Vie <= 0 {
		fmt.Print("Personnage est mort \n\n")
	} else {
		fmt.Print("Vie : " + fmt.Sprintf("%.2f", nain.Vie) + "\n")
		fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", nain.VieMax) + "\n")
		fmt.Print("Mana : " + fmt.Sprintf("%.2f", nain.Mana) + "\n")
		fmt.Print("Mana Max : " + fmt.Sprintf("%.2f", nain.ManaMax) + "\n")
		fmt.Print("Force : " + fmt.Sprintf("%.2f", nain.Force) + "\n")
		fmt.Print("Agilité : " + fmt.Sprintf("%.2f", nain.Agilite) + "\n")
		fmt.Print("Armure : " + fmt.Sprintf("%.2f", nain.Armure) + "\n")
		fmt.Print("Classe : " + nain.Classe + "\n")
		fmt.Print("Expertise en forge : " + nain.expertiseEnForge + "\n")
		fmt.Print("Résistance à l'alcool : " + fmt.Sprintf("%.2f", nain.resistanceAlcool) + "\n\n")
		fmt.Print("Armes : " + perso.Arme.Nom + "\n")
		fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
		afficherStatutPersonnage(nain.Personnage, true)
	}
}

func GobelinInfo(gob Gobelin) {
	fmt.Printf("	Gobelin\n")
	fmt.Print("Vie : " + fmt.Sprintf("%.2f", gob.Vie) + "\n")
	fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", gob.VieMax) + "\n")
	fmt.Print("Niveau de menace : " + fmt.Sprintf("%.2f", gob.NiveauDeMenace) + "\n")
}

func ElfeInfo(perso Personnage) {
	fmt.Printf("	Elfe\n")
	if elf.Vie <= 0 {
		fmt.Print("Personnage est mort \n\n")
	} else {
		fmt.Print("Vie : " + fmt.Sprintf("%.2f", elf.Vie) + "\n")
		fmt.Print("Vie Max : " + fmt.Sprintf("%.2f", elf.VieMax) + "\n")
		fmt.Print("Affinité avec la nature : " + fmt.Sprintf("%.2f", elf.affiniteNature) + "\n")
		fmt.Print("Longévité : " + fmt.Sprintf("%d", elf.longevite) + "\n")
		fmt.Print("Armes : " + perso.Arme.Nom + "\n")
		fmt.Print("Dégats : " + fmt.Sprintf("%.2f", perso.Arme.Degats) + "\n\n")
		afficherStatutPersonnage(elf.Personnage, true)
	}
}

func MesStatistiques() {
	switch PersonnageSelected.Classe {
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

func afficherStatistiques() {
	clearConsole()
	AfficherStats_Mage(mage.Personnage)
	AfficherStats_Chevalier(chevalier.Personnage)
	AfficherStats_Nain(nain.Personnage)
	ElfeInfo(elf.Personnage)
}
