package cliente

type Comprar interface {
	OrdenarSandwich() //Orden
}

type Cliente struct {
	Nombre Nombre
	Pais   Pais
}

type Nombre string

type Pais string

type Cantidad int

// func (c *Cliente) OrdenarSandwich(cantidad Cantidad, precio sandwich.Precio)/* Orden */ {}

func NuevoCliente(nombre Nombre, pais Pais) Cliente {
	return Cliente{
		Nombre: nombre,
		Pais:   pais,
	}
}
