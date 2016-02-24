package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	game.Screen().SetLevel(level)
	game.Start()
}
