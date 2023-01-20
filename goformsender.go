package main

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	Form struct {
		Name     string
		Subject  string
		Message  string
		Fields   []Field
		Contacts []string
	}

	Field struct {
		Name     string
		Required bool
		Value    string `yaml:"default"`
	}

	Mail struct {
		From    string
		To      []string
		Subject string
		Message string
	}
)

const (
	templates_folder   string = ".formtemplates"
	smtp_server_name   string = "smtp.gmail.com"
	smtp_server_port   int    = 587
	smtp_user_name     string = "email@test.com"
	smtp_user_password string = "*****"
)

// Form

func newFormFromTemplate(path string) (*Form, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	f := &Form{}
	if err := yaml.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}

	return f, nil
}

func (f *Form) validate() error {
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

func (f *Form) formattedMessage() string {
	var fieldsf []string

	for _, field := range f.Fields {
		fieldsf = append(fieldsf, fmt.Sprintf("%s: %s", field.Name, field.Value))
	}

	strfields := strings.Join(fieldsf, "\n")
	msg := strings.ReplaceAll(f.Message, "{{form}}", strfields)

	return msg
}

// Email

func (m *Mail) Send(host string, port int, username string, password string) error {
	addr := fmt.Sprintf("%s:%d", smtp_server_name, smtp_server_port)
	auth := smtp.PlainAuth("", username, password, host)
	emailMsg := "From: " + username + "\r\n" +
		"To: " + strings.Join(m.To, ",\r\n") + "\r\n" +
		"Subject:" + m.Subject + "\r\n" +
		"\r\n" + m.Message + "\r\n"

	if err := smtp.SendMail(addr, auth, username, m.To, []byte(emailMsg)); err != nil {
		return err
	}

	return nil
}

// Logging

func logInfo(msg string, params ...any) {
	logMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02"), fmt.Sprintf(msg, params...))
	fmt.Println(logMsg)
}

// CLI

func inputField(f *Field) {
	var fieldMeta string

	if f.Value != "" {
		fieldMeta += fmt.Sprintf("(default: %s)", f.Value)
	}

	if f.Required {
		fieldMeta += "*"
	}

	fmt.Printf("%s%s: ", f.Name, fieldMeta)
	fmt.Scanln(&f.Value)
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(errors.New("missing form name parameter"))
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	form, err := newFormFromTemplate(fmt.Sprintf("%s/%s/%s", homedir, templates_folder, os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(form.Fields); i++ {
		inputField(&form.Fields[i])
	}

	if err := form.validate(); err != nil {
		log.Fatal(err)
	}

	logInfo("form is valid")
	logInfo("sending mail")

	mail := &Mail{
		From:    smtp_user_name,
		To:      form.Contacts,
		Subject: form.Subject,
		Message: form.formattedMessage(),
	}
	err = mail.Send(smtp_server_name, smtp_server_port, smtp_user_name, smtp_user_password)
	if err != nil {
		log.Fatal(err)
	}
}
