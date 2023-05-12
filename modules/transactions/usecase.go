package transactions

import (
	"Pos-Design/modules/products"
	"fmt"
	"time"
)

type UseCase struct {
	Repo        Repository
	ProductRepo products.Repository
}

func (usecase UseCase) GetAllTransactions() ([]Transaction, error) {
	transactions, err := usecase.Repo.GetAllTransactions()
	return transactions, err
}

func (usecase UseCase) GetTransactionById(id int) (*Transaction, error) {
	transaction, err := usecase.Repo.GetTransactionById(id)
	return transaction, err
}

func (usecase UseCase) CreateTransaction(req *CreateTransactionRequest ) (*Transaction, error) {
	items := []TransactionItems{}
	totalPrice := 0

	// for _, i := range req.Items {
	// 	product, err := usecase.ProductRepo.GetProductById(int(i.ProductID))
	// 	if err != nil {
	// 		return nil, fmt.Errorf("product id not found %d", i.ProductID)
	// 	}

	// 	if i.Quantity > product.Stock {
	// 		return nil, fmt.Errorf("stock is not enough %s", product.Name)
	// 	}

	// 	subTotal := (int(i.Quantity)) * product.Price

	// 	item := &TransactionItems{
	// 		ProductID: (int(i.ProductID)),
	// 		Quantity: i.Quantity,
	// 		Price: subTotal,
	// 	}

	// 	items = append(items, *item)

	// 	totalPrice += subTotal
	// 	product.Stock = product.Stock - i.Quantity

	// 	err = usecase.ProductRepo.UpdateProductById(int(i.ProductID), product)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("data can't be changed")
	// 	}
	// }

	// transaction := &Transaction{
	// 	TimeStamp: time.Now(),
	// 	Total: totalPrice,
	// 	Items: items,
	// }

	// err := usecase.Repo.CreateTransaction(transaction)
	// if err != nil {
	// 	return nil, fmt.Errorf("data can't added")
	// }

	// newTransaction, err := usecase.Repo.GetTransactionById(transaction.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("id transaction not found")
	// }
	for _, i := range req.Items {
		product, err := usecase.ProductRepo.GetProductById(i.ProductID)
		if err != nil {
			return nil, fmt.Errorf("produk dengan ID tidak ditemukan: %d", i.ProductID)
		}
	
		if i.Quantity > product.Stock {
			return nil, fmt.Errorf("stok tidak mencukupi untuk produk: %s", product.Name)
		}
	
		subTotal := i.Quantity * product.Price
	
		item := &TransactionItems{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			Price:     subTotal,
		}
	
		items = append(items, *item)
	
		totalPrice += subTotal
		product.Stock -= i.Quantity
	
		err = usecase.ProductRepo.UpdateProductById(i.ProductID, product)
		if err != nil {
			return nil, fmt.Errorf("gagal memperbarui produk")
		}
	}
	
	transaction := &Transaction{
		TimeStamp: time.Now(),
		Total:     totalPrice,
		Items:     items,
	}
	
	err := usecase.Repo.CreateTransaction(transaction)
	if err != nil {
		return nil, fmt.Errorf("gagal menambahkan data")
	}
	
	newTransaction, err := usecase.Repo.GetTransactionById(transaction.Id)
	if err != nil {
		return nil, fmt.Errorf("ID transaksi tidak ditemukan")
	}
	
	return newTransaction, nil
}

