# Hangman Web


## Description
Hangman Web reprends la logique de jeux de hangman-classic, il est développée en Go (Golang) avec une interface web crée a partir d'éléments graphique entiérement dessiné par l'équipe. Le jeu permet aux joueurs de deviner des mots lettre par lettre, avec un nombre limité de tentatives, les lettres utilisées sont affichées.
Il y a 3 niveaux de dificultées disponnible dans le jeux "easy", "medium" et "hard".


## Technologies Utilisées
- Backend : Go (Golang). 
- Frontend : HTML, CSS.
- Krita : éléments graphique.


## Installation

1. Clonez le repository :
```bash
git clone https://github.com/jolanallen/hangman-web.git
cd hangman-web
```

## Lancement du Jeu

1. Démarrez le serveur :
```bash
go run main.go
```

2. Ouvrez votre navigateur et accédez à :
```
http://localhost:3030
```



## Fonctionnalités
- Interface web interactive
- Sélection aléatoire des mots
- Compteur de tentatives restantes
- Pages de victoire et de défaite personnalisées
- Trois niveau de dificultés
- Possibilité d'ajouter des mots personaliser dans les fichiers qui contiennent les liste de mots. 
- possibilité de modifier les dessin des état du hangman pour crée un hangman avec des états personalisé


## Comment Jouer
1. Lancez le serveur et accédez à la page d'accueil
2. Cliquez sur "Nouvelle Partie" pour commencer
3. Devinez les lettres en cliquant sur les touches ou en utilisant votre clavier
4. Vous avez un nombre limité de tentatives pour trouver le mot

## Contribution
Les contributions sont les bienvenues ! Pour contribuer :
1. Forkez le projet.
2. Créez une nouvelle branche (`git checkout -b contributeur`)
3. Committez vos changements (`git commit -m 'Ajout de fonctionnalité'`)
4. Pushez vers la branche (`git push origin contributeur`)
5. Ouvrez une Pull Request


## Auteurs
- Jolan allen, oscars gonsalez, noa, kéo.



