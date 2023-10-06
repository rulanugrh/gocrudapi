package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Untuk konfigurasi env
type App struct {
	Database struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}

	Host string
	Port string
}

var app *App

func GetConnection() *gorm.DB {
	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Database.User, conf.Database.Pass, conf.Database.Host, conf.Database.Port, conf.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Cant connect to database : %v", err)
	}

	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	config := App{}
	err := godotenv.Load()
	if err != nil {
		config.Database.Host = "localhost"
		config.Database.Port = "3306"
		config.Database.Name = "gocrud"
		config.Database.User = "root"
		config.Database.Pass = ""

		config.Host = "localhost"
		config.Port = "8080"

		return &config
	}

	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Pass = os.Getenv("DB_PASS")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.Name = os.Getenv("DB_NAME")
	config.Database.User = os.Getenv("DB_USER")

	config.Host = os.Getenv("APP_HOST")
	config.Port = os.Getenv("APP_PORT")

	return &config
}
