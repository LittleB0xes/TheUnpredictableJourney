package main

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type LevelChunck struct {
	width         int
	height        int
	collisionGrid []int
	mapData       []int
	backData      []int
}

func createLevelChunck(w, h int) LevelChunck {
	grid := make([]int, w*h)
	levelMap := make([]int, w*h)
	backMap := make([]int, w*h)

	// Background randomization
	for i := 0; i < 200; i++ {
		sx := 10 + rand.Intn(3)
		sy := rand.Intn(3)

		x := rand.Intn(w)
		y := rand.Intn(h)

		backMap[x+y*w] = sx + 32*sy
	}

	for i := 0; i < w*h; i++ {
		if i/w == h-1 {
			grid[i] = 1
			levelMap[i] = 258 + 32 + rand.Intn(4)

		} else if i/w == h-2 {
			grid[i] = 1
			levelMap[i] = 258 + rand.Intn(4)
		}
	}

	// second path
	for i := 0; i < w*h; i++ {
		// create the test ladder
		if i%w == i/w && (i/w)%3 == 0 {
			grid[i] = 1
			levelMap[i] = 258
		}
	}
	// some pit
	pits_number := 4
	min_dist := 10
	max_dist := 20
	previous_pit_x := 0
	for i := 0; i < pits_number; i++ {

		// pit creation
		new_pit_x := previous_pit_x + min_dist + rand.Intn(max_dist-min_dist)

		// pit with variable length
		pit_width := 2 + rand.Intn(2)
		for j := 0; j < pit_width; j++ {
			for k := 0; k < 2; k++ {
				grid[new_pit_x+j+(h-2+k)*w] = 0
				levelMap[new_pit_x+j+(h-2+k)*w] = 0
			}
		}
		previous_pit_x = new_pit_x

	}

	return LevelChunck{
		width:         w,
		height:        h,
		collisionGrid: grid,
		mapData:       levelMap,
		backData:      backMap,
	}

}

func (chunck *LevelChunck) Draw(screen *ebiten.Image, tileset *ebiten.Image, camX, camY float64) {

	for index, tile := range chunck.backData {
		if tile != 0 {
			op := &ebiten.DrawImageOptions{}

			// 32is the tileset width (in tile)
			sourceX := 8 * (tile % 32)
			sourceY := 8 * (tile / 32)
			op.GeoM.Translate(float64(index%chunck.width)*8-0.25*camX, float64(index/chunck.width)*8-0.25*camY)
			screen.DrawImage(tileset.SubImage(image.Rect(sourceX, sourceY, sourceX+8, sourceY+8)).(*ebiten.Image), op)

		}

	}

	for index, tile := range chunck.mapData {
		if tile != 0 {
			op := &ebiten.DrawImageOptions{}

			// 32is the tileset width (in tile)
			sourceX := 8 * (tile % 32)
			sourceY := 8 * (tile / 32)
			op.GeoM.Translate(float64(index%chunck.width)*8-camX, float64(index/chunck.width)*8-camY)
			screen.DrawImage(tileset.SubImage(image.Rect(sourceX, sourceY, sourceX+8, sourceY+8)).(*ebiten.Image), op)

		}

	}

}
