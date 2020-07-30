package main

import (
	"mygo/player"
	"mygo/referee"
)

func main() {
	Grade,Poker,Array,Name:=player.Player("小明")
	Grade2,Poker2,Array2,Name2:=player.Player("小刚")
	referee.Referee(Grade,Poker,Array,Name,Grade2,Poker2,Array2,Name2)
}
