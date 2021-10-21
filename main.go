package main

import (
	_ "image/png"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// Holds all game data
type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

// Create a new game object and initialize
func NewGame() *Game {
	g := &Game{}
	world, tags := InitializeWorld()
	g.Map = NewGameMap()
	g.WorldTags = tags
	g.World = world
	return g
}

// Update called each tic
func (g *Game) Update() error {

	TryMovePlayer(g)

	return nil
}

// Draw called each draw cylce
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

// Return screen dimensions
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
