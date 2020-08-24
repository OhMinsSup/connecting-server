package database

import (
	"connecting-server/lib"
	"connecting-server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/pborman/uuid"
	"log"
	"reflect"
	"strings"
)

func Initialize() (*gorm.DB, error) {
	// viper package read .env
	dbUser := lib.GetEnvWithKey("POSTGRES_USER")
	dbPassword := lib.GetEnvWithKey("POSTGRES_PASSWORD")
	dbName := lib.GetEnvWithKey("POSTGRES_DB")
	dbHost := lib.GetEnvWithKey("POSTGRES_HOST")
	dbPort := lib.GetEnvWithKey("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting
	dbConfig := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", dbConfig)

	if err != nil {
		log.Println("database connection error:", err)
	}

	// set connection pool
	db.DB().SetMaxIdleConns(10)
	// Logs SQL
	db.LogMode(true)
	db.Set("gorm:table_options", "charset=utf8")
	// created uuid
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", BeforeCreateUUID)
	log.Println("Connected to database")
	Migrate(db)

	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.UserProfile{},
		&model.UserMeta{},
		&model.AuthToken{})

	db.Model(&model.UserProfile{}).AddForeignKey("user_ref", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.UserMeta{}).AddForeignKey("user_ref", "users(id)", "CASCADE", "CASCADE")
	db.Model(&model.AuthToken{}).AddForeignKey("user_ref", "users(id)", "CASCADE", "CASCADE")

	log.Println("Auto Migration has beed processed")
}

func BeforeCreateUUID(scope *gorm.Scope) {
	reflectValue := reflect.Indirect(reflect.ValueOf(scope.Value))
	if strings.Contains(string(reflectValue.Type().Field(0).Tag), "uuid") {
		uuid.SetClockSequence(-1)
		scope.SetColumn("id", uuid.NewUUID().String())
	}
}

func Inject(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
