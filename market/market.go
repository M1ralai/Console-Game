package market

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/M1ralai/RPG/supposed"
)

var decision, Gold int

func Armour_Market() {
	fmt.Println("Which part of armour are you interested? \n 1. Head armour \n 2. Chestplate armour \n 3. Leg armour \n ")
}
func Weapons_market() {
	Gold = 300000
	fmt.Printf("%sName%sDamage%sPenetration%sAgility%sCost%s\n", strings.Repeat(".", 5), strings.Repeat(".", 26), strings.Repeat(".", 4), strings.Repeat(".", 3), strings.Repeat(".", 6), strings.Repeat(".", 8))
	for i := 0; i <= 7; i++ {
		fmt.Println(i+1, ". ", Weapon_Name[i], strings.Repeat(".", 30-len(Weapon_Name[i])), Weapon_Damage[i], strings.Repeat(".", 10-len(strconv.Itoa(Weapon_Damage[i]))), Weapon_Armour_penetration[i], strings.Repeat(".", 10-len(strconv.Itoa(Weapon_Armour_penetration[i]))), Weapon_Agility[i], strings.Repeat(".", 10-len(strconv.Itoa(Weapon_Agility[i]))), Weapon_Cost[i], strings.Repeat(".", 10-len(strconv.Itoa(Weapon_Cost[i]))))
	}
	fmt.Println("Which one do you wanna buy?")
	supposed.Default_Quit()
	for {
		fmt.Scanln(&decision)
		if decision <= 8 {
			for i := 0; i <= 7; i++ {
				if i == decision-1 {
					if Gold < Weapon_Cost[i] {
						fmt.Println("You dont have enough money to buy ", Weapon_Name[i])
					} else {
						Weapon_Owning[i] = true
						fmt.Println("You bought ", Weapon_Name[i], "do you wanna equip it?\n 1. Yes \n 2. No ")
						for {
							fmt.Scanln(&decision)
							if decision == 1 {
								fmt.Println("You equiped ", Weapon_Name[i])
								for a := 7; a <= 7; a++ {
									Weapon_Using[a] = false
								}
								Weapon_Using[i] = true
								break
							} else if decision == 2 {
								fmt.Println("Okay you going back to market menu")
								break
							} else {
								fmt.Println("Please choose one of them.")
							}
						}
					}
					break
				}
			}
			break
		} else if decision == 9 {
			break
		} else if decision == 10 {
			fmt.Println("Goodbye sir.")
			os.Exit(1)
		} else {
			fmt.Println("Please select one of the options.")
		}
	}
}
