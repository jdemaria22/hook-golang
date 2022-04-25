package gui

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var disable = flag.Bool("disable", false, "disable all widgets")

func CreateUi(name string) {
	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(600)), app.Title(name), app.Decorated(true))
		if err := loop(w, name); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window, name string) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				if *disable {
					gtx = gtx.Disabled()
				}
				gui(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}

var (
	list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
	progress            = float32(0)
	progressIncrementer chan float32
	topLabel            = "Framework-hook-Go"
	CheckboxOrbwalker   = new(widget.Bool)
	CheckboxWards       = new(widget.Bool)
	CheckboxRanges      = new(widget.Bool)
	CheckboxSpells      = new(widget.Bool)
	CheckboxTraps       = new(widget.Bool)
	CheckboxLastHit     = new(widget.Bool)
	float               = new(widget.Float)
)

type (
	D = layout.Dimensions
	C = layout.Context
)

func init() {
	CheckboxOrbwalker.Value = true
	CheckboxWards.Value = true
	CheckboxRanges.Value = true
	CheckboxSpells.Value = true
	CheckboxTraps.Value = true
	CheckboxLastHit.Value = true
}

func gui(gtx layout.Context, th *material.Theme) layout.Dimensions {
	paint.Fill(gtx.Ops,
		color.NRGBA{R: 84, G: 184, B: 144, A: 0xFF},
	)
	widgets := []layout.Widget{
		func(gtx C) D {
			title := material.H4(th, topLabel)
			title.TextSize = unit.Dp(30)
			title.Color = color.NRGBA{R: 0, G: 88, B: 134, A: 0xFF}
			return layout.Center.Layout(gtx,
				title.Layout,
			)
		},
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Px(unit.Dp(10))
			gtx.Constraints.Max.Y = gtx.Constraints.Min.Y
			dr := image.Rectangle{Max: gtx.Constraints.Min}
			paint.LinearGradientOp{
				Stop1:  layout.FPt(dr.Min),
				Stop2:  layout.FPt(dr.Max),
				Color1: color.NRGBA{R: 0x10, G: 0xff, B: 0x10, A: 0xFF},
				Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
			}.Add(gtx.Ops)
			defer clip.Rect(dr).Push(gtx.Ops).Pop()
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{
				Size: gtx.Constraints.Max,
			}
		},
		func(gtx C) D {
			title := material.H6(th, "Orbwalker")
			title.TextSize = unit.Dp(20)
			title.Color = color.NRGBA{R: 0, G: 88, B: 134, A: 0xFF}
			return layout.Center.Layout(gtx,
				title.Layout,
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(
					material.CheckBox(th, CheckboxOrbwalker, "Obwalker").Layout,
				),
			)
		},
		func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(1, material.Slider(th, float, 0, 100).Layout),
				layout.Rigid(func(gtx C) D {
					return layout.UniformInset(unit.Dp(8)).Layout(gtx,
						material.Body1(th, fmt.Sprintf("%.2f", float.Value)).Layout,
					)
				}),
			)
		},
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Px(unit.Dp(10))
			gtx.Constraints.Max.Y = gtx.Constraints.Min.Y
			dr := image.Rectangle{Max: gtx.Constraints.Min}
			paint.LinearGradientOp{
				Stop1:  layout.FPt(dr.Min),
				Stop2:  layout.FPt(dr.Max),
				Color1: color.NRGBA{R: 0x10, G: 0xff, B: 0x10, A: 0xFF},
				Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
			}.Add(gtx.Ops)
			defer clip.Rect(dr).Push(gtx.Ops).Pop()
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{
				Size: gtx.Constraints.Max,
			}
		},
		func(gtx C) D {
			title := material.H6(th, "Drawings")
			title.TextSize = unit.Dp(20)
			title.Color = color.NRGBA{R: 0, G: 88, B: 134, A: 0xFF}
			return layout.Center.Layout(gtx,
				title.Layout,
			)
		},
		func(gtx C) D {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(material.CheckBox(th, CheckboxRanges, "Draw Ranges").Layout),
				layout.Rigid(material.CheckBox(th, CheckboxSpells, "Draw Spells").Layout),
				layout.Rigid(material.CheckBox(th, CheckboxWards, "Draw Wards").Layout),
				layout.Rigid(material.CheckBox(th, CheckboxTraps, "Draw Traps").Layout),
				layout.Rigid(material.CheckBox(th, CheckboxLastHit, "Draw Last Hit Minions").Layout),
			)
		},
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Px(unit.Dp(10))
			gtx.Constraints.Max.Y = gtx.Constraints.Min.Y
			dr := image.Rectangle{Max: gtx.Constraints.Min}
			paint.LinearGradientOp{
				Stop1:  layout.FPt(dr.Min),
				Stop2:  layout.FPt(dr.Max),
				Color1: color.NRGBA{R: 0x10, G: 0xff, B: 0x10, A: 0xFF},
				Color2: color.NRGBA{R: 0x10, G: 0x10, B: 0xff, A: 0xFF},
			}.Add(gtx.Ops)
			defer clip.Rect(dr).Push(gtx.Ops).Pop()
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{
				Size: gtx.Constraints.Max,
			}
		},
	}

	return material.List(th, list).Layout(gtx, len(widgets), func(gtx C, i int) D {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})
}
