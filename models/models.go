package models

import "time"

type Address struct {
	ID int `json:"id,primary_key"`
	//User        []User
	AddressName string    `gorm:"size:255"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Food struct {
	ID   int    `json:"id,primary_key"`
	Name string `gorm:"size:255"`
	//OrderItems []OrderItems
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID         int    `json:"id,primary_key"`
	FirstName  string `gorm:"size:255"`
	SecondName string `gorm:"size:255"`
	//Orders     []Orders
	AddressID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Orders struct {
	ID int `json:"id,primary_key"`
	//OrderItems []OrderItems
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItems struct {
	ID        int `json:"id,primary_key"`
	OrderId   int
	FoodID    int
	Price     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
