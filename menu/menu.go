package menu

import (
	"fmt"
)

/*
Muestra un menú con las opciones de los diferentes paises disponibles.
Retorna el valor que corresponde a determinado pais
*/
func ShowCountryMenu() (int, error) {
	var opcion int

	fmt.Println("De que país esta realizando la orden?\n1 Argentina\n2 Uruguay\n3 Chile\n4 Mexico\n5 Venezuela\n?")

	_, err := fmt.Scan(&opcion)

	if err != nil {
		fmt.Println("La opción ingresada no es válida")
		return 0, err
	}

	// en este if statement generar un error especifico
	if opcion < 1 || opcion > 5 {
		fmt.Println("La opción ingresada no es válida")
		return 0, err
	}

	return opcion, nil
}

func ShowMainMenu(moneda string, precioDolar float64) (int, error) {
	var opcion int
	fmt.Printf("Cada Sandwich está: $%f %s (5 USD)\n", precioDolar*5, moneda)
	fmt.Println("Seleccionar una opción:\n1 Nueva Compra\n2 Salir")

	_, err := fmt.Scan(opcion)

	if err != nil {
		fmt.Println("La opción ingresada no es válida")
		return 0, err
	}

	// en este if statement generar un error especifico
	if opcion > 1 || opcion < 0 {
		fmt.Println("La opción ingresada no es válida")
		return 0, err
	}

	return opcion, nil
}
