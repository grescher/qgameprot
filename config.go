package main

const (
	/* Map Dimensions */

	// MapWidth determines the map width.
	MapWidth = 5
	// MapHeight determines the map height.
	MapHeight = 5
	// MapMax - number of the last sector on the map.
	MapMax = MapHeight * MapWidth

	/* Bitwise constants for the map's sector status */

	// SectEmp - sector is clear (nobody owns it).
	SectEmp = 0x0
	// SectHum - sector is utilized by the human player.
	SectHum = 0x1
	// SectBot - sector is utilized by the bot player.
	SectBot = 0x2
)
