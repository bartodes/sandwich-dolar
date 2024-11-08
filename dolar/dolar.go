package dolar

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Status struct {
	estado string `json:"estado"`
}

type Dolar struct {
	compra float64 `json:"compra"`
	venta  float64 `json:"venta"`
}

/* Retorna el valor al que cotiza el dolar en la moneda del país ingresado */
func DolarApiStatus() bool {
	var estado Status

	resp, err := http.Get("https://dolarapi.com/v1/estado")

	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if err := json.Unmarshal(body, &estado); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func GetDolarValue(pais string) /* (, error) */ {
	var paises = [...]string{"Argentina", "Chile", "Mexico", "Bolivia", "Uruguay"}
	// if pais not in paises return // Generar un nuevo error

	// for que recorre el array para chequear que pais sea uno de los valores admitidos
	// y determinar la url de la api
	for _, v := range paises {

		// return
	}

	// traer toda la info de api dolar

	// substraer tipo de cambio segun pais

	// retornar tipo de cambio
	return 0.0, nil
}
