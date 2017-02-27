package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/agtorre/gocolorize"
)

// ColorInfo returns a string with the info color
func ColorInfo(message string) string {
	info := gocolorize.Colorize{Fg: gocolorize.Magenta}
	return info.Paint(message)
}

// ColorWarning returns a string with the warning color
func ColorWarning(message string) string {
	warning := gocolorize.Colorize{Fg: gocolorize.Yellow}
	return warning.Paint(message)
}

// ColorError returns a string with the error color
func ColorError(message string) string {
	err := gocolorize.Colorize{Fg: gocolorize.Red}
	err.ToggleBold()
	return err.Paint(message)
}

// ColorSuccess returns a string with the success color
func ColorSuccess(message string) string {
	success := gocolorize.Colorize{Fg: gocolorize.Green}
	return success.Paint(message)
}

// LogInfo logs a string with the info color
func LogInfo(message string) {
	fmt.Println(ColorInfo(message))
}

// LogWarning logs a string with the warning color
func LogWarning(message string) {
	fmt.Println(ColorWarning(message))
}

// LogError logs a string with the error color
func LogError(message string) {
	fmt.Println(ColorError(message))
}

// LogSuccess logs a string with the success color
func LogSuccess(message string) {
	fmt.Println(ColorSuccess(message))
}

// JSONString turns an empty interface into a JSON string
func JSONString(value interface{}) (string, error) {
	json, err := json.MarshalIndent(value, "", "  ")
	if err == nil {
		return string(json), nil
	}
	return "", err
}

// Die logs an error and exits
func Die(err error) {
	if err != nil {
		LogError(err.Error())
	}
	os.Exit(1)
}
