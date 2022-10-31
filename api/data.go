package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	models "github.com/mannx/caltrac/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// GetMealEntryWithDate returns all MealEntry data for a specific day
func GetMealEntryWithDate(c echo.Context, db *gorm.DB) error {
	var month, year, day int

	err := echo.QueryParamsBinder(c).
		Int("month", &month).
		Int("year", &year).
		Int("day", &day).
		BindError()
	if err != nil {
		return LogAndReturnError(c, "[GetMealEntryWithDate] Unable to bind parameters for date", err)
	}

	date := datatypes.Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))

	var meals []models.MealEntry
	res := db.Where("Date = ?", date).Find(&meals)
	if res.Error != nil {
		return LogAndReturnError(c, "[GetMealEntryWithDate] Unable to retrieve data.", res.Error)
	}

	return c.JSON(http.StatusOK, &meals)
}
