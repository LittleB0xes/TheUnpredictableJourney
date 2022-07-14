package main

// Some useful enum or type
type AnimationState int64

const (
	Walk = iota
	Idle
)

type Vec2 struct {
	x float64
	y float64
}

func (v *Vec2) is_zero() bool {
	if v.x == 0.0 && v.y == 0.0 {
		return true
	} else {
		return false
	}
}

type Rect struct {
	x float64
	y float64
	w float64
	h float64
}

type Teleporter struct {
	destination  string
	destX        float64
	destY        float64
	auto         bool
	collisionBox Rect
}
