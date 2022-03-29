package renderer

type ScreenPosition struct {
	X float32
	Y float32
}

const (
	FLOAT_DIVISOR    uint32  = 2
	COORDW_BAD_VALUE float32 = 1
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

	screenX := (float32(renderer.Width/FLOAT_DIVISOR) * middleX) + (middleX + float32(renderer.Width/FLOAT_DIVISOR))
	screenY := -(float32(renderer.Height/FLOAT_DIVISOR) * middleY) + (middleY + float32(renderer.Height/FLOAT_DIVISOR))

	screenPosition.X = screenX
	screenPosition.Y = screenY

	return screenPosition
}

func OnScren(renderer Renderer, x float32, y float32, z float32) bool {
	var offsetX float32 = 0
	var offsetY float32 = 0
	return x > -offsetX && x < (float32(renderer.Width)+offsetX) && y > -offsetY && y < (float32(renderer.Height)+offsetY)
}
