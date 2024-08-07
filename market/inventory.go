package market

import (
	"fmt"
	"os"

	"github.com/M1ralai/RPG/stats"
	"github.com/M1ralai/RPG/supposed"
)

func Inventory_Visualization() {
	for {
		fmt.Println("Which type of inventory you looking for? \n 1. Weapon Inventory \n 2. Armour Inventory")
		supposed.Default_Quit()
		fmt.Scanln(&decision)
		switch {
		case decision == 1:
			Weapon_Inventory()
		case decision == 2:
		}
		if decision == 9 {
			break
		} else if decision == 10 {
			fmt.Println("Goodbye Sir")
			os.Exit(1)
		}

	}
}
func Weapon_Inventory() {
	for i := 0; i <= 7; i++ {
		if Weapon_Owning[i] {
			if Weapon_Using[i] {
				fmt.Println(i+1, ". you have and using ", Weapon_Name[i], " right now.")
			} else {
				fmt.Println(i+1, ". you have ", Weapon_Name[i])
			}
		} else {
			fmt.Println(i+1, ". you don't have ", Weapon_Name[i])
		}
	}
	fmt.Println("You can choose one of you have in inventory.")
	supposed.Default_Quit()
	for {
		fmt.Scanln(&decision)
		if decision <= 8 {
			for i := 0; i <= 7; i++ {
				if decision-1 == i {
					if Weapon_Owning[i] {
						fmt.Println("You equiped: ", Weapon_Name[i])
						stats.Item_Damage = Weapon_Damage[i]
						stats.Item_armour_penetration = Weapon_Armour_penetration[i]
						stats.Item_WAgility = Weapon_Agility[i]
						for a := 0; a <= 7; a++ {
							Weapon_Using[a] = false
						}
						Weapon_Using[i] = true
					} else {
						fmt.Println("You dont have. ", Weapon_Name[i])
					}
				}
			}
		} else if decision == 9 {
			break
		} else if decision == 10 {
			fmt.Println("Goodbye sir.")
			os.Exit(1)
		} else {
			fmt.Println("Please choose one of them")
		}
	}
}
