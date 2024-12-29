package dolar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Status struct {
	Estado string `json:"estado"`
}

type Dolar struct {
	Compra float64 `json:"compra"`
	Venta  float64 `json:"venta"`
}

// Implementar el uso del siguiente struct:
type Pais struct {
	Url       string `json:"url"`
	CodigoISO string `json:"codigoISO"`
}

var (
	errCountry = errors.New("invalid country")

	paises = map[string]Pais{
		"Argentina": {
			Url:       "https://dolarapi.com/v1/dolares/oficial",
			CodigoISO: "ARS",
		},
		"Chile": {
			Url:       "https://cl.dolarapi.com/v1/cotizaciones/usd",
			CodigoISO: "CLP",
		},
		"Mexico": {
			Url:       "https://mx.dolarapi.com/v1/cotizaciones/usd",
			CodigoISO: "MXN",
		},
		"Bolivia": {
			Url:       "https://bo.dolarapi.com/v1/dolares/oficial",
			CodigoISO: "BOP",
		},
		"Uruguay": {
			Url:       "https://uy.dolarapi.com/v1/cotizaciones/usd",
			CodigoISO: "UYU",
		},
	}
)

/* Generar funcion Must para reducir codigo {if err != nil} */

/* Retorna el valor al que cotiza el dolar en la moneda del país ingresado */
func dolarApiStatus() (bool, error) {
	var estado Status

	resp, err := http.Get("https://dolarapi.com/v1/estado")

	if err != nil {
		// log.Fatal(err)
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// log.Fatal(err)
		return false, err
	}

	if err := json.Unmarshal(body, &estado); err != nil {
		// log.Fatal(err)
		return false, err
	}

	return true, nil
}

func GetValue(pais string) (float64, error) {
	var dolar Dolar
	if ok, err := dolarApiStatus(); !ok {
		return 0.0, err
	}

	if _, existe := paises[pais]; !existe {
		return 0.0, fmt.Errorf("error: %w", errCountry)
	}
	url := paises[pais].Url

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return 0.0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return 0.0, err
	}

	if err := json.Unmarshal(body, &dolar); err != nil {
		log.Fatal(err)
		return 0.0, err
	}

	value := (dolar.Compra + dolar.Venta) / 2

	// traer toda la info de api dolar

	// substraer tipo de cambio segun pais

	// retornar tipo de cambio
	return float64(value), nil
}

func GetCurrency(pais string) (string, error) {

	if _, existe := paises[pais]; !existe {
		return "", fmt.Errorf("error: %w", errCountry)
	}

	return paises[pais].CodigoISO, nil
}
