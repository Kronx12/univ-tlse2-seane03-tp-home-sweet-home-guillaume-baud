package main

import (
	. "home_sweet_home/peterlib"
)

func drawWalls(sideLength int) {
	Color("#0070c0")
	Down()
	South()
	Forward(sideLength)
	East()
	Forward(sideLength)
	North()
	Forward(sideLength)
	Up()
}

func drawRooftop(sideLength int) {
	Color("#ff0000")
	Down()
	Pivote(30)
	Forward(sideLength)
	Pivote(120)
	Forward(sideLength)
	Pivote(120)
	Forward(sideLength)
	Up()
}

func drawHouses(houseCount, sideLength int) {
	for range houseCount {
		drawRooftop(sideLength)
		drawWalls(sideLength)
		East()
		Forward(1)
		North()
	}
}

func main() {
	drawHouses(3, 5)

	// --- Lancer l’animation ---
	Play()
}
