package migrations

import (
	"food-delivery/database"
	"food-delivery/models"
)

func Migrate() {
	database.GetDB().AutoMigrate(models.User{}, models.Food{}, models.Address{}, models.Orders{}, models.OrderItems{})
}
