package main

import (
	"testing"
)

// TestAddVieMaxVie simule l'ajout de vie à un personnage dont la vie est déjà au maximum.
func TestAddVieMaxVie(t *testing.T) {
	var p Personnage
	p.Vie = 100
	p.VieMax = 100
	vieAjoutee := 10.0
	expected := true

	result := AddVie(&p, vieAjoutee)
	if result != expected {
		t.Errorf("AddVie(%v, %v) = %v; want %v", p, vieAjoutee, result, expected)
	}
	if p.Vie != p.VieMax {
		t.Errorf("La vie du personnage devrait être à son maximum (%v), got %v", p.VieMax, p.Vie)
	}
}
