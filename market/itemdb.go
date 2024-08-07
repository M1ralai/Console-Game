package market

type weapon struct {
	name               string
	damage             int
	armour_penetration int
	agility            int
	cost               int
	owning             bool
	using              bool
}
type armour struct {
	name    string
	armour  int
	agility int
	cost    int
	owning  bool
}

var Weapon_Name []string
var Weapon_Damage, Weapon_Armour_penetration, Weapon_Agility, Weapon_Cost []int
var Weapon_Owning, Weapon_Using []bool

func Itemdb_init() {
	One_handed_wooden_sword := weapon{name: "One Handed Wooden Sword", damage: 25, armour_penetration: 5, agility: 20, cost: 50}
	One_handed_copper_sword := weapon{name: "One Handed Copper Sword", damage: 50, armour_penetration: 10, agility: 15, cost: 100}
	One_handed_iron_sword := weapon{name: "One Handed Iron Sword", damage: 100, armour_penetration: 20, agility: 10, cost: 200}
	One_handed_obsidian_sword := weapon{name: "One Handed Obsidian Sword", damage: 200, armour_penetration: 40, agility: 5, cost: 400}
	Two_handed_wooden_sword := weapon{name: "Two Handed Wooden Sword", damage: 50, armour_penetration: 10, agility: 10, cost: 100}
	Two_handed_copper_sword := weapon{name: "Two Handed Copper Sword", damage: 100, armour_penetration: 25, agility: 5, cost: 200}
	Two_handed_iron_sword := weapon{name: "Two Handed Iron Sword", damage: 200, armour_penetration: 50, agility: 0, cost: 400}
	Two_handed_obsidian_sword := weapon{name: "Two Handed Obsidian Sword", damage: 500, armour_penetration: 125, agility: 0, cost: 1000}
	Weapon_Name = []string{One_handed_wooden_sword.name, One_handed_copper_sword.name, One_handed_iron_sword.name, One_handed_obsidian_sword.name, Two_handed_wooden_sword.name, Two_handed_copper_sword.name, Two_handed_iron_sword.name, Two_handed_obsidian_sword.name}
	Weapon_Damage = []int{One_handed_wooden_sword.damage, One_handed_copper_sword.damage, One_handed_iron_sword.damage, One_handed_obsidian_sword.damage, Two_handed_wooden_sword.damage, Two_handed_copper_sword.damage, Two_handed_iron_sword.damage, Two_handed_obsidian_sword.damage}
	Weapon_Armour_penetration = []int{One_handed_wooden_sword.armour_penetration, One_handed_copper_sword.armour_penetration, One_handed_iron_sword.armour_penetration, One_handed_obsidian_sword.armour_penetration, Two_handed_wooden_sword.armour_penetration, Two_handed_copper_sword.armour_penetration, Two_handed_iron_sword.armour_penetration, Two_handed_obsidian_sword.armour_penetration}
	Weapon_Agility = []int{One_handed_wooden_sword.agility, One_handed_copper_sword.agility, One_handed_iron_sword.agility, One_handed_obsidian_sword.agility, Two_handed_wooden_sword.agility, Two_handed_copper_sword.agility, Two_handed_iron_sword.agility, Two_handed_obsidian_sword.agility}
	Weapon_Cost = []int{One_handed_wooden_sword.cost, One_handed_copper_sword.cost, One_handed_iron_sword.cost, One_handed_obsidian_sword.cost, Two_handed_wooden_sword.cost, Two_handed_copper_sword.cost, Two_handed_iron_sword.cost, Two_handed_obsidian_sword.cost}
	Weapon_Owning = []bool{One_handed_wooden_sword.owning, One_handed_copper_sword.owning, One_handed_iron_sword.owning, One_handed_obsidian_sword.owning, Two_handed_wooden_sword.owning, Two_handed_copper_sword.owning, Two_handed_iron_sword.owning, Two_handed_obsidian_sword.owning}
	Weapon_Using = []bool{One_handed_wooden_sword.using, One_handed_copper_sword.using, One_handed_iron_sword.using, One_handed_obsidian_sword.using, Two_handed_wooden_sword.using, Two_handed_copper_sword.using, Two_handed_iron_sword.using, Two_handed_obsidian_sword.using}
}
