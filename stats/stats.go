package stats

import "fmt"

var Item_Damage, Item_Armour, Item_armour_penetration, Item_WAgility, Item_AAgility int

func Stat_initialize() {
	fmt.Println("Hello World")
}
func Stat_Visualize() {
	fmt.Printf("\nHere is your stats from items: \n Damage: %d \n Armour: %d \n Armour Penetration: %d \n Agility from weapon: %d  \n Agility from armour: %d \n", Item_Damage, Item_Damage, Item_armour_penetration, Item_WAgility, Item_AAgility)
}
