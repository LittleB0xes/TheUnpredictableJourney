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

	// Floor Generation
	for i := 0; i < w*h; i++ {
		if i/w == h-1 {
			grid[i] = 1
			levelMap[i] = 258 + 32 + rand.Intn(4)

		} else if i/w == h-2 {
			grid[i] = 1
			levelMap[i] = 258 + rand.Intn(4)
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
	// Add some levelling
	x_pos := 0
	for x_pos < w {
		current_hight := h - 2
		if rand.Intn(100) < 20 {
			max_length := 4 + rand.Intn(6)
			for tile := 0; tile < max_length; tile++ {

				// Dont build over pit
				if grid[x_pos+tile+current_hight*w] == 1 && x_pos+tile < w {
					grid[x_pos+tile+(current_hight-1)*w] = 1
					levelMap[x_pos+tile+(current_hight-1)*w] = 258 + rand.Intn(4)

				}
			}
			x_pos += max_length + 3

		} else {
			x_pos += 1
		}

	}

	x_pos = 0
	for x_pos < w {
		current_hight := h - 3
		if rand.Intn(100) < 35 {
			max_length := 4 + rand.Intn(6)
			for tile := 0; tile < max_length; tile++ {

				// Dont build over pit
				if grid[x_pos+tile+current_hight*w] == 1 && grid[x_pos+tile-1+current_hight*w] == 1 && x_pos+tile < w {
					grid[x_pos+tile+(current_hight-1)*w] = 1
					levelMap[x_pos+tile+(current_hight-1)*w] = 258 + rand.Intn(4)

				}
			}
			x_pos += max_length + 3

		} else {
			x_pos += 1
		}

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
