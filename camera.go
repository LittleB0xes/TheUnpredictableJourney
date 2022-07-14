package main

type Camera struct {
	x       float64
	y       float64
	width   float64
	height  float64
	marginX float64
	marginY float64
}

func createCamera(marginX, marginY float64, w, h int) Camera {
	return Camera{
		x:       0,
		y:       0,
		width:   float64(w),
		height:  float64(h),
		marginX: marginX,
		marginY: marginY,
	}
}

func (camera *Camera) Update(x, y float64) {
	camera.x = x - 0.5*camera.width
	camera.y = 0
	if camera.x < 0 {
		camera.x = 0
	}
	//camera.y = y - 0.5*camera.height

}
