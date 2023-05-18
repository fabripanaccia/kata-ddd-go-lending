package domain

type Ddd interface {
	ValePorUnMetodo() string
}

type ddd struct{}

func (d *ddd) ValePorUnMetodo() string {
	return "Vale por una respuesta"
}
