package services

import (
	"log"

	"github.com/gustavosvalentim/go-form-sender/pkg/ports"
	"github.com/gustavosvalentim/go-form-sender/pkg/tui"
)

type SendService struct {
	storage ports.StoragePort
}

func NewSendService(s ports.StoragePort) *SendService {
	return &SendService{s}
}

func (s *SendService) Send(name string) {
	form := s.storage.GetFormTemplateFromName(name)

	for i := 0; i < len(form.Fields); i++ {
		fd := &form.Fields[i]
		if val := tui.InputField(*fd); val != "" {
			fd.Value = val
		}
	}

	err := form.Validate()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range form.SendTo {
		adapter, err := ports.GetSenderAdapterFromName(s.Name)
		if err != nil {
			log.Fatal(err)
		}

		adapter.Send(form.FormattedMessage(), s.Contacts)
	}
}
