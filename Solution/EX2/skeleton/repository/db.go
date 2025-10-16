package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Coca Cola", Price: 2, Description: "Refreshing soda"},
		{ID: 2, Name: "Fanta", Price: 2, Description: "Orange flavored drink"},
		{ID: 3, Name: "Sprite", Price: 2, Description: "Lemon-lime soda"},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, Amount: 2, CreatedAt: time.Now().Add(-10 * time.Minute)}, // 10 minutes ago
		{DrinkID: 2, Amount: 1, CreatedAt: time.Now().Add(-5 * time.Minute)},  // 5 minutes ago
		{DrinkID: 3, Amount: 3, CreatedAt: time.Now()},                        // now
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {

	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
