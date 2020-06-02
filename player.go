package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/* Types */

// Coords type contains coorginates on the game map.
type Coords struct {
	Row int
	Col int
}

// Player type.
type Player struct {
	Name     string
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
	p.Name = "Bot Player"
	p.IsBot = true
	rand.Seed(time.Now().UnixNano())
}

// SetHuman marks the player as human.
func (p *Player) SetHuman() {
	p.Name = "Human Player"
	p.IsBot = false
}

// InitPos sets up the initial player's sector.
func (p *Player) InitPos() {
	var choice int
	if !p.IsBot {
		// The human player initial sector.
		fmt.Println("Choose sector to start:")
		choice := selectSector(MapMax)
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
		// do...while loop analogue
		for {
			tmpLoc = toCoords(selectSector(MapMax))
			err := checkNextMove(p.Location, tmpLoc)
			if err != nil {
				fmt.Println(err)
				fmt.Print("Select the sector next to your current location: ")
				continue
			}
			reSetPlayerLocation(p, tmpLoc)
			break
		}
	} else {
		for {
			tmpLoc = toCoords(rand.Intn(MapMax) + 1)
			err := checkNextMove(p.Location, tmpLoc)
			if err != nil {
				continue
			}
			reSetPlayerLocation(p, tmpLoc)
			break
		}
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
	if p.IsBot {
		m[p.Location.Row][p.Location.Col].Status -= SectBot
	} else {
		m[p.Location.Row][p.Location.Col].Status -= SectHum
	}
}

// Fight simulates the fight and determines the winner among the players.
func Fight(players ...*Player) {
	winner := rand.Intn(len(players))
	fmt.Println("The winner is", players[winner].Name)
}

/* Functions for the local use */

func selectSector(lim int) (input int) {
	// Scan int from stdin.
	input = getInt()
	// Check whether ch is in range of map, if it's not reprompt the user.
	for input < 1 || input > lim {
		fmt.Printf("incorrect value, try with <1-%d>: ", lim)
		input = getInt()
	}
	return input
}

func getInt() (n int) {
	// Create a new bufio.Reader to read keyboard input.
	reader := bufio.NewReader(os.Stdin)
	// Read all the user input untill newline character.
	input, err := reader.ReadString('\n')
	checkErr(err)
	// Remove the whitespaces (the newline chararcter).
	input = strings.TrimSpace(input)
	n, err = strconv.Atoi(input)
	checkErr(err)
	return n
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Sets a new user location; clears the previous sector.
func reSetPlayerLocation(p *Player, newLoc Coords) {
	p.ClearSector(&MapSectors)
	p.Location = newLoc
	p.CaptureSector(&MapSectors)
}

// Converts map sector number into map coordinates.
func toCoords(n int) (c Coords) {
	c.Row = (n - 1) / MapWidth
	c.Col = (n - 1) % MapWidth
	return c
}

// Checks if the player can make a move to a new sector.
func checkNextMove(plrLoc, newLoc Coords) (err error) {
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
