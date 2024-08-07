package main

import (
	"github.com/M1ralai/RPG/hub"
	"github.com/M1ralai/RPG/intro"
	"github.com/M1ralai/RPG/market"
)

func main() {
	market.Itemdb_init()
	intro.Lore()
	hub.Main_Hub()
}
