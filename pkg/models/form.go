package models

import (
	"errors"
	"fmt"
	"strings"
)

type Form struct {
	Name    string
	Message string
	Fields  []Field
	SendTo  []SendToAdapter
}

type Field struct {
	Name     string
	Required bool
	Value    string `yaml:"default"`
}

type SendToAdapter struct {
	Name     string
	Contacts []string
}

func (f *Form) Validate() error {
	if len(f.Name) <= 0 {
		return errors.New("field name is required")
	}

	for _, field := range f.Fields {
		if field.Required && len(field.Value) <= 0 {
			return fmt.Errorf("field %s is required", field.Name)
		}
	}

	return nil
}

func (f *Form) FormattedMessage() string {
	var fieldsf []string

	for _, field := range f.Fields {
		fieldsf = append(fieldsf, fmt.Sprintf("%s: %s", field.Name, field.Value))
	}

	strfields := strings.Join(fieldsf, "\n")
	msg := strings.ReplaceAll(f.Message, "{{form}}", strfields)

	return msg
}
