package ports

import "teste/internal/adapters/zip/model"

type ZipAdapter interface {
	GetAddressFromCorreiosByZipCode(cep string) (model.ViaCEPResponse, error)
}
