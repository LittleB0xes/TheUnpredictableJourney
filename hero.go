package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Hero struct {
	x             float64
	y             float64
	vx            float64
	vy            float64
	dir_x         float64
	dir_y         float64
	jumpForce     float64
	collisionBox  Rect
	state         AnimationState
	animationList map[AnimationState]AnimationData
	sprite        AnimatedSprite
	onFloor       bool
}

func NewHero(xo, yo float64, sLib SpriteLib) Hero {

	data := make(map[AnimationState]AnimationData)
	data[Walk] = sLib["hero_walk"]
	data[Idle] = sLib["hero_idle"]
	return Hero{
		x:             xo,
		y:             yo,
		vx:            0,
		vy:            0,
		dir_x:         0,
		dir_y:         0,
		jumpForce:     4,
		collisionBox:  Rect{x: 0, y: 0, w: 8, h: 8},
		state:         Walk,
		animationList: data,
		sprite:        NewAnimatedSprite(xo, yo, "hero_idle", sLib),
		onFloor:       false,
	}
}

func (e *Hero) Update(level *LevelChunck) {

	current_state := e.state

	// Input checking
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		e.dir_x = -1
		e.state = Walk
		e.sprite.flipH = true

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		e.dir_x = 1
		e.state = Walk
		e.sprite.flipH = false
	} else {
		e.dir_x = 0
		e.dir_y = 0
		e.state = Idle
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) && e.onFloor {
		e.vy = -e.jumpForce
		e.onFloor = false
	}
	if e.state != current_state {
		e.sprite.SetAnimation(e.animationList[e.state])
	}

	// Some gravity
	e.vy += 0.25
	// Motion
	if e.dir_x != 0 {
		e.vx = e.dir_x
	} else {
		e.vx *= 0.9
	}
	//~~~~~~~~ IntGrid Collision Checking

	e.onFloor = false

	destXa := int((e.x + e.vx) / 8.0)
	destYa := int(e.y / 8.0)

	destXb := int(e.x / 8.0)
	destYb := int((e.y + e.vy) / 8.0)

	tiles := [4]([2]int){
		[2]int{0, 0},
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{1, 1},
	}

	// Collision on x
	for _, tile := range tiles {
		i := tile[0]
		j := tile[1]
		if e.HasCollision(destXa+i, destYa+j, e.vx, 0, level) != 0 {
			e.vx = 0
			break
		}
	}
	for _, tile := range tiles {
		i := tile[0]
		j := tile[1]
		if e.HasCollision(destXb+i, destYb+j, 0, e.vy, level) != 0 {
			if e.vy > 0 {
				e.onFloor = true
			}
			e.vy = 0

			break
		}
	}

	//if e.y > 130 {
	//	e.y = 130
	//	e.onFloor = true
	//}

	e.x += e.vx
	e.y += e.vy

	// 4 and 8 are her to adjust the "pivot" with collision
	// because the sprite is bigger than 8px
	e.sprite.x = e.x - 4
	e.sprite.y = e.y - 8
}
func (e *Hero) HasCollision(x, y int, dx, dy float64, level *LevelChunck) int {
	if x >= level.width || y >= level.height || x < 0 || y < 0 {
		return 0
	}
	value := level.collisionGrid[x+y*level.width]
	if (intersestRect(e.GetCollisionBox(dx, dy), Rect{x: float64(x * 8), y: float64(y * 8), w: 8, h: 8}) && value != 0) {
		return value
	} else {
		return 0
	}
}

func (e *Hero) GetCollisionBox(dx, dy float64) Rect {
	return Rect{
		x: e.x + e.collisionBox.x + dx,
		y: e.y + e.collisionBox.y + dy,
		w: e.collisionBox.w,
		h: e.collisionBox.h,
	}
}

func intersestRect(box1, box2 Rect) bool {
	r1_right := box1.x + box1.w
	r1_left := box1.x
	r1_top := box1.y
	r1_bottom := box1.y + box1.h

	r2_right := box2.x + box2.w
	r2_left := box2.x
	r2_top := box2.y
	r2_bottom := box2.y + box2.h

	collided := r1_left <= r2_right && r1_right >= r2_left && r1_top <= r2_bottom && r1_bottom >= r2_top
	return collided

}
