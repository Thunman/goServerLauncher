package main

import (
	"fmt"
	"os"

	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	fmt.Println("Starting program")
	serverLauncher := app.New()

	launcherWindow := serverLauncher.NewWindow("Node.js Server Controller")

	startButton := widget.NewButtonWithIcon("Start Server", theme.MediaPlayIcon(), func() {
		startServer()
	})

	stopButton := widget.NewButtonWithIcon("Stop Server", theme.MediaStopIcon(), func() {
		stopServer()
	})

	buttonsContainer := container.NewVBox(
		container.New(layout.NewCenterLayout(), container.NewPadded(layout.NewSpacer()), startButton),
		container.New(layout.NewCenterLayout(), container.NewPadded(layout.NewSpacer()), stopButton),
	)
	content := container.NewBorder(nil, nil, nil, nil,
		container.NewVBox(
			buttonsContainer,
		),
	)
	launcherWindow.SetContent(content)
	launcherWindow.ShowAndRun()
}

var cmd *exec.Cmd

func startServer() {
	nodePath := os.Getenv("NODE_PATH")
	if nodePath == "" {
		panic("NODE_PATH not set")
	}
	cmd = exec.Command("node", nodePath, "start")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func stopServer() {
	if cmd != nil && cmd.Process != nil {
		nodePath := os.Getenv("NODE_PATH")
		if nodePath == "" {
			panic("NODE_PATH not set")
		}
		cmd = exec.Command("node", nodePath, "stop")
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}
}
