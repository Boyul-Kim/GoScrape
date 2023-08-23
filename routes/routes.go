package router

import (
	"goscrape/handler"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(server *fiber.App) {

	scraper := server.Group("/scraper")
	scraper.Get("query", handler.ScrapeGoQueryCostco)
	scraper.Get("/albertsons", handler.ScrapeGoQueryAlbertsons)
	scraper.Get("/hmart", handler.ScrapeGoQueryHmart)
	scraper.Get("/target", handler.ScrapeGoQueryTarget)
}
