package colorify

import "github.com/GalacticDocs/Colors-GO/checkers"

type IBrightBGColors struct {
	BgBlackBright   func(string) string
	BgRedBright     func(string) string
	BgGreenBright   func(string) string
	BgYellowBright  func(string) string
	BgBlueBright    func(string) string
	BgMagentaBright func(string) string
	BgCyanBright    func(string) string
	BgWhiteBright   func(string) string
}

func BrightBGColors() *IBrightBGColors {
	Colors := &IBrightBGColors{
		BgBlackBright:   func(str string) string { return checkers.Initiate(100, 49, "", str) },
		BgRedBright:     func(str string) string { return checkers.Initiate(101, 49, "", str) },
		BgGreenBright:   func(str string) string { return checkers.Initiate(102, 49, "", str) },
		BgYellowBright:  func(str string) string { return checkers.Initiate(103, 49, "", str) },
		BgBlueBright:    func(str string) string { return checkers.Initiate(104, 49, "", str) },
		BgMagentaBright: func(str string) string { return checkers.Initiate(105, 49, "", str) },
		BgCyanBright:    func(str string) string { return checkers.Initiate(106, 49, "", str) },
		BgWhiteBright:   func(str string) string { return checkers.Initiate(107, 49, "", str) },
	}

	return Colors
}
