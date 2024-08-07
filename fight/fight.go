package fight

import (
	"fmt"

	"github.com/M1ralai/RPG/level"
)

var decision int

func Scarab() {
	fmt.Println("You choosed scarab. Do you accept? \n 1. Yes \n 2. No")
	for {
		fmt.Scanln(&decision)
		if decision == 1 {
			Fight("scarab", 100+level.LVL*20, 25+level.LVL*25, 50+level.LVL*25)
		} else if decision == 2 {
			break
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}
func Zombie() {
	fmt.Println("You choosed zombie. Do you accept? \n 1. Yes \n 2. No")
	for {
		fmt.Scanln(&decision)
		if decision == 1 {
			Fight("zombie", 150+level.LVL*30, 100+level.LVL*35, 50+level.LVL*50)
		} else if decision == 2 {
			break
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}
func Skeleton() {
	fmt.Println("You choosed skeleton. Do you accept? \n 1. Yes \n 2. No")
	for {
		fmt.Scanln(&decision)
		if decision == 1 {
			Fight("skeleton", 150+level.LVL*50, 100+level.LVL*30, 100) //change last variable as (damage*85)/100
		} else if decision == 2 {
			break
		} else {
			fmt.Println("Please choose one of the options.")
		}
	}
}

func Fight(name string, monster_hp int, monster_damage int, monster_armour int) {

}
