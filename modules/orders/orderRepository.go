package orders

import "github.com/evrintobing/bivrost_example_task2/models"

type OrderRepository interface {
	GetItems() (*[]models.Items, error)
	GetOrders() (*[]models.Order, error)
	AddOrder(order *models.Order) (*models.Order, error)
}
