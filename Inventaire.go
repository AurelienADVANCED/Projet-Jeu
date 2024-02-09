package main

import (
	items "GoProjet/Items"
	"fmt"
	"math"
)

var inventaire Inventaire

func (inv *Inventaire) NombreTotalItemsStackes() int {
	total := 0
	for _, item := range inv.Items {
		total += item.GetStack()
	}
	return total
}

func (inv *Inventaire) AddItems(itemToAdd items.Item, xDest int, yDest int, c *Carte) {
	if c != nil {
		c.Grille[xDest][yDest].Contenu = nil
	}

	for i := 0; i < len(inv.Items) && itemToAdd.GetStack() > 0; i++ {
		item := inv.Items[i]
		// Vérifier si l'item existe déjà dans l'inventaire et si c'est le même type
		if item.GetNom() == itemToAdd.GetNom() && item.TypeItem() == itemToAdd.TypeItem() {
			spaceAvailable := item.StackMax() - item.GetStack()

			if itemToAdd.GetStack() <= spaceAvailable {
				// L'item peut être empilé dans l'item existant
				item.SetStack(item.GetStack() + itemToAdd.GetStack())
				inv.Items[i] = item
				return
			} else {
				// Ajouter autant que possible à l'item existant
				item.SetStack(item.StackMax())
				inv.Items[i] = item
				itemToAdd.SetStack(itemToAdd.GetStack() - spaceAvailable)
			}
		}
	}

	// Ajouter le reste en créant de nouveaux Items si nécessaire
	for itemToAdd.GetStack() > 0 && inv.Taille < inv.Taille_max {
		newStack := min(itemToAdd.GetStack(), itemToAdd.StackMax())
		newItem := itemToAdd.Clone()
		newItem.SetStack(newStack)
		inv.Items = append(inv.Items, newItem)
		inv.Taille++
		itemToAdd.SetStack(itemToAdd.GetStack() - newStack)
	}

	if itemToAdd.GetStack() > 0 {
		// Gérer le cas où l'inventaire est plein et qu'il reste encore des Items à ajouter
		fmt.Println("Inventaire plein, certains Items n'ont pas été ajoutés")

		// Ajouter les Items restants à la carte
		c.Grille[xDest][yDest].Contenu = itemToAdd
	}
}

// Fonction helper pour trouver le minimum de deux valeurs
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RemoveItems(itemToRemove items.Item, quantityToRemove int) {
	for quantityToRemove > 0 {
		// Identifier l'index de l'item avec le stack le plus petit
		minStackIndex := -1
		minStack := math.MaxInt32 // Utiliser la valeur maximale pour initialiser

		for i, item := range inventaire.Items {
			if item.GetNom() == itemToRemove.GetNom() && (minStackIndex == -1 || item.GetStack() < minStack) {
				minStackIndex = i
				minStack = item.GetStack()
			}
		}

		// Si aucun item correspondant n'est trouvé, sortir de la boucle
		if minStackIndex == -1 {
			break
		}

		item := inventaire.Items[minStackIndex]
		currentStack := item.GetStack()

		if quantityToRemove < currentStack {
			// Réduire la quantité de l'item
			item.SetStack(currentStack - quantityToRemove)
			inventaire.Items[minStackIndex] = item // Mettre à jour l'item dans l'inventaire
			return
		} else {
			// Retirer complètement l'item de l'inventaire
			quantityToRemove -= currentStack
			inventaire.Items = append(inventaire.Items[:minStackIndex], inventaire.Items[minStackIndex+1:]...)
			inventaire.Taille--
		}
	}
}

func AfficherInventaire() {
	for {
		viderCacheClavier()
		clearConsole()
		fmt.Println("Taille de l'inventaire:", inventaire.Taille, "/", inventaire.Taille_max)
		// Récuperer le nombre d'item en tout dans l'inventaire
		nombreTotalItems := inventaire.NombreTotalItemsStackes()
		fmt.Println("Nombre total d'Items:", nombreTotalItems)
		fmt.Println()
		for index, item := range inventaire.Items {
			fmt.Printf("Item  N°%d: %s x %d %s\n", index+1, item.GetNom(), item.GetStack(), item.GetSymbole())
		}

		var choix int
		fmt.Println("Choisissez un item à utiliser (0 pour annuler):")
		fmt.Scan(&choix)
		if choix == 0 {
			return
		} else if choix > 0 && choix <= len(inventaire.Items) {
			item := inventaire.Items[choix-1]
			if item.TypeItem() == 1 {
				// Arme
				switchWeapon(PersonnageSelected.Arme, item.(*items.Arme))

			} else if item.TypeItem() == 2 {
				// Nourriture
				AddVie(PersonnageSelected, item.(*items.Nourriture).VieRecup)
				RemoveItems(item, 1)
			}
		}
	}
}

func switchWeapon(weapon1 *items.Arme, weapon2 *items.Arme) {
	PersonnageSelected.Arme = weapon2
	RemoveItems(weapon2, 1)
	inventaire.AddItems(weapon1, 0, 0, nil)
}

type Inventaire struct {
	Items      []items.Item `json:"items"`
	Taille     int          `json:"taille"`
	Taille_max int          `json:"taille_max"`
}
