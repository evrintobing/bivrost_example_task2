package repository

import (
	"github.com/evrintobing/bivrost_example_task2/models"
	"github.com/evrintobing/bivrost_example_task2/modules/orders"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) orders.OrderRepository {
	return &repository{
		db: db,
	}
}

// AddOrder implements oreders.OrderRepository
func (r *repository) AddOrder(order *models.Order) (*models.Order, error) {
	db := r.db.Create(&order)
	if db.Error != nil {
		return nil, db.Error
	}
	return order, nil
}

// GetItems implements oreders.OrderRepository
func (r *repository) GetItems() (*[]models.Items, error) {
	var item []models.Items
	db := r.db.Table("orders").Find(&item)
	if db.Error != nil {
		return nil, db.Error
	}
	return &item, nil
}

// GetOrders implements oreders.OrderRepository
func (r *repository) GetOrders() (*[]models.Order, error) {
	var order []models.Order
	db := r.db.Table("orders").Find(&order)
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}
