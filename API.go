package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Position struct {
	X int
	Y int
}

func InitilisationAPI() {
	http.HandleFunc("/joueur", getJoueur)
	http.HandleFunc("/joueurs", getAllPlayer)
	http.HandleFunc("/carte", getCarte)
	http.HandleFunc("/inventaire", getInventaire)
	http.HandleFunc("/position", GetPositionPlayer)
	http.HandleFunc("/deplacer", setPlayerPosition)

	go func() {
		err := http.ListenAndServe(":8002", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func getCarte(w http.ResponseWriter, r *http.Request) {

	mapInfo, err := json.Marshal(carte)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	// Répondre avec les informations de la carte au format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapInfo)
}

func getJoueur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&PersonnageSelected)
}

func getAllPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&OtherPersonnage)
}

func getInventaire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&inventaire)
}

func OpenInventaire(w http.ResponseWriter, r *http.Request) {

}

func GetPositionPlayer(w http.ResponseWriter, r *http.Request) {
	xOrig, yOrig, trouve := carte.trouverPositionPersonnage(PersonnageSelected)
	if trouve {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&Position{xOrig, yOrig})
	} else {
		http.Error(w, "Erreur lors de la recherche de la position du joueur", http.StatusInternalServerError)
	}

}

func setPlayerPosition(w http.ResponseWriter, r *http.Request) {

	var position Position
	err := json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la position du joueur", http.StatusInternalServerError)
		return
	}

	message := carte.deplacerSiPossible(PersonnageSelected, position.X, position.Y)

	clearConsole()
	carte.afficher()
	displayMenu()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&message)

}
