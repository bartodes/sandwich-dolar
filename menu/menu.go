package menu

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	errOpcion = errors.New("la opción ingresada no esta permitida")
	opts      = map[string]interface{}{
		"1": "Argentina",
		"2": "Uruguay",
		"3": "Chile",
		"4": "Mexico",
		"5": "Bolivia",
	}
)

/* Utiliza fmt.Scan, hace el error handling y retorna un string con lo que leyo como input. */
func mustScan() string {
	var x string
	_, err := fmt.Scan(&x)
	if err != nil {
		panic(err)
	}

	return x
}

// funcion para hacer conversiones de string a distintos tipos segun tipo de parametros
func mustConv[T uint64 | float64](s string) T {
	// Intentamos convertir a uint64 primero
	if _, ok := any(new(T)).(*uint64); ok {
		// Si T es uint64, intentamos parsear como uint64
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			return any(v).(T)
		}
	}

	// Si falló la conversión como uint64, intentamos convertirlo a float64
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return any(v).(T)
	}

	// Si ambas conversiones fallan, hacemos panic
	panic(fmt.Sprintf("failed to convert '%s' to either uint64 or float64", s))
}

/*
Muestra un menú con las opciones de los diferentes paises disponibles.
Retorna el valor que corresponde a determinado pais
*/
func CountryMenu() (string, error) {
	fmt.Print("De que país esta realizando la orden?\n1 Argentina\n2 Uruguay\n3 Chile\n4 Mexico\n5 Bolivia\n? ")

	opcionPais := mustScan()

	if _, ok := opts[opcionPais]; !ok {
		return "", fmt.Errorf("error: %w", errOpcion)
	}

	return opts[opcionPais].(string), nil
}

func MainMenu(precioSandwitch, precioDolar float64, moneda string) (string, uint64, float64, error) {
	var opcion, cant uint64
	var name string
	var montoPago, totalDolares float64

	precioSandwitchPesos := precioSandwitch * precioDolar

	fmt.Printf("\nCada Sandwich está: $%2.f %s (%2.f USD)\n", precioSandwitchPesos, moneda, precioSandwitch)
	fmt.Print("Seleccionar una opción:\n1 Nueva Compra\n2 Salir\n? ")

	opcion = mustConv[uint64](mustScan())

	if opcion > 1 {
		panic(errOpcion)
	}

	if opcion == 1 {
		for {
			fmt.Print("\nCual es su nombre? ")
			name = mustScan()

			if name != "" {
				break
			}

			fmt.Println("No puede no ingresar un nombre!")
		}

		for {
			fmt.Printf("%s, ¿Cuantos sandwichs quiere? ", name)
			cant = mustConv[uint64](mustScan())

			if cant >= 1 {
				break
			}

			fmt.Println("No puede ingresar menos de 1!")
		}

		fmt.Println()

		totalDolares := precioSandwitch * float64(cant)
		totalPesos := precioSandwitchPesos * float64(cant)

		fmt.Printf("El precio total es: %2.f\n", totalPesos)

		for {
			fmt.Print("Con que monto vas a abonar? ")
			montoPago = mustConv[float64](mustScan())

			if montoPago >= totalPesos {
				break
			}

			fmt.Println("No puede ingresar un monto menor al importe de la compra!")
		}

		vuelto := montoPago - totalPesos

		if vuelto > 0 {
			fmt.Printf("\nSu vuelto: %2.f %s\n", vuelto, moneda)
		}

		return name, cant, totalDolares, nil
	}

	fmt.Println("Goodbye!")
	os.Exit(0)

	return name, cant, totalDolares, nil
}
