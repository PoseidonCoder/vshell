package builtins

import (
	"github.com/marcusolsson/tui-go"
	"os"

	"vshell/src/util"
)

func RunCustom(primary string, args []string, root *tui.Box) (err error, found bool) {
	found = true

	//TODO: add clear command that clears the "output" box
	switch primary {
	case "cd":
		if len(args) > 0 {
			err = os.Chdir(args[0])
		} else {
			err = util.ErrNoPath
			return
		}
	case "ls":
		if len(args) > 0 {
			err = Ls(args[0], root)
		} else {
			err = util.ErrNoPath
		}
	case "mkdir":
		for _, arg := range args {
			if err = os.Mkdir(arg, os.ModeDir); err != nil {
				return
			}
		}
	default:
		found = false
	}

	return
}
