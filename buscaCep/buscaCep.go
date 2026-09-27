package buscacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Ibge         struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"ibge"`
}

func (v ViaCep) String() string {
	return fmt.Sprintf(
		"CEP:          %s\n"+
			"Logradouro:   %s\n"+
			"Complemento:  %s\n"+
			"Unidade:      %s\n"+
			"Bairro:       %s\n"+
			"Localidade:   %s\n"+
			"UF:           %s\n"+
			"Estado:       %s\n"+
			"Região:       %s\n"+
			"IBGE:         %s\n"+
			"GIA:          %s\n"+
			"DDD:          %s\n"+
			"SIAFI:        %s",
		v.Cep,
		v.Logradouro,
		v.Complemento,
		v.Unidade,
		v.Bairro,
		v.Localidade,
		v.Uf,
		v.Estado,
		v.Regiao,
		v.Ibge,
		v.Gia,
		v.Ddd,
		v.Siafi,
	)
}

func (b BrasilApi) String() string {
	return fmt.Sprintf(
		"CEP:         %s\n"+
			"State:       %s\n"+
			"City:        %s\n"+
			"Neighborhood:%s\n"+
			"Street:      %s\n"+
			"Service:     %s\n"+
			"IBGE City:   %s\n"+
			"IBGE State:  %s",
		b.Cep,
		b.State,
		b.City,
		b.Neighborhood,
		b.Street,
		b.Service,
		b.Ibge.City,
		b.Ibge.State,
	)
}

func BuscaViaCep(cep string) (*ViaCep, error) {
	if cep == "" {
		return nil, fmt.Errorf("invalid cep input")
	}

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	body, err := fetch(url)
	if err != nil {
		return nil, err
	}

	var viacep ViaCep
	json.Unmarshal(body, &viacep)

	return &viacep, nil
}

func BuscaBrasilApi(cep string) (*BrasilApi, error) {
	if cep == "" {
		return nil, fmt.Errorf("invalid cep input")
	}

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	body, err := fetch(url)
	if err != nil {
		return nil, err
	}

	var brasilapi BrasilApi
	json.Unmarshal(body, &brasilapi)

	return &brasilapi, nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failled to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}
