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
/* type Pais struct {
	SubDomain string `json:"subDomain"`
	Libre     bool   `json:"libre"`
	CodigoISO string `json:"codigoISO"`
}
*/

var (
	errCountry   = errors.New("invalid country")
	errCodigoISO = errors.New("no 'codigoISO' for the country")
	errSubDomain = errors.New("no subdomain for the country")

	paises = map[string]map[string]interface{}{
		"Argentina": {
			"subDomain": "",
			"libre":     false,
			"codigoISO": "ARS",
		},
		"Chile": {
			"subDomain": "cl",
			"libre":     true,
			"codigoISO": "CLP",
		},
		"Mexico": {
			"subDomain": "mx",
			"libre":     true,
			"codigoISO": "MXN",
		},
		"Bolivia": {
			"subDomain": "bo",
			"libre":     false,
			"codigoISO": "BOP",
		},
		"Uruguay": {
			"subDomain": "uy",
			"libre":     true,
			"codigoISO": "UYU",
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

func getApiUrl(pais string) (string, error) {

	// Validamos si el país existe en el mapa
	if _, existe := paises[pais]; !existe {
		return "", fmt.Errorf("error: %w", errCountry)
	}

	subDomain, ok := paises[pais]["subDomain"]
	if !ok {
		return "", fmt.Errorf("error: %w", errSubDomain)
	}

	if subDomain == "" {
		return "https://dolarapi.com/v1/dolares/oficial", nil
	}

	if paises[pais]["libre"].(bool) {
		return fmt.Sprintf("https://%s.dolarapi.com/v1/cotizaciones/usd", paises[pais]["subDomain"]), nil
	}

	return fmt.Sprintf("https://%s.dolarapi.com/v1/dolares/oficial", paises[pais]["subDomain"]), nil
}

func GetValue(pais string) (float64, error) {
	var dolar Dolar
	if ok, err := dolarApiStatus(); !ok {
		return 0.0, err
	}

	url, err := getApiUrl(pais)

	if err != nil {
		return 0.0, err
	}

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

	currency, ok := paises[pais]["codigoISO"]

	if !ok {
		return "", fmt.Errorf("error: %w", errCodigoISO)
	}

	return currency.(string), nil
}
