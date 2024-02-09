package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func InitilisationAPI() {
	http.HandleFunc("/joueur", getJoueur)
	http.HandleFunc("/joueurs", getAllPlayer)
	http.HandleFunc("/carte", getCarte)
	http.HandleFunc("/inventaire", getInventaire)

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

func setPlayerPosition(w http.ResponseWriter, r *http.Request) {

}
