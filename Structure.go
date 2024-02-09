package main

import (
	items "GoProjet/Items"
)

type Personnage struct {
	Vie              float64
	VieMax           float64
	Arme             *items.Arme
	Mana             float64 // Pour les personnages pouvant utiliser de la magie
	ManaMax          float64
	Force            float64 // Pour les personnages physiques
	Agilite          float64 // Pour les personnages physiques
	Armure           float64
	Classe           string
	RayonDeplacement int
}

type Monstre struct {
	Vie              float64
	VieMax           float64
	Arme             *items.Arme
	NiveauDeMenace   float64
	Classe           string
	RayonDeplacement int
}

type PersonnageRef struct {
	personnage *Personnage
	nom        string
	estVivant  func() bool
}

type MonstreRef struct {
	monstre   *Monstre
	nom       string
	estVivant func() bool
}

type Case struct {
	Personnage *Personnage `json:"personnage"`
	Monstre    *Monstre    `json:"monstre"`
	Contenu    items.Item  `json:"contenu"`
}

type Carte struct {
	Grille  [][]Case `json:"grille"`
	Largeur int      `json:"largeur"`
	Hauteur int      `json:"hauteur"`
}
