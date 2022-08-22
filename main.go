package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zyzmoz/go-basic-crm/database"
	"github.com/zyzmoz/go-basic-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDb() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "crm.db")

	if err != nil {
		panic("failed to create database connection")
	}

	fmt.Println("Connected to the database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDb()
	setupRoutes(app)
	app.Listen(":8080")
	defer database.DBConn.Close()
}
