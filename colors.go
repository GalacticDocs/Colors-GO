package colors

import (
	"github.com/GalacticDocs/Colors-GO/colorify"
)

func Reset(str string) string {
	return colorify.Format().Reset(str)
}
func Bold(str string) string {
	return colorify.Format().Bold(str)
}
func Dim(str string) string {
	return colorify.Format().Dim(str)
}
func Italic(str string) string {
	return colorify.Format().Italic(str)
}
func Underline(str string) string {
	return colorify.Format().Underline(str)
}
func Inverse(str string) string {
	return colorify.Format().Inverse(str)
}
func Hidden(str string) string {
	return colorify.Format().Hidden(str)
}
func StrikeThrough(str string) string {
	return colorify.Format().Strikethrough(str)
}

func Black(str string) string {
	return colorify.Base().Black(str)
}
func Red(str string) string {
	return colorify.Base().Red(str)
}
func Green(str string) string {
	return colorify.Base().Green(str)
}
func Yellow(str string) string {
	return colorify.Base().Yellow(str)
}
func Blue(str string) string {
	return colorify.Base().Blue(str)
}
func Magenta(str string) string {
	return colorify.Base().Magenta(str)
}
func Cyan(str string) string {
	return colorify.Base().Cyan(str)
}
func White(str string) string {
	return colorify.Base().White(str)
}
func Gray(str string) string {
	return colorify.Base().Gray(str)
}

func BlackBright(str string) string {
	return colorify.BrightColors().BlackBright(str)
}
func RedBright(str string) string {
	return colorify.BrightColors().RedBright(str)
}
func GreenBright(str string) string {
	return colorify.BrightColors().GreenBright(str)
}
func YellowBright(str string) string {
	return colorify.BrightColors().YellowBright(str)
}
func BlueBright(str string) string {
	return colorify.BrightColors().BlueBright(str)
}
func MagentaBright(str string) string {
	return colorify.BrightColors().MagentaBright(str)
}
func CyanBright(str string) string {
	return colorify.BrightColors().CyanBright(str)
}
func WhiteBright(str string) string {
	return colorify.BrightColors().WhiteBright(str)
}

func BgBlack(str string) string {
	return colorify.BGColors().BgBlack(str)
}
func BgRed(str string) string {
	return colorify.BGColors().BgRed(str)
}
func BgGreen(str string) string {
	return colorify.BGColors().BgGreen(str)
}
func BgYellow(str string) string {
	return colorify.BGColors().BgYellow(str)
}
func BgBlue(str string) string {
	return colorify.BGColors().BgBlue(str)
}
func BgMagenta(str string) string {
	return colorify.BGColors().BgMagenta(str)
}
func BgCyan(str string) string {
	return colorify.BGColors().BgCyan(str)
}
func BgWhite(str string) string {
	return colorify.BGColors().BgWhite(str)
}

func BgBlackBright(str string) string {
	return colorify.BrightBGColors().BgBlackBright(str)
}
func BgRedBright(str string) string {
	return colorify.BrightBGColors().BgRedBright(str)
}
func BgGreenBright(str string) string {
	return colorify.BrightBGColors().BgGreenBright(str)
}
func BgYellowBright(str string) string {
	return colorify.BrightBGColors().BgYellowBright(str)
}
func BgBlueBright(str string) string {
	return colorify.BrightBGColors().BgBlueBright(str)
}
func BgMagentaBright(str string) string {
	return colorify.BrightBGColors().BgMagentaBright(str)
}
func BgCyanBright(str string) string {
	return colorify.BrightBGColors().BgCyanBright(str)
}
func BgWhiteBright(str string) string {
	return colorify.BrightBGColors().BgWhiteBright(str)
}
