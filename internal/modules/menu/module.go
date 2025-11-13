package menu

import (
	menupage "p_dm_aa01_hafsa/internal/modules/menu/page"
	"p_dm_aa01_hafsa/internal/server"
	"p_dm_aa01_hafsa/util"

	"github.com/gofiber/fiber/v2"
)

func Init(server *server.Server) {
	router := server.GetRouter()

	// Menu page
	router.Get("/menu", func(c *fiber.Ctx) error {
		return util.Render(c, menupage.MenuPage())
	})
}
