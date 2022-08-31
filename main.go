package main

import (
	"encoding/base64"
	"fmt"
	"github.com/skadiD/text2png/core"
	"image/color"
)

func main() {
	img := core.Draw{
		XY: [4]int{0, 0, 400, 80},
		Text: []core.Text{
			{String: "中文 test", Color: color.RGBA{R: 255, G: 255, B: 0, A: 255}, Size: 24},
		},
		BgColor:     color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		DefaultFont: "MicrosoftYaHei",
	}
	fmt.Println("data:image/png;base64," + base64.StdEncoding.EncodeToString(img.Create()))
}
