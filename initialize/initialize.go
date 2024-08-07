package initialize

import (
	"github.com/M1ralai/RPG/level"
	"github.com/M1ralai/RPG/market"
	"github.com/M1ralai/RPG/stats"
)

func Initialize() {
	level.Level_init()
	market.Itemdb_init()
	stats.Stat_initialize()
}
