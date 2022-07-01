package app

import (
	"os"

	"mvcProjectV1/database"
	"mvcProjectV1/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html"
)

func CreateApp() *fiber.App {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Static("/", "./dist")
	database.ConnectDB()

	os.Setenv("TOEKN_SECRET_KEY", "secret_key_test_me")

	routes.Web(app)
	routes.Api(app)
	// err := filepath.Walk("./routes",
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}

	// 		log.Println(path, info.Size())
	// 		return nil
	// 	})
	// if err != nil {
	// 	log.Println(err)
	// }
	// routes.Api(app)

	return app
}
