package sandwich

type Sandwich struct {
	Precio float64
}

/* Retorna importe total de la compra de sandwiches segun la cantidad ingresada, estimado en dolares*/
func (s *Sandwich) GetTotalValue(cant uint) float64 {
	return s.Precio * float64(cant)
}

/* Genera un nuevo Sandwich */
func NuevoSandwich(precio float64) Sandwich {
	return Sandwich{
		Precio: precio,
	}
}
