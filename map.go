package main

import (
	"fmt"
)

/* Types */

// Sector contains information about sector status.
type Sector struct {
	ID     int  // The sector's number.
	Status byte // Is it empty, or is taken by player.
}

// Map contains statuses of all its sectors.
type Map [MapHeight][MapWidth]Sector

/* Package scope variables */

// MapSectors is the map for the game
var MapSectors Map

/* Package methods */

// Init sets up initial state of the map.
func (m *Map) Init() {
	var num int
	for i := range m {
		for j := range m[i] {
			num++
			m[i][j].ID = num
		}
	}
}

// ShowMap shows up the current map condition with the players on it if any.
func (m *Map) ShowMap() {
	for i := range m {
		for j := range m[i] {
			printSector(m[i][j])
		}
		fmt.Println()
	}
}

/* Functions for the local use */

// printSector reads the sector's status and prints string with its condition.
func printSector(s Sector) {
	switch s.Status {
	case SectEmp:
		fmt.Printf("%3d", s.ID)
	case SectHum:
		fmt.Print("  H")
	case SectBot:
		fmt.Print("  B")
	case SectHum + SectBot:
		fmt.Print(" HB")
	default:
		panic("unrecognized sector status")
	}
}
