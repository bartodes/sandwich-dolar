package sandwich

type Sandwich struct {
	Precio Precio
}

type Precio float64

func NuevoSandwich(precio Precio) Sandwich {
	return Sandwich{
		Precio: precio,
	}
}
