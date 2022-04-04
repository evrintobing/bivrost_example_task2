package usecase

import (
	"github.com/evrintobing/bivrost_example_task2/models"
	"github.com/evrintobing/bivrost_example_task2/modules/orders"
)

type usecase struct {
	repo orders.OrderRepository
}

func NewOrderUsecase(repo orders.OrderRepository) orders.OrderUsecase {
	return &usecase{
		repo: repo,
	}
}

// AddOrder implements orders.OrderUsecase
func (uc *usecase) AddOrder(idProduk int, jumlahProduk int) (*models.Order, error) {
	var order models.Order
	order.IDProduk = idProduk
	order.JumlahProduk = jumlahProduk

	data, err := uc.repo.AddOrder(&order)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetItems implements orders.OrderUsecase
func (uc *usecase) GetItems() (*[]models.Items, error) {
	data, err := uc.repo.GetItems()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetOrder implements orders.OrderUsecase
func (uc *usecase) GetOrder() (*[]models.GetOrder, error) {
	data, err := uc.repo.GetOrders()
	if err != nil {
		return nil, err
	}
	return data, nil
}
