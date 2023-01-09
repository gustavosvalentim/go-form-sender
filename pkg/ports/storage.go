package ports

import (
	"fmt"
	"log"
	"os"

	"github.com/gustavosvalentim/go-form-sender/pkg/models"
	"gopkg.in/yaml.v2"
)

type StoragePort interface {
	GetFormTemplateFromName(name string) models.Form
}

type FileSystemStorage struct{}

func (f *FileSystemStorage) GetFormTemplateFromName(name string) models.Form {
	path := fmt.Sprintf("%s/%s", models.GetConfiguration().TemplatesDir, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	form := models.Form{}
	err = yaml.Unmarshal(b, &form)
	if err != nil {
		log.Fatal(err)
	}

	return form
}
