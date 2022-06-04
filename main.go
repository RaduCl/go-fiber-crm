package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/raducl/go-fiber-crm/database"
	"github.com/raducl/go-fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	leadsRoute := "/v1/leads"
	leadsRouteById := leadsRoute + "/:id"

	app.Get(leadsRouteById, lead.GetLead)
	app.Delete(leadsRouteById, lead.DeleteLead)
	app.Get(leadsRoute, lead.GetLeads)
	app.Post(leadsRoute, lead.CreateLead)
}

func initDb() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to the DB")
	}
	fmt.Println("Db connection opened successfully.")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Db migrated")
}

func main() {
	app := fiber.New()
	initDb()
	setupRoutes(app)
	fmt.Println("App started on port 3000")
	app.Listen(3000)
	defer database.DBConn.Close()
}
