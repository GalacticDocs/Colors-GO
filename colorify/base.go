package colorify

import "github.com/GalacticDocs/Colors-GO/checkers"

type IBaseColors struct {
	Black   func(string) string
	Red     func(string) string
	Green   func(string) string
	Yellow  func(string) string
	Blue    func(string) string
	Magenta func(string) string
	Cyan    func(string) string
	White   func(string) string
	Gray    func(string) string
}

func Base() *IBaseColors {
	Colors := &IBaseColors{
		Black:   func(str string) string { return checkers.Initiate(30, 39, "", str) },
		Red:     func(str string) string { return checkers.Initiate(31, 39, "", str) },
		Green:   func(str string) string { return checkers.Initiate(32, 39, "", str) },
		Yellow:  func(str string) string { return checkers.Initiate(33, 39, "", str) },
		Blue:    func(str string) string { return checkers.Initiate(34, 39, "", str) },
		Magenta: func(str string) string { return checkers.Initiate(35, 39, "", str) },
		Cyan:    func(str string) string { return checkers.Initiate(36, 39, "", str) },
		White:   func(str string) string { return checkers.Initiate(37, 39, "", str) },
		Gray:    func(str string) string { return checkers.Initiate(90, 39, "", str) },
	}

	return Colors
}
