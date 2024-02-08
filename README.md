# Jeu de rôle simple en Go

Ce projet est un jeu de rôle simple développé en Go, où les joueurs peuvent explorer une carte, combattre des monstres, collecter des objets et améliorer leurs personnages. Le jeu est actuellement en cours de développement et est destiné à illustrer les concepts de base de la programmation en Go.

## Fonctionnalités actuelles

- Génération de carte aléatoire.
- Personnage joueur avec des statistiques de base.
- Monstres se déplaçant aléatoirement sur la carte.
- Combat entre le joueur et les monstres.
- Collecte d'objets (armes, nourriture, etc.).
- Inventaire pour stocker des objets.
- Amélioration des statistiques du joueur.

## Prérequis

Assurez-vous d'avoir Go installé sur votre système. Vous pouvez télécharger Go à partir du [site officiel](https://golang.org/).

## Comment jouer

1. Clonez ce dépôt sur votre machine :

git clone https://github.com/votre-utilisateur/jeu-go.git


2. Accédez au répertoire du jeu :

cd jeu-go


3. Exécutez le jeu :

go run .


4. Suivez les instructions à l'écran pour jouer au jeu.

## API (Application Programming Interface)

Le jeu dispose également d'une API pour récupérer des informations sur la carte. Voici comment l'utiliser :

- Récupérer toutes les informations de la carte :

GET /api/carte


Cette requête renverra un JSON contenant toutes les données de la carte, y compris la grille, les personnages, les monstres, etc.
