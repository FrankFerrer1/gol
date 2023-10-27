package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	inject := initializeAll()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		inject.Database.Stats()
		return c.SendString("Hello World!")
	})

	//Test query
	app.Get("/test", func(c *fiber.Ctx) error {
		query := "SELECT id, assetname, links FROM links"

		// Execute the SQL query
		rows, err := inject.Database.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		// Create a struct to store the query results
		type Link struct {
			ID        int    // Assuming id is an integer
			AssetName string // Assuming assetname is a string
			Links     string // Assuming links is a string
		}

		// Create a slice to store the query results
		var results []Link

		// Process the query results
		for rows.Next() {
			var link Link
			if err := rows.Scan(&link.ID, &link.AssetName, &link.Links); err != nil {
				return err
			}
			results = append(results, link)
		}
		return c.JSON(results)
	})

	app.Listen(":9000")
}
