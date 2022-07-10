package main

import (
	"gogame/scenes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

const (
	LOGICAL_WIDTH  = 1920 / 6
	LOGICAL_HEIGHT = 1080 / 6

	WINDOW_WIDTH  = 1920 / 2
	WINDOW_HEIGHT = 1080 / 2
)

type Game struct {
	ecs ebiten.Game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ecs.Layout(LOGICAL_WIDTH, LOGICAL_HEIGHT)
}

func (g *Game) Update() error {
	// start := time.Now()

	err := g.ecs.Update()

	// fmt.Printf("%s took %v\n", "Update()", time.Since(start))

	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	// start := time.Now()

	g.ecs.Draw(screen)

	// fmt.Printf("%s took %v\n", "Draw()", time.Since(start))
}

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Go Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)

	g := &Game{ecs: engine.NewGame(&scenes.MainScene{})}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
