package ports

import "fmt"

type SenderPort interface {
	Send(message string, to []string) error
}

type DebugSenderAdapter struct{}

func (s *DebugSenderAdapter) Send(message string, to []string) error {
	fmt.Print(message)
	return nil
}

func GetSenderAdapterFromName(name string) (SenderPort, error) {
	switch name {
	case "debug":
		return &DebugSenderAdapter{}, nil
	default:
		return nil, fmt.Errorf("sender adapter name %s not found", name)
	}
}
