package transactions

import (
	"Pos-Design/modules/products"
	"context"
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

func (usecase UseCase) GetTransactionById(ctx context.Context) (*Transaction, error) {

	idprms := ctx.Value("idPrms")
	id := idprms.(int)

	transaction, err := usecase.Repo.GetTransactionById(id)
	return transaction, err
}

func (usecase UseCase) CreateTransaction(ctx context.Context, req *CreateTransactionRequest ) (*Transaction, error) {

	id_admin := ctx.Value("id_admin")
	
	
	items := []TransactionItems{}
	totalPrice := 0
	
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
	
		err = usecase.ProductRepo.EditProductById(i.ProductID, product)
		if err != nil {
			return nil, fmt.Errorf("gagal memperbarui produk")
		}
	}
	
	transaction := &Transaction{
		AdminID: id_admin.(int),
		TimeStamp: time.Now(),
		Total:     totalPrice,
		Items:     items,
		
	}

	fmt.Println(transaction.AdminID)
	
	err := usecase.Repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	
	newTransaction, err := usecase.Repo.GetTransactionById(transaction.Id)
	if err != nil {
		return nil, err
	}
	
	return newTransaction, nil
}

