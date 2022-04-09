package gui

import (
	"framework-memory-go/src/renderer"
	"framework-memory-go/src/unitmanager"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var simpleShader *ebiten.Shader

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
	emptyImage.Fill(color.Black)
}

func DrawChamps(screen *ebiten.Image) {
	purpleClr := color.RGBA{255, 0, 255, 255}
	for i := 0; i < len(unitmanager.UNITMANAGER.Champions); i++ {
		ran := unitmanager.UNITMANAGER.Champions[i].AttackRange
		pos := unitmanager.UNITMANAGER.Champions[i].Position
		posScreen := renderer.WorldToScreen(renderer.RENDERER, pos.X, pos.Y, pos.Z)
		DrawCircle(screen, posScreen.X, posScreen.Y, ran, purpleClr)
	}
}

func DrawCircle(screen *ebiten.Image, x float32, y float32, radius float32, clr color.RGBA) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	var path vector.Path

	path.MoveTo(x, y)
	path.Arc(x, y, radius, 0, math.Pi*2, vector.Clockwise)
	path.Arc(x, y, radius-5, 0, math.Pi*2, vector.Clockwise)
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

var (
	emptyImage = ebiten.NewImage(3, 3)
)
