package root

import (
	homepage "p_dm_aa01_hafsa/internal/modules/root/page"
	"p_dm_aa01_hafsa/internal/server"
	"p_dm_aa01_hafsa/util"

	"github.com/gofiber/fiber/v2"
)

func Init(server *server.Server) {
	router := server.GetRouter()

	// Home page
	router.Get("/", func(c *fiber.Ctx) error {
		return util.Render(c, homepage.HomePage())
	})
}
