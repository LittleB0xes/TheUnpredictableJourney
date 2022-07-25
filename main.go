package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	tileset   *ebiten.Image
	spriteLib SpriteLib
	camera    Camera
	hero      Hero
	level     LevelChunck
}

const (
	CanvasWidth  int = 320 //240 //320
	CanvasHeight int = 96  //180
)

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	img, _, err := ebitenutil.NewImageFromFile("./assets/tileset.png")
	if err != nil {
		log.Fatal("AnimatedSprite - Error when openning file", err)
	}
	sLib := loadSpriteLibrary("./assets/atlas.json")
	return &Game{
		tileset:   img,
		spriteLib: sLib,
		camera:    createCamera(0, 0, 96, 12),
		hero:      NewHero(0, 48, sLib),
		level:     createCave(96, 12),
		//level:     createLevelChunck(96, 12),
	}
}

func (g *Game) Update() error {
	g.hero.Update(&g.level)
	g.camera.Update(g.hero.x, g.hero.y)

	// Debug inputs
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		g.hero = NewHero(0, 20, g.spriteLib)
		g.level = createCave(96, 12)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0x2b, 0x2b, 0x26, 0xff})

	//ebitenutil.DebugPrint(screen, "Hello World")
	g.level.Draw(screen, g.tileset, g.camera.x, g.camera.y)
	g.hero.sprite.Draw(screen, g.camera.x, g.camera.y)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return CanvasWidth, CanvasHeight
	//return 240, 96
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("An Unpredictable Journey")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
