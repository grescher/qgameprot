package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

/* Types */

// Coords type contains coorginates on the game map.
type Coords struct {
	Row int
	Col int
}

// Player type.
type Player struct {
	Location Coords // Tracks in which sector the player currently is.
	IsBot    bool   // True if a player is bot.
}

/* Variables for the local use */

// initHumPos keeps the initial sector value for the human player.
// Used in the bot's first sector initialization.
var initHumPos int

/* Package methods */

// SetBot marks the player as bot.
func (p *Player) SetBot() {
	p.IsBot = true
}

// SetHuman marks the player as human.
func (p *Player) SetHuman() {
	p.IsBot = false
}

// InitPos sets up the initial player's sector.
func (p *Player) InitPos() {
	var choice int
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Choose sector to start:")
		choice := readUserInput(MapMax)
		initHumPos, p.Location = choice, toCoords(choice)
	} else {
		// The bot player initial sector.
		choice = rand.Intn(MapMax) + 1
		for choice == initHumPos {
			// Reassign location if it's equal to the human player location.
			choice = rand.Intn(MapMax) + 1
		}
		p.Location = toCoords(choice)
		fmt.Printf("The opponent's clan has picked sector %d\n", choice)
	}
	p.CaptureSector(&MapSectors)
}

// Move sets up the next player's position on the map.
func (p *Player) Move() {
	var tmpLoc Coords
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Input the sector number to move:")
		for {
			tmpLoc = toCoords(readUserInput(MapMax))
			err := checkNextMove(p.Location, tmpLoc)
			if err != nil {
				fmt.Println(err)
				fmt.Print("Select the sector next to your current location: ")
				continue
			}
			p.ClearSector(&MapSectors)
			p.Location = tmpLoc
			p.CaptureSector(&MapSectors)
			break
		}
	} else {

	}
}

// CaptureSector sets the status of the map sector for the respective player.
func (p *Player) CaptureSector(m *Map) {
	if p.IsBot {
		m[p.Location.Row][p.Location.Col].Status += SectBot
	} else {
		m[p.Location.Row][p.Location.Col].Status += SectHum
	}
}

// ClearSector marks the map sector as no-one's.
func (p *Player) ClearSector(m *Map) {
	m[p.Location.Row][p.Location.Col].Status = SectEmp
}

func readUserInput(lim int) (input int) {
	// Scan int from stdin.
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		log.Fatal(err)
	}
	// Check whether ch is in range of map, if it's not reprompt the user.
	for input < 1 || input > lim {
		fmt.Printf("incorrect value, try with <1-%d>: ", lim)
		_, err = fmt.Scanf("%d", &input)
		if err != nil {
			log.Fatal(err)
		}
	}
	return input
}

// Converts map sector number into map coordinates.
func toCoords(n int) (c Coords) {
	c.Row = (n - 1) / MapWidth
	c.Col = (n - 1) % MapWidth
	return c
}

// Checks if the player can make a move to a new sector.
func checkNextMove(plrLoc Coords, newLoc Coords) (err error) {
	diff := newLoc.Col - plrLoc.Col
	if diff > 1 || diff < -1 {
		return errors.New("unable to move through a sector")
	}
	diff = newLoc.Row - plrLoc.Row
	if diff > 1 || diff < -1 {
		return errors.New("unable to move through a sector")
	}
	return nil
}
