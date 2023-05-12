package products

type Product struct {
	Id    int    `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
