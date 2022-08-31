package core

import (
	"bytes"
	"github.com/fexli/logger"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
)

type Draw struct {
	XY      [4]int
	Text    []Text
	BgColor color.RGBA

	DefaultFont string
	defaultSize float64
	canvas      *image.RGBA
	ctx         *freetype.Context
	cache       map[string]*truetype.Font
}
type Text struct {
	String  string
	Color   color.RGBA
	Size    float64
	NewLine bool
}

func (d *Draw) Create() []byte {
	canvas := image.NewRGBA(image.Rect(d.XY[0], d.XY[1], d.XY[2], d.XY[3]))
	draw.Draw(canvas, canvas.Bounds(), image.NewUniform(d.BgColor), image.Pt(0, 0), draw.Src)
	ctx := freetype.NewContext()
	ctx.SetDPI(144)
	ctx.SetFont(d.init())
	ctx.SetClip(canvas.Bounds())
	ctx.SetDst(canvas)
	ctx.SetHinting(font.HintingNone)
	ctx.SetSrc(image.NewUniform(color.RGBA{R: 168, G: 69, B: 66, A: 1}))
	d.canvas, d.ctx = canvas, ctx
	d.print()
	return d.output()
}
func (d *Draw) print() {
	d.defaultSize = 24
	pt := freetype.Pt(int(d.defaultSize), int(d.ctx.PointToFixed(d.defaultSize)>>6))
	for _, str := range d.Text {
		d.ctx.SetFontSize(str.Size)
		_, _ = d.ctx.DrawString(str.String, pt)
		pt.Y += d.ctx.PointToFixed(str.Size * 1.5)
	}
}
func (d *Draw) output() []byte {
	b := new(bytes.Buffer)
	if err := png.Encode(b, d.canvas); err != nil {
		logger.RootLogger.Error(logger.WithContent("图片绘制失败", err))
		return nil
	}
	return b.Bytes()
}
