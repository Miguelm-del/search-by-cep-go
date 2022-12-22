package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		website := "http://viacep.com.br/ws/"
		url := website + cep + "/json"
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when making the request: %v\n", err)

		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error read response: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro parse response %v\n", err)
		}
		fmt.Println("CEP:    ", data.Cep)
		fmt.Println("Rua:    ", data.Logradouro)
		fmt.Println("Bairro: ", data.Bairro)
		fmt.Println("UF:     ", data.Uf)
		fmt.Println("Ibge:   ", data.Ibge)
		fmt.Println("Ddd:    ", data.Ddd)
		fmt.Println("Siafi:  ", data.Siafi)
	}

}
