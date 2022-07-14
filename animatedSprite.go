package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tile struct {
	sx int
	sy int
	sw int
	sh int
}

type AnimatedSprite struct {
	texture        *ebiten.Image
	x              float64
	y              float64
	w              float64
	h              float64
	anchorX        float64
	sx             int
	sy             int
	sw             int
	sh             int
	frame          int
	animationSpeed int
	elapsedTime    int
	currentFrame   int
	camX           float64
	camY           float64
	isPlaying      bool
	flipH          bool
}

func NewAnimatedSprite(xo, yo float64, name string, spriteLib SpriteLib) AnimatedSprite {
	data := spriteLib[name]
	path := data.Path

	img, _, err := ebitenutil.NewImageFromFile("./" + path)
	if err != nil {
		log.Fatal("AnimatedSprite - Error when openning file", err)
	}
	return AnimatedSprite{
		texture:        img,
		x:              xo,
		y:              yo,
		w:              float64(data.W),
		h:              float64(data.H),
		anchorX:        0.5,
		sx:             data.X,
		sy:             data.Y,
		sw:             data.W,
		sh:             data.H,
		frame:          data.Frame,
		animationSpeed: data.Speed,
		elapsedTime:    0,
		currentFrame:   0,
		camX:           0,
		camY:           0,
		isPlaying:      true,
		flipH:          false,
	}

}

func (s *AnimatedSprite) Draw(screen *ebiten.Image, camX, camY float64) {
	//var anchorX float64 = 0.5

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-s.anchorX*float64(s.w), 0)
	if s.flipH {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(s.x-camX+s.anchorX*float64(s.w), s.y-camY)
	source_x := s.sx + s.currentFrame*s.sw

	//Translate with camera and reset anchor translations
	screen.DrawImage(s.texture.SubImage(image.Rect(source_x, s.sy, source_x+s.sw, s.sy+s.sh)).(*ebiten.Image), op)
	if s.isPlaying {
		s.elapsedTime++
	}

	if s.isPlaying && s.elapsedTime >= s.animationSpeed {
		s.currentFrame++
		s.elapsedTime = 0
	}
	if s.currentFrame >= s.frame {
		s.currentFrame = 0
	}
}

func (s *AnimatedSprite) SetAnimation(data AnimationData) {
	s.sx = data.X
	s.sy = data.Y
	s.sw = data.W
	s.sh = data.H
	s.frame = data.Frame
	s.animationSpeed = data.Speed
	s.elapsedTime = 0
	s.currentFrame = 0
}

func (s *AnimatedSprite) Pause() {
	s.isPlaying = false
}

func (s *AnimatedSprite) Rewind() {
	s.currentFrame = 0
	s.elapsedTime = 0
}

func (s *AnimatedSprite) Play() {
	s.isPlaying = true
}

func (s *AnimatedSprite) SetSpeed(speed int) {
	s.animationSpeed = speed
}
func (s AnimatedSprite) GetRect() Rect {
	return Rect{
		x: s.x,
		y: s.y,
		w: s.w,
		h: s.h,
	}
}
