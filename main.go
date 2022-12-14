package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	models "github.com/mannx/caltrac/models"
)

// Version of the current build/release
var (
	BuildVersion string = ""
	BuildTime    string = ""
)

// connection to the database for our entire run
var DB *gorm.DB

// name of the database we are using
var dbName string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msgf("CalorieTracker")
	log.Info().Msgf("  => Version: %v", BuildVersion)
	log.Info().Msgf("  => Build Time: %v", BuildTime)

	log.Info().Msg("Initializing environment...")
	Environment.Init()

	dbName = filepath.Join(Environment.DataPath, "db.db")

	log.Info().Msg("Initializing database...")
	log.Debug().Msgf("  => Database path: %v", dbName)
	dbo, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open database...")
	}

	// if we used DB initially, we would be creating a local DB variable and not saving into the global?
	//	recieving error about not using variable DB
	DB = dbo

	log.Info().Msg("Auto migrating the database...")
	migrateDB()

	log.Info().Msg("Initialiing server and middleware")

	e := initServer()

	log.Info().Msg("Starting server...")
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting server down")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// migrateDB ensures all selected models have all fields in the database
func migrateDB() {
	// TODO as neded
	DB.AutoMigrate(&models.ItemName{})
	DB.AutoMigrate(&models.MealEntry{})
}
