package renderer

type ScreenPosition struct {
	X float32
	Y float32
}

const (
	FLOAT_DIVISOR          float32 = 2.0
	COORDW_BAD_VALUE       float32 = 1.0
	SCREEN_MOVE_FULLSCRENN float32 = 22
)

func WorldToScreen(renderer Renderer, x float32, y float32, z float32) ScreenPosition {
	var screenPosition ScreenPosition
	coordX := x*renderer.ViewProjMatrix[0] + y*renderer.ViewProjMatrix[4] + z*renderer.ViewProjMatrix[8] + renderer.ViewProjMatrix[12]
	coordY := x*renderer.ViewProjMatrix[1] + y*renderer.ViewProjMatrix[5] + z*renderer.ViewProjMatrix[9] + renderer.ViewProjMatrix[13]
	coordW := x*renderer.ViewProjMatrix[3] + y*renderer.ViewProjMatrix[7] + z*renderer.ViewProjMatrix[11] + renderer.ViewProjMatrix[15]

	if coordW < COORDW_BAD_VALUE {
		coordW = COORDW_BAD_VALUE
	}

	middleX := coordX / coordW
	middleY := coordY / coordW

	screenX := (1920 / FLOAT_DIVISOR * middleX) + (middleX + 1920/FLOAT_DIVISOR)
	screenY := -(1080 / FLOAT_DIVISOR * middleY) + (middleY + 1080/FLOAT_DIVISOR)

	screenPosition.X = screenX
	screenPosition.Y = screenY - SCREEN_MOVE_FULLSCRENN

	return screenPosition
}

func OnScren(renderer Renderer, x float32, y float32, z float32) bool {
	var offsetX float32 = 0
	var offsetY float32 = 0
	return x > -offsetX && x < (float32(renderer.Width)+offsetX) && y > -offsetY && y < (float32(renderer.Height)+offsetY)
}
