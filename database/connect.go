package database

import (
	"fmt"
	"log"
	"mvcProjectV1/app/model"
	"mvcProjectV1/config"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	models := []interface{}{&model.Company{}, &model.Note{}, &model.User{}, &model.Tests{}}
	DB.AutoMigrate(models...)
	fmt.Println("Database Migrated")
	fmt.Println("Connection Opened to Database")
}
func Query(model interface{}, r *fiber.Ctx) *gorm.DB {
	_query := DB.Model(model)
	_query.Scopes(Page(r))

	return _query
}
func Param(r *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Query("id") != "" {
			db.Where("id = ?", r.Query("id"))
		}
		if r.Query("sort") != "" {
			matched, _ := (regexp.MatchString(`^-`, r.Query("sort")))
			sort := r.Query("sort")
			if matched {
				db.Order(fmt.Sprintf("%s DESC", strings.Replace(sort, "-", "", 1)))
			} else {
				db.Order(fmt.Sprintf("%s ASC", sort))
			}
		}
		return db
	}
}
func Page(r *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	// go func(db *gorm.DB) *gorm.DB {
	// var count int64
	// query.Count(&count)
	// r.Append("Total", fmt.Sprintf("%d", count))

	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.Query("page"))
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(r.Query("page_size"))
		switch {
		case pageSize <= 0:
			pageSize = 10
		}
		var offset int
		if page == 1 {
			offset = 0
		} else {
			offset = (page - 1) * pageSize
		}

		r.Append("Page", fmt.Sprintf("%d", page))
		r.Append("Offset", fmt.Sprintf("%d", offset))
		r.Append("Limit", fmt.Sprintf("%d", pageSize))
		return db.Limit(pageSize).Offset(offset)
	}
}

//////////////////////////////////////////////////////////
