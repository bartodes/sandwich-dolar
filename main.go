package main

import (
	"fmt"
	"sandwich-dolar/dolar"
)

/* sandwich vale 5 dolares (Usar api para conseguir precio local)
Primer Menú (Solo la primera vez al ingresar al programa)
1 Argentina
2 Uruguay
3 Chile
4 Mexico
5 Venezuela

Segundo Menú (Repetitivo)
Mostrar Precio unitario(local) del Sandwich
1 Nueva orden
2 Salir

Nueva compra debe pedir:
Nombre
Cantidad Sandwich

Total (Guardarlo en Dolares)

Pedir dinero
Mostrar el vuelto
Desea continuar? y --> insert en una mysql con xampp "o" guardar en un txt

Todo plus que crean que es necesario es bienvenido */

type Status struct {
	Estado string `json:"estado"`
	Random int    `json:"aleatorio"`
}

func main() {
	// menu.ShowCountryMenu()
	// menu.ShowMainMenu("args", 1.0)
	// Scan de nombre del cliente y catnidad total de sandwitches
	// imprimir total, solicitar importe con el que va a abonar y calcular e imprimir el vuelto

	// Confirmacion de nueva orden

	dolar, err := dolar.GetValue("Bolivia")
	if err != nil {
		panic(err)
	}
	fmt.Println(dolar)
}
