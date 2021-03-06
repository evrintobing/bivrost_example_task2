package orders

import "github.com/evrintobing/bivrost_example_task2/models"

type OrderUsecase interface {
	GetOrder() (*[]models.GetOrder, error)
	GetItems() (*[]models.Items, error)
	AddOrder(idProduk, jumlahProduk int) (*models.Order, error)
}
