package colorify

import "github.com/GalacticDocs/Colors-GO/checkers"

type IBGColors struct {
	BgBlack   func(string) string
	BgRed     func(string) string
	BgGreen   func(string) string
	BgYellow  func(string) string
	BgBlue    func(string) string
	BgMagenta func(string) string
	BgCyan    func(string) string
	BgWhite   func(string) string
}

func BGColors() *IBGColors {
	Colors := &IBGColors{
		BgBlack:   func(str string) string { return checkers.Initiate(40, 49, "", str) },
		BgRed:     func(str string) string { return checkers.Initiate(41, 49, "", str) },
		BgGreen:   func(str string) string { return checkers.Initiate(42, 49, "", str) },
		BgYellow:  func(str string) string { return checkers.Initiate(43, 49, "", str) },
		BgBlue:    func(str string) string { return checkers.Initiate(44, 49, "", str) },
		BgMagenta: func(str string) string { return checkers.Initiate(45, 49, "", str) },
		BgCyan:    func(str string) string { return checkers.Initiate(46, 49, "", str) },
		BgWhite:   func(str string) string { return checkers.Initiate(47, 49, "", str) },
	}

	return Colors
}
