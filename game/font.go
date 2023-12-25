package game

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/phuslu/log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusBigFont font.Face
)

func init() {
	// initialize font
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal().Err(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    100,
		DPI:     100,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal().Err(err)
	}
}
