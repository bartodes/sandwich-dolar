package main

import (
	"fmt"
	"sandwich-dolar/dolar"
	"sandwich-dolar/menu"
	"sandwich-dolar/sandwich"
	"sandwich-dolar/store"
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
Insertar en DB
Volver a mostrar menu principal */

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func main() {
	db := store.OpenDB("mypassword", "sandwich-dolar")
	defer db.Close()

	store.TestDb(db)

	s := sandwich.NuevoSandwich(5.0)

	country := Must(menu.CountryMenu())
	valueDolar := Must(dolar.GetValue(country))
	currency := Must(dolar.GetCurrency(country))

	for {
		orden := Must(menu.MainMenu(s.Precio, valueDolar, currency))
		id := store.InstertVenta(db, orden)
		fmt.Printf("Orden guardada en la base de datos, con id: %d\n", id)
	}
}
