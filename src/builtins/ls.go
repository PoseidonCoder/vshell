package builtins

import (
	"github.com/marcusolsson/tui-go"
	"io/ioutil"
	"strconv"
)

func Ls(dir string, root *tui.Box) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var size string
	fileTable := tui.NewTable(0, 0)
	fileTable.AppendRow(tui.NewLabel("Files"), tui.NewLabel("Size (bytes)"))
	fileTable.SetBorder(true)

	dirTable := tui.NewTable(0, 0)
	dirTable.AppendRow(tui.NewLabel("Directories"))
	dirTable.SetBorder(true)

	for _, file := range files {
		if file.IsDir() {
			dirTable.AppendRow(tui.NewLabel(file.Name()))
		} else {
			size = strconv.Itoa(int(file.Size()))
			fileTable.AppendRow(tui.NewLabel(file.Name()), tui.NewLabel(size))
		}
	}

	root.Append(dirTable)
	root.Append(fileTable)
	return nil
}
