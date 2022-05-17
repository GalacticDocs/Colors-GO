package colorify

import "github.com/GalacticDocs/Colors-GO/checkers"

type IFormatColors struct {
	Reset         func(string) string
	Bold          func(string) string
	Dim           func(string) string
	Italic        func(string) string
	Underline     func(string) string
	Inverse       func(string) string
	Hidden        func(string) string
	Strikethrough func(string) string
}

func Format() *IFormatColors {
	Colors := &IFormatColors{
		Reset:     func(str string) string { return checkers.Initiate(0, 0, "", str) },
		Bold:      func(str string) string { return checkers.Initiate(1, 22, "\x1b[22m\x1b[1m", str) },
		Dim:       func(str string) string { return checkers.Initiate(2, 22, "\x1b[22m\x1b[2m", str) },
		Italic:    func(str string) string { return checkers.Initiate(3, 23, "\x1b[23m\x1b[3m", str) },
		Underline: func(str string) string { return checkers.Initiate(4, 24, "\x1b[24m\x1b[4m", str) },
		Inverse:   func(str string) string { return checkers.Initiate(7, 27, "\x1b[27m\x1b[7m", str) },
		Hidden:    func(str string) string { return checkers.Initiate(8, 28, "\x1b[28m\x1b[8m", str) },
		Strikethrough: func(str string) string {
			return checkers.Initiate(9, 29, "\x1b[29m\x1b[9m", str)
		},
	}

	return Colors
}
