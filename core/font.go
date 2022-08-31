package core

import (
	"github.com/fexli/logger"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
)

func (d *Draw) init() *truetype.Font {
	if d.DefaultFont == "" {
		d.DefaultFont = "UbuntuMono-R"
	}
	d.cache = make(map[string]*truetype.Font)
	return d.loadFont(d.DefaultFont)
}
func (d *Draw) loadFont(fontName string) *truetype.Font {
	if d.cache[fontName] != nil {
		return d.cache[fontName]
	}
	file := "assets/fonts/" + fontName + ".ttf"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("加载字体失败", err))
		return nil
	}
	f, err := freetype.ParseFont(b)
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("解析字体失败", err))
		return nil
	}
	d.cache[fontName] = f
	return f
}
