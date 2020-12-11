package builtins

import (
	"github.com/marcusolsson/tui-go"
	"io/ioutil"
	"strconv"
)

//Ls lists files and directories in  an organized horizontal box
func Ls(dir string, root *tui.Box) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var size string
	var fileAmount int
	var dirs int

	fileTable := tui.NewTable(0, 0)
	fileTable.AppendRow(tui.NewLabel("Files"), tui.NewLabel("Size (bytes)"))
	fileTable.SetBorder(true)

	dirTable := tui.NewTable(0, 0)
	dirTable.AppendRow(tui.NewLabel("Directories"))
	dirTable.SetBorder(true)

	for _, file := range files {
		if file.IsDir() {
			dirs++

			dirTable.AppendRow(tui.NewLabel(file.Name()))
		} else {
			fileAmount++

			size = strconv.Itoa(int(file.Size()))
			fileTable.AppendRow(tui.NewLabel(file.Name()), tui.NewLabel(size))
		}
	}

	//yeah...I wish there was a better way to do this
	dirContent := tui.NewHBox()
	if fileAmount > 0 {
		dirContent.Append(fileTable)
	}

	if dirs > 0 {
		dirContent.Append(dirTable)
	}

	root.Append(dirContent)

	return nil
}
