package main

import (
	"sandwich-dolar/dolar"
	"sandwich-dolar/menu"
	"sandwich-dolar/sandwich"
)

/* sandwich vale 5 dolares (Usar api para conseguir precio local)
Primer Menú (Solo la primera vez al ingresar al programa)
1 Argentina
2 Uruguay
3 Chile
4 Mexico
5 Bolivia

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

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func main() {
	// fmt.Println(uint(5.5))
	s := sandwich.NuevoSandwich(5.0)
	country := Must(menu.CountryMenu())

	valueDolar := Must(dolar.GetValue(country))
	currency := Must(dolar.GetCurrency(country))

	menu.MainMenu(s.Precio, valueDolar, currency)

	// Scan de nombre del cliente y catnidad total de sandwitches
	// imprimir total, solicitar importe con el que va a abonar y calcular e imprimir el vuelto

	// Confirmacion de nueva orden

}
