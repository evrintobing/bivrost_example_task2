package repository

import (
	"log"

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
	dataOrder := models.Order{
		IDProduk:     order.IDProduk,
		JumlahProduk: order.JumlahProduk,
	}
	db := r.db.Create(&dataOrder)
	if db.Error != nil {
		return nil, db.Error
	}
	return &dataOrder, nil
}

// GetItems implements oreders.OrderRepository
func (r *repository) GetItems() (*[]models.Items, error) {
	var item []models.Items
	db := r.db.Raw("Select * from items").Scan(&item)
	log.Println(item, "isi item")
	if db.Error != nil {
		return nil, db.Error
	}
	return &item, nil
}

// GetOrders implements oreders.OrderRepository
func (r *repository) GetOrders() (*[]models.GetOrder, error) {
	var order []models.GetOrder
	db := r.db.Raw("Select * from orders").Scan(&order)
	if db.Error != nil {
		return nil, db.Error
	}
	return &order, nil
}
