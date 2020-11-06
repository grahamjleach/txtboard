package txtboard

import (
	"strings"
)

type input struct {
	action string
	object string
	parameters []string
}

func (i input) exit() bool {
	return i.action == "exit"
}

func (i input) valid() bool {
	return i.action != "" && i.object != ""
}

func (i input) isMovement() bool {
	return i.action == "go"
}

func (i input) direction() direction {
	return directionFromString(i.object)
}

func clean(in []string) []string {
	var out []string

	for _, s := range in {
		if trimmed := strings.TrimSpace(s); trimmed != "" {
			out = append(out, trimmed)
		}
	}

	return out
}
