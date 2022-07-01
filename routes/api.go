package routes

import (
	"fmt"
	"log"
	"mvcProjectV1/app/model"
	"mvcProjectV1/database"
	"mvcProjectV1/utils/pkg"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Api(app *fiber.App) {
	api := app.Group("/api/")

	api.Get("/1", func(c *fiber.Ctx) error {
		return c.JSON(app)
	})
	api.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err == nil {
			c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		}
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "upload complete"})
	})
	api.Post("/log", func(c *fiber.Ctx) error {

		err, out, errout := pkg.Shellout("ls -ltr")
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println("--- stdout ---")
		fmt.Println(out)
		fmt.Println("--- stderr ---")
		fmt.Println(errout)
		return c.JSON(app)
	})
	api.Get("/test", func(c *fiber.Ctx) error {
		var tests []model.Tests
		db := database.DB
		db.Model(&tests).Scopes(database.Param(c)).Scopes(database.Page(c)).Find(&tests)

		// pagination := database.Pagination{}
		// db := database.DB
		// query := db.Model(&tests)
		// query.Scopes(database.Paginate(tests, &pagination, db))

		return c.JSON(fiber.Map{"status": "success", "data": tests})
		// return c.JSON(pagination)
	})
	api.Post("/test", func(c *fiber.Ctx) error {

		var tests []model.Tests
		// db := database.DB

		for i := 0; i < 10000; i++ {
			var test model.Tests
			test.Name = "test-0" + strconv.Itoa(i)
			test.LastName = "LastName-0" + strconv.Itoa(i)
			tests = append(tests, test)
		}
		db := database.DB
		db.Create(&tests)
		// for i := 0; i < 5000; i++ {
		// 	database.DB.Transaction(func(tx *gorm.DB) error {
		// 		var tests model.Tests
		// 		input := new(model.Tests)
		// 		if err := c.BodyParser(input); err != nil {
		// 			return err
		// 		}
		// 		tests.Name = input.Name + strconv.Itoa(i)
		// 		tests.LastName = input.LastName + strconv.Itoa(i)
		// 		tx.Create(&tests)
		// 		return nil
		// 	})
		// }

		return c.JSON(fiber.Map{"status": "success", "data": "tests"})
	})
	api.Post("/companys", func(c *fiber.Ctx) error {
		type Input struct {
			Name string `json:"name" xml:"name" form:"name"`
		}
		input := new(Input)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		// Create from map
		db := database.DB
		var company model.Company
		company.Name = input.Name
		db.Create(&company)

		return c.JSON(fiber.Map{"status": "success", "message": "create complete", "data": company})
	})
	api.Post("/users", func(c *fiber.Ctx) error {
		// type Input struct {
		// 	Name      string `json:"name" xml:"name" form:"name"`
		// 	CompanyID string `json:"company_id" xml:"company_id" form:"company_id"`
		// }

		input := new(model.User)
		if err := c.BodyParser(input); err != nil {
			// log.Println(input.CompanyID)
			return err
		}

		// Create from map
		db := database.DB
		var user model.User
		user.Name = input.Name
		user.CompanyID = input.CompanyID
		db.Create(&user)

		return c.JSON(fiber.Map{"status": "success", "message": "create complete", "data": user})
	})
	api.Post("/company-users", func(c *fiber.Ctx) error {
		type Input struct {
			Name string `json:"company-name" xml:"company-name" form:"company-name"`
		}
		inputCompany := new(Input)
		if err := c.BodyParser(inputCompany); err != nil {
			return err
		}
		inputUser := new(model.User)
		if err := c.BodyParser(inputUser); err != nil {
			return err
		}

		// Create from map
		db := database.DB

		db.Transaction(func(tx *gorm.DB) error {
			var company model.Company
			tx.Where("name = ?", inputCompany.Name).Find(&company)
			if company.ID == uuid.Must(uuid.Parse("00000000-0000-0000-0000-000000000000")) {
				company.Name = inputCompany.Name
				tx.Create(&company)
			}

			var user model.User
			user.Name = inputUser.Name
			user.CompanyID = company.ID
			tx.Create(&user)
			return nil
		})

		return c.JSON(fiber.Map{"status": "success", "message": "create complete"})
	})
	api.Get("/users", func(c *fiber.Ctx) error {
		var Total int64
		db := database.DB
		var user []model.User

		db.Model(&user).
			// Scopes(database.Page(c)).
			Preload("Company").Count(&Total).Find(&user)
		// log.Println(Total)
		if len(user) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}

		// pkg.Paging(&pkg.Param{
		// 	Total: Total,
		// 	Ctx:   c,
		// }, &paginator)
		// Else return notes
		return c.JSON(user)
	})
	api.Get("/notes", func(c *fiber.Ctx) error {
		var Total int64
		db := database.DB
		var notes []model.Note
		db.Model(notes).
			// Scopes(database.Page(c)).
			Where("notes.id = 'f56a88f8-1dcc-4f1f-94b2-1ec7b8308fb3'").
			Count(&Total).
			Find(&notes)

		// db = db.Where("notes.id = 'f56a88f8-1dcc-4f1f-94b2-1ec7b8308fb3'")
		// page, _ := strconv.Atoi(c.Query("page"))
		// pageSize, _ := strconv.Atoi(c.Query("page_size"))
		// pkg.Paging(&pkg.Param{
		// 	Total: Total,
		// 	Ctx:   c,
		// }, &paginator)
		// utils.Query(&notes)
		// find all notes in the database
		// db.Find(&notes)

		// If no note is present return an error
		if len(notes) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
		}

		// Else return notes
		return c.JSON(notes)
	})

}
