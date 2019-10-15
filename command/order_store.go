package command

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OrdersStore struct {
	db *gorm.DB
}

type Order struct {
	gorm.Model
	UserId string
	Order  string
}

func NewOrdersStore() (*OrdersStore, error) {
	db, err := gorm.Open("sqlite3", "orders.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Order{})
	return &OrdersStore{db}, nil
}

func (o OrdersStore) Add(userId string, order *string) error {
	return o.db.Create(&Order{UserId: userId, Order: *order}).Error
}

func (o OrdersStore) Remove(userId string, order *string) error {
	return o.db.Where(&Order{UserId: userId, Order: *order}).Delete(Order{}).Error
}

func (o OrdersStore) List() ([]Order, error) {
	var orders []Order

	result := o.db.Where(Order{}).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o OrdersStore) Close() error {
	if o.db != nil {
		return o.db.Close()
	}

	return nil
}
