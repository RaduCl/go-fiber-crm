package lead

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/raducl/go-fiber-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	ctx.JSON(lead)
}

func GetLeads(ctx *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead

	db.Find(&leads)
	ctx.JSON(leads)
}

func CreateLead(ctx *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)

	if err := ctx.BodyParser(&lead); err != nil {
		ctx.Status(fiber.StatusServiceUnavailable).Send(err)
		return
	}

	db.Create(&lead)
	ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var lead Lead

	result := db.Delete(&lead, id)

	if result.RowsAffected == 0 {
		ctx.Send(fiber.StatusNotFound, fmt.Sprintf("Lead with id: %d was not found", lead.ID))
		return
	}

	ctx.Send("deleted")
}
