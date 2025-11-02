Données:
```
couleur du toit = rouge
couleur des murs = bleu
nombre de maison = 3

# Dimensions choisis arbitrairement
largeur maison = 5
```

Raffinage 1:
```
dessiner 3 maisons de 5 de large
```

Raffinage 2:
```
dessinerMaisons(nbrMaisons, largeurMaison):
	indexMaison = 0
	tant que indexMaison < nbrMaisons:
		dessiner une maison de largeurMaison de large
		se rendre à la position suivante
		incrémenter indexMaison
dessiner 3 maisons de 5 de large => dessinerMaisons(3, 5)
```

Raffinage 3:
```
dessinerMaison(largeur):
	dessiner un toit
	dessiner les murs
	
preparePosition():
	s'orienter vers l'Est
	avancer de 1
	s'orienter vers le Nord

dessinerMaisons(nbrMaisons, largeurMaison):
	indexMaison = 0
	tant que indexMaison < nbrMaisons:
		dessiner une maison de largeurMaison de large => dessinerMaison(5)
		se rendre à la position suivante => preparePosition()
		incrémenter indexMaison
dessinerMaisons(3, 5)
```

Raffinage 4:
```
dessinerToit(pLargeur):
	utiliser le rouge
	baisser le stylo
	tourner de 30 degrés
	avancer de pLargeur
	tourner de 120 degrés
	avancer de pLargeur
	tourner de 120 degrés
	lever le stylo
	
dessinerMurs(pLargeur):
	utiliser le bleu
	baisser le stylo
	tourner a gauche
	avancer de pLargeur
	tourner a gauche
	avancer de pLargeur
	tourner a gauche
	avancer de pLargeur
	lever le stylo

dessinerMaison(pLargeur):
	dessiner un toit de largeur pLargeur => dessinerToit(pLargeur)
	dessiner les murs de largeur pLargeur => dessinerMurs(pLargeur)
	
preparerPosition():
	s'orienter vers l'Est
	avancer de 1
	s'orienter vers le Nord

dessinerMaisons(nbrMaisons, largeurMaison):
	indexMaison = 0
	tant que indexMaison < nbrMaisons:
		dessinerMaison(largeurMaison)
		preparerPosition()
		incrémenter indexMaison
dessinerMaisons(3, 5)
```

Mise en algorithme:
```
dessinerToit(pLargeur):
	utiliser le rouge
	baisser le stylo
	tourner de 30 degrés
	avancer de pLargeur
	tourner de 120 degrés
	avancer de pLargeur
	tourner de 120 degrés
	lever le stylo

dessinerMurs(pLargeur):
	utiliser le bleu
	baisser le stylo
	tourner a gauche
	avancer de pLargeur
	tourner a gauche
	avancer de pLargeur
	tourner a gauche
	avancer de pLargeur
	lever le stylo

dessinerMaison(pLargeur):
	dessinerToit(pLargeur)
	dessinerMurs(pLargeur)
	
preparerPosition():
	s'orienter vers l'Est
	avancer de 1
	s'orienter vers le Nord

dessinerMaisons(nbrMaisons, largeurMaison):
	indexMaison = 0
	tant que indexMaison < nbrMaisons:
		dessinerMaison(largeurMaison)
		preparerPosition()
		indexMaison = indexMaison + 1

Algorithme Principal:
	dessinerMaisons(3, 5)
```