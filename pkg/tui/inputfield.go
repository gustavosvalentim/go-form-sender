package tui

import (
	"fmt"

	"github.com/gustavosvalentim/go-form-sender/pkg/models"
)

func InputField(fd models.Field) string {
	var fieldMeta string
	var ret string

	if fd.Value != "" {
		fieldMeta += fmt.Sprintf(" (default: %s)", fd.Value)
	}

	if fd.Required {
		fieldMeta += "*"
	}

	fmt.Printf("%s%s: ", fd.Name, fieldMeta)
	fmt.Scanln(&ret)

	return ret
}
