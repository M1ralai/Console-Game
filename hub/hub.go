package hub

import (
	"fmt"
	"os"

	"github.com/M1ralai/RPG/fight"
	"github.com/M1ralai/RPG/initialize"
	"github.com/M1ralai/RPG/level"
	"github.com/M1ralai/RPG/market"
	"github.com/M1ralai/RPG/supposed"
)

var decision int

func Main_Hub() {
	initialize.Initialize()
	for {
		fmt.Println("Welcome to this world's brave warrior, please choose where do you wanted to go. \n 1. Monster Hub \n 2. Marketplace \n 3. Level \n 4. Inventory")
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			Monster_Hub()
		case decision == 2:
			Market_Hub()
		case decision == 3:
			level.Level()
		case decision == 4:
			market.Inventory_Visualization()
		default:
			fmt.Println("Please choose one of the options")
		}
	}
}

func Monster_Hub() {
	for {
		fmt.Println("Welcome to the Monster Hub, \n 1. Scarab \n 2. Zombie \n 3. Skeleton")
		supposed.Default_Quit()
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			fight.Scarab()
		case decision == 2:
			fight.Zombie()
		case decision == 3:
			fight.Skeleton()
		}
		if decision == 9 {
			break
		} else if decision == 10 {
			fmt.Println("Goodbye sir.")
			os.Exit(1)
		}
	}
}
func Market_Hub() {
	for {
		fmt.Println("Where do ypu wanna go? \n 1. Weaponsmith \n 2. Armoursmith")
		supposed.Default_Quit()
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			market.Weapons_market()
		case decision == 2:
			market.Armour_Market()
		}
		if decision == 9 {
			break
		} else if decision == 10 {
			fmt.Println("Goodbye sir.")
			os.Exit(1)
		}
	}
}
