package gui

import (
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/unitmanager"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var simpleShader *ebiten.Shader

const (
	NUMPOINTS           = 60
	STEP_VALUE          = 6.2831 / NUMPOINTS
	THETA_VALUE_INITIAL = 0.0
)

func init() {
	var err error

	simpleShader, err = ebiten.NewShader([]byte(`
		package gui

		func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
			return color
		}
	`))
	if err != nil {
		panic(err)
	}
	emptyImage.Fill(color.White)
}

func DrawCircle(screen *ebiten.Image, pos unitmanager.GamePosition, radius float32, tickness float32, clr color.RGBA) {
	var path vector.Path
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.EvenOdd}
	op.Address = ebiten.AddressUnsafe

	path = createCirclePath(pos, radius, path)
	path = createCirclePath(pos, radius-tickness, path)

	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	// screen.DrawTrianglesShader(vertices, indices, simpleShader, &ebiten.DrawTrianglesShaderOptions{FillRule: ebiten.EvenOdd})
	screen.DrawTriangles(vertices, indices, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)

}

var (
	emptyImage = ebiten.NewImage(3, 3)
)

func DrawCircleFilled(screen *ebiten.Image, pos unitmanager.GamePosition, radius float32, clr color.RGBA) {
	var path vector.Path
	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.EvenOdd}
	op.Address = ebiten.AddressUnsafe

	path = createCirclePath(pos, radius, path)

	vertices, indices := path.AppendVerticesAndIndicesForFilling(nil, nil)

	redScaled := float32(clr.R) / 255
	greenScaled := float32(clr.G) / 255
	blueScaled := float32(clr.B) / 255
	alphaScaled := float32(clr.A) / 255

	for i := range vertices {
		v := &vertices[i]

		v.ColorR = redScaled
		v.ColorG = greenScaled
		v.ColorB = blueScaled
		v.ColorA = alphaScaled
	}
	screen.DrawTrianglesShader(vertices, indices, simpleShader, &ebiten.DrawTrianglesShaderOptions{FillRule: ebiten.EvenOdd})
}

func createCirclePath(pos unitmanager.GamePosition, radius float32, path vector.Path) vector.Path {
	step := STEP_VALUE
	theta := THETA_VALUE_INITIAL

	firsPos := renderer.WorldToScreen(renderer.RENDERER,
		pos.X+radius*float32(math.Cos(theta)),
		pos.Y,
		pos.Z-radius*float32(math.Sin(theta)))
	path.MoveTo(firsPos.X, firsPos.Y)
	for theta < 6.2831 {
		posScreen := renderer.WorldToScreen(renderer.RENDERER,
			pos.X+radius*float32(math.Cos(theta)),
			pos.Y,
			pos.Z-radius*float32(math.Sin(theta)))
		path.LineTo(posScreen.X, posScreen.Y)
		theta += step
	}
	path.LineTo(firsPos.X, firsPos.Y)
	return path
}

func createCirclePath2(pos unitmanager.GamePosition, radius float32, path vector.Path) vector.Path {
	step := STEP_VALUE
	theta := THETA_VALUE_INITIAL

	firsPos := renderer.WorldToScreen(renderer.RENDERER,
		pos.X+radius*float32(math.Cos(theta)),
		pos.Y,
		pos.Z-radius*float32(math.Sin(theta)))
	path.MoveTo(firsPos.X, firsPos.Y)

	for theta < 6.2831 {
		posScreen := renderer.WorldToScreen(renderer.RENDERER,
			pos.X+radius*float32(math.Cos(theta)),
			pos.Y,
			pos.Z-radius*float32(math.Sin(theta)))
		path.Arc(posScreen.X, posScreen.Y, radius, 0, math.Pi/2, vector.Clockwise)
		theta += step
	}

	// path.Arc(firsPos.X, firsPos.Y, radius, 0, math.Pi/2, vector.Clockwise)
	return path
}
