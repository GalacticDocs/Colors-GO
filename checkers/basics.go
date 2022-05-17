package checkers

import (
	"fmt"
	"strings"

	conf "github.com/GalacticDocs/Colors-GO/config"
)

type ColorStringify struct {
	stringify func(string) string
}

func Start() {
	conf.Inititation()

	// argv := os.Args

	// isDisabled := StringToBool(conf.Get("NO_COLOR")) || StringItemExists(argv, "--no-color") || StringItemExists(argv, "--no-color")
	// isForced := StringToBool(conf.Get("FORCE_COLOR")) || StringItemExists(argv, "--force-color") || StringItemExists(argv, "--color")
	// isWindows := runtime.GOOS == "windows"

	// isCompatibleTerminal := isatty.IsTerminal(os.Stdout.Fd()) && conf.Get("TERM") != "dumb"

	// IsColorSupported = !isDisabled && (isForced || isWindows || isCompatibleTerminal)

}

func replaceClose(index int, str string, close, replace string) string {
	var (
		head string = str[0:index] + replace
		tail string = str[index:len(close)]
		next int    = strings.Index(tail, close)
	)

	if next < 0 {
		return head + tail
	} else {
		return head + replaceClose(next, tail, close, replace)
	}
}

func clearBleed(index int, str string, open string, close string, replace string) string {
	if index < 0 {
		return open + str + close
	} else {
		return open + replaceClose(index, str, close, replace) + close
	}
}

func filterEmpty(open string, close string, replace string) *ColorStringify {
	var at int = len(open) + 1

	return &ColorStringify{
		stringify: func(str string) string {
			if len(str) > 0 {
				return clearBleed(
					strings.Index(""+str, close[at:len(close)]),
					str,
					open,
					close,
					replace,
				)
			} else {
				return ""
			}
		},
	}
}

func Initiate(open int, close int, replace string, msg string) string {
	var (
		open_value  = fmt.Sprint(open)
		close_value = fmt.Sprint(close)
	)

	return filterEmpty("\x1b["+open_value+"m", "\x1b["+close_value+"m", replace).stringify(msg)
}

// type ColorCreationData struct{
// 	useColor bool;
// }

// func CreateColors() *IColors {

// }
