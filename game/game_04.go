package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	entity *tl.Entity
}

func (player *Player) Draw(screen *tl.Screen) {
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		x, y := player.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.entity.SetPosition(x+1, y)
		case tl.KeyArrowLeft:
			player.entity.SetPosition(x-1, y)
		case tl.KeyArrowUp:
			player.entity.SetPosition(x, y-1)
		case tl.KeyArrowDown:
			player.entity.SetPosition(x, y+1)
		}
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	player := Player{
		entity: tl.NewEntity(1, 1, 1, 1),
	}
	// Set the character at position (0, 0) on the entity.
	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
}
