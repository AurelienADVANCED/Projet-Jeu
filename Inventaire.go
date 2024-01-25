package main

import (
	items "GoProjet/Items"
	"fmt"
)

var inventaire Inventaire

func (inv *Inventaire) NombreTotalItemsStackes() int {
	total := 0
	for _, item := range inv.items {
		total += item.GetStack()
	}
	return total
}

func (inv *Inventaire) AddItems(itemToAdd items.Item) {
	for i, item := range inventaire.items {
		// Vérifier si l'item existe déjà dans l'inventaire et si c'est le même type
		if item.GetNom() == itemToAdd.GetNom() && item.TypeItem() == itemToAdd.TypeItem() {
			spaceAvailable := item.StackMax() - item.GetStack()

			if itemToAdd.GetStack() <= spaceAvailable {
				// Utiliser une méthode SetGetStack pour modifier la quantité
				item.SetStack(item.GetStack() + itemToAdd.GetStack())
				inventaire.items[i] = item
				return
			} else {
				// Remplir le stack actuel et ajouter le reste après
				item.SetStack(item.StackMax())
				inventaire.items[i] = item
				// Créer un nouvel item avec la quantité restante
				remainingItem := item.Clone()
				remainingItem.SetStack(itemToAdd.GetStack() - spaceAvailable)
				inventaire.AddItems(remainingItem)
				return
			}
		}
	}

	// Si l'item n'existe pas dans l'inventaire, l'ajouter si il y a de la place
	if inventaire.taille < inventaire.taille_max {
		inventaire.items = append(inventaire.items, itemToAdd)
		inventaire.taille++
	}
}

func RemoveItems(itemToRemove items.Item, quantityToRemove int) {
	for i := 0; i < len(inventaire.items); {
		item := inventaire.items[i]
		if item.GetNom() == itemToRemove.GetNom() {
			currentStack := item.GetStack()
			if quantityToRemove < currentStack {
				// Réduire la quantité de l'item
				item.SetStack(currentStack - quantityToRemove)
				inventaire.items[i] = item // Mettre à jour l'item dans l'inventaire
				return
			} else {
				// Retirer l'item complètement et ajuster la quantité à retirer
				quantityToRemove -= currentStack
				inventaire.items = append(inventaire.items[:i], inventaire.items[i+1:]...)
				inventaire.taille--
				// Pas besoin d'incrémenter i car on a retiré un élément
			}
		} else {
			i++ // Incrémenter seulement si on n'a pas retiré d'élément
		}
	}
}

func AfficherInventaire() {
	for {
		viderCacheClavier()
		clearConsole()
		fmt.Println("Taille de l'inventaire:", inventaire.taille, "/", inventaire.taille_max)
		// Récuperer le nombre d'item en tout dans l'inventaire
		nombreTotalItems := inventaire.NombreTotalItemsStackes()
		fmt.Println("Nombre total d'items:", nombreTotalItems)
		fmt.Println()
		for index, item := range inventaire.items {
			fmt.Printf("Item N°%d: %s x %d\n", index+1, item.GetNom(), item.GetStack())
		}

		var choix int
		fmt.Println("Choisissez un item à utiliser (0 pour annuler):")
		fmt.Scan(&choix)
		if choix == 0 {
			return
		} else if choix > 0 && choix <= len(inventaire.items) {
			item := inventaire.items[choix-1]
			if item.TypeItem() == 1 {
				// Arme
			} else if item.TypeItem() == 2 {
				// Nourriture
				AddVie(*PersonnageSelect, item.(items.Nourriture).VieRecup)
				RemoveItems(item, choix-1)
			}

		}
	}
}

type Inventaire struct {
	items      []items.Item
	taille     int
	taille_max int
}
