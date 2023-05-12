package products

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) GetAllProducts() ([]Product, error) {
	products, err := usecase.Repo.GetAllProducts()
	return products, err
}

func (usecase UseCase) GetProductById(id int) (*Product, error) {
	product, err := usecase.Repo.GetProductById(id)
	return product, err
}

func (usecase UseCase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}

func (usecase UseCase) UpdateProductById(id int, product *Product) error {
	err := usecase.Repo.UpdateProductById(id, product)
	return err
}

func (usecase UseCase) DeleteProductById(id int) error {
	err := usecase.Repo.DeleteProductById(id)
	return err
}