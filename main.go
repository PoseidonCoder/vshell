package main

import (
	"log"
	"os/exec"
	"strings"
	"vshell/src/builtins"

	"github.com/marcusolsson/tui-go"
)

var root *tui.Box
var output *tui.Box

func main() {
	root = tui.NewVBox()

	theme := tui.NewTheme()
	theme.SetStyle("normal", tui.Style{Bg: tui.ColorBlack, Fg: tui.ColorWhite})

	theme.SetStyle("label.fatal", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})
	theme.SetStyle("label.info", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorMagenta})

	//Output area
	output = tui.NewVBox()
	output.SetTitle("Output")
	output.SetBorder(true)
	output.SetSizePolicy(tui.Expanding, tui.Expanding)
	root.Append(output)

	scroll := tui.NewScrollArea(output)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)
	input.OnSubmit(func(entry *tui.Entry) {
		if err := execInput(entry.Text()); err != nil {
			errorLabel := tui.NewLabel(err.Error())
			errorLabel.SetStyleName("fatal")

			output.Append(errorLabel)
		}

		input.SetText("")
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	root.Append(inputBox)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Up", func() { scroll.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() { scroll.Scroll(0, 1) })
	ui.SetKeybinding("Left", func() { scroll.Scroll(-1, 0) })
	ui.SetKeybinding("Right", func() { scroll.Scroll(1, 0) })
	ui.SetTheme(theme)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func execInput(input string) error {
	// Clean input of unnecessary characters
	input = strings.TrimSuffix(input, "\r")

	args := strings.Split(input, " ")
	primary := args[0]

	//TODO: add command parser
	err, found := builtins.RunCustom(primary, args[1:], output)
	if found {
		return err
	}

	out, err := exec.Command(primary, args[1:]...).Output()

	info := tui.NewLabel(string(out))
	info.SetStyleName("info")
	output.Append(info)

	return err
}
