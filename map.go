package main

import (
	"fmt"

	"github.com/fatih/color"
)

/* Types */

// Sector contains information about sector status.
type Sector struct {
	ID     int  // The sector's number.
	Status byte // Is it empty, or is taken by player.
}

// Map contains statuses of all its sectors.
type Map [MapHeight][MapWidth]Sector

// Coords type contains coorginates on the game map.
type Coords struct {
	Row int
	Col int
}

/* Package scope variables */

// MapSectors is the map for the game
var MapSectors Map

/* Package methods */

// Init sets up initial state of the map.
func (m *Map) Init() {
	var num int = 0
	for i := range m {
		for j := range m[i] {
			num++
			m[i][j].ID = num
			m[i][j].Status = SectEmp
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

// IsAvailable checks if there are still free sectors on the map.
func (m *Map) IsAvailable() bool {
	for i := range m {
		for j := range m[i] {
			if m[i][j].Status == SectEmp {
				return true
			}
		}
	}
	return false
}

// ToCoords converts map sector number into map coordinates.
func ToCoords(n int) (c Coords) {
	c.Row = (n - 1) / MapWidth
	c.Col = (n - 1) % MapWidth
	return c
}

/* Functions for the local use */

// printSector reads the sector's status and prints string with its condition.
func printSector(s Sector) {
	switch s.Status {
	case SectEmp:
		fmt.Printf("%3d", s.ID)
	case SectHum:
		c := color.New(color.FgGreen)
		c.Printf("%3d", s.ID)
		c.DisableColor()
	case SectBot:
		c := color.New(color.FgBlue)
		c.Printf("%3d", s.ID)
		c.DisableColor()
	case SectHum + SectBot:
		c := color.New(color.FgRed)
		c.Print(" HB")
		c.DisableColor()
	default:
		panic("unrecognized sector status")
	}
}
