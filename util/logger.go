package util

import (
	"fmt"
	"os"

	"github.com/agtorre/gocolorize"
)

func ColorInfo(message string) string {
	info := gocolorize.Colorize{Fg: gocolorize.Magenta}
	return info.Paint(message)
}

func ColorWarning(message string) string {
	warning := gocolorize.Colorize{Fg: gocolorize.Yellow}
	return warning.Paint(message)
}

func ColorError(message string) string {
	err := gocolorize.Colorize{Fg: gocolorize.Red}
	err.ToggleBold()
	return err.Paint(message)
}

func ColorSuccess(message string) string {
	success := gocolorize.Colorize{Fg: gocolorize.Green}
	return success.Paint(message)
}

func LogInfo(message string) {
	fmt.Println(ColorInfo(message))
}

func LogWarning(message string) {
	fmt.Println(ColorWarning(message))
}

func LogError(message string) {
	fmt.Println(ColorError(message))
}

func LogSuccess(message string) {
	fmt.Println(ColorSuccess(message))
}

func Die(err error) {
	if err != nil {
		LogError(err.Error())
	}
	os.Exit(1)
}
