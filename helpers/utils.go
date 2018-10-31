package helpers

import (
	"strings"

	"github.com/bclicn/color"
)

// ColorizedDenied events
func ColorizedDenied(event string) string {
	if strings.Contains(event, "отклонен") || strings.Contains(event, "Запрет") {
		return color.Red(event)
	}

	return event
}
