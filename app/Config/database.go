package config

import (
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
  "github.com/joho/godotenv"
  "log"
  "os"
)

type Database struct {
  dbname string
  user string
  password string
  port string
  driver string

  host string
}

func Connect() *gorm.DB {

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  config := Database{
    dbname: os.Getenv("APP_DATABASE"),
		user: os.Getenv("APP_USERNAME"),
		password: os.Getenv("APP_PASSWORD"),
		driver: os.Getenv("APP_CONNECTION"),
    port: os.Getenv("APP_DB_PORT"),
    host: os.Getenv("APP_DB_HOST"),
  }

  dsn := config.user+":"+config.password+"@tcp("+config.host+":"+config.port+")/"+config.dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
  
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  
  if err != nil {
    panic("failed to connect database")
  }

	return db
}