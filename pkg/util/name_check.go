package util

import "strings"

// FixControllerName to fix user input controller name
func FixControllerName(controller string) string{
	currentController := strings.Title(controller)

	if strings.Contains(controller, "Controller") {
		currentController = strings.Trim(controller, "Controller")
	} else if strings.Contains(controller, "controller") {
		currentController = strings.Trim(controller, "controller")
	}
	return currentController
}
