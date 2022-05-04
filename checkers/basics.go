package checkers

import (
	"errors"
	"fmt"
	"os"

	"github.com/GalacticDocs/unitylogger/tree/GOlang/config"
)

func start() {
	var (
		environment = checkers.GetEnvironment()
		argv        = os.Args
	)
}

func init(open int, close int, replace string) (string, error) {
	filter_empty("\x1b["+open+"m", "\x1b["+close+"m", replace)
}
