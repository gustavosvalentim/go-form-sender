package main

import (
	"errors"
	"log"
	"os"

	"github.com/gustavosvalentim/go-form-sender/pkg/models"
	"github.com/gustavosvalentim/go-form-sender/pkg/ports"
	"github.com/gustavosvalentim/go-form-sender/pkg/services"
)

func main() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	models.InitConfiguration(userDir + "/.formsenderrc")

	switch os.Args[1] {
	case "send":
		formName := os.Args[2]
		svc := services.NewSendService(new(ports.FileSystemStorage))
		svc.Send(formName)
	case "set":
		name, val := os.Args[2], os.Args[3]
		models.SetConfiguration(name, val)
		models.PersistConfiguration()
	default:
		log.Fatal(errors.New("not implemented"))
	}
}
