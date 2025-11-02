package main

import (
	. "home_sweet_home/peterlib"
)

func drawRooftop(sideLength int) {
	Color("red")
	Down()
	Pivote(30)
	Forward(sideLength)
	Pivote(120)
	Forward(sideLength)
	Pivote(120)
	Forward(sideLength)
	Up()
}

func drawWalls(sideLength int) {
	Color("blue")
	Down()
	Left()
	Forward(sideLength)
	Left()
	Forward(sideLength)
	Left()
	Forward(sideLength)
	Up()
}

func drawHouse(sideLength int) {
	drawRooftop(sideLength)
	drawWalls(sideLength)
}

func preparePosition() {
	East()
	Forward(1)
	North()
}

func drawHouses(houseCount int, sideLength int) {
	houseIndex := 0
	for houseIndex < houseCount {
		drawHouse(sideLength)
		preparePosition()
		houseIndex += 1
	}
}

func main() {
	drawHouses(3, 5)

	// --- Lancer l’animation ---
	Play()
}
