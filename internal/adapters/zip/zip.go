package zip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"teste/internal/adapters/zip/model"
	"teste/internal/ports"
)

type ZipAdapter struct{}

func NewZipAdapter() ports.ZipAdapter {
	return ZipAdapter{}
}

func (adapter ZipAdapter) GetAddressFromCorreiosByZipCode(cep string) (model.ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return model.ViaCEPResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.ViaCEPResponse{}, errors.New("falha ao consultar o CEP: resposta não OK")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ViaCEPResponse{}, err
	}

	var result model.ViaCEPResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return model.ViaCEPResponse{}, err
	}

	if result.Erro == "true" {
		return model.ViaCEPResponse{}, errors.New("CEP não encontrado")
	}

	return result, nil
}
