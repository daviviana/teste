package adapters

import "teste/internal/ports"

type EmailAdapter struct{}

func NewEmailAdapter() ports.EmailAdapter {
	return &EmailAdapter{}
}

func (adapter EmailAdapter) SendEmail() {}
