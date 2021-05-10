package product

type Product struct {
	tableName struct{} `pg:"product"`
	ID        int      `json:"id" pg:"id,pk"`
	Name      string   `json:"name" pg:"name"`
	Price     float64  `json:"price" pg:"price"`
}

type ProductRepository interface {
	GetProducts() ([]*Product, error)
}
