package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zyzmoz/go-basic-crm/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead
	db.Find(&lead, id)

	return c.JSON(lead)
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead

	db.Find(&leads)
	return c.JSON(leads)
}

func CreateLead(c *fiber.Ctx) error {
	db := database.DBConn

	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).JSON(err)
		return err
	}

	db.Create(&lead)

	return c.JSON(lead)

}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead
	db.First(&lead, id)

	if lead.Name == "" {
		return c.Status(500).SendString("No Lead found")

	}
	db.Delete(&lead)

	return c.SendString("Lead successfully deleted")
}
