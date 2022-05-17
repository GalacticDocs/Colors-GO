package colorify

import "github.com/GalacticDocs/Colors-GO/checkers"

type IBrightColors struct {
	BlackBright   func(string) string
	RedBright     func(string) string
	GreenBright   func(string) string
	YellowBright  func(string) string
	BlueBright    func(string) string
	MagentaBright func(string) string
	CyanBright    func(string) string
	WhiteBright   func(string) string
}

func BrightColors() *IBrightColors {
	Colors := &IBrightColors{
		BlackBright:   func(str string) string { return checkers.Initiate(90, 39, "\x1b[1m", str) },
		RedBright:     func(str string) string { return checkers.Initiate(91, 39, "\x1b[1m", str) },
		GreenBright:   func(str string) string { return checkers.Initiate(92, 39, "\x1b[1m", str) },
		YellowBright:  func(str string) string { return checkers.Initiate(93, 39, "\x1b[1m", str) },
		BlueBright:    func(str string) string { return checkers.Initiate(94, 39, "\x1b[1m", str) },
		MagentaBright: func(str string) string { return checkers.Initiate(95, 39, "\x1b[1m", str) },
		CyanBright:    func(str string) string { return checkers.Initiate(96, 39, "\x1b[1m", str) },
		WhiteBright:   func(str string) string { return checkers.Initiate(97, 39, "\x1b[1m", str) },
	}

	return Colors
}
