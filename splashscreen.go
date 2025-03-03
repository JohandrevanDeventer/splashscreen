package splashscreen

import (
	"fmt"

	"github.com/johandrevandeventer/textutils"
)

func PrintSplashScreen() {

	fmt.Print(textutils.ColorText(textutils.Blue, `
██████  ██    ██ ██████  ██  ██████  ██████  ███    ██     ██████  ███    ███ ███████ 
██   ██ ██    ██ ██   ██ ██ ██      ██    ██ ████   ██     ██   ██ ████  ████ ██      
██████  ██    ██ ██████  ██ ██      ██    ██ ██ ██  ██     ██████  ██ ████ ██ ███████ 
██   ██ ██    ██ ██   ██ ██ ██      ██    ██ ██  ██ ██     ██   ██ ██  ██  ██      ██ 
██   ██  ██████  ██████  ██  ██████  ██████  ██   ████     ██████  ██      ██ ███████`))

	fmt.Println()
	fmt.Println()
}
