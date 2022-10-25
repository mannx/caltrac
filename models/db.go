package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ItemEntry contains a single item name, and its ID is linked to other entries
type ItemName struct {
	gorm.Model

	Name string `gorm:"column:Name"`
}

// ItemEntry contains the information for a given meal
type MealEntry struct {
	gorm.Model

	Name         uint // id of the ItemName table
	Date         datatypes.Date
	CalorieCount uint
	Servings     uint
}
