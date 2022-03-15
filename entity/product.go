package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock < 3 {
		status = "Stock Terbatas"
	} else {
		status = "Masih Banyak Kok"
	}
	return status
}