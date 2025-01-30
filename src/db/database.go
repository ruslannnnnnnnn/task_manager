package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"repos/task_manager/src/entity"
	"repos/task_manager/src/utils"
	"strconv"

	"github.com/joho/godotenv"
)

const env_path = "/app/.env"

type DatabaseConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func getDBConfig() (string, error) {
	err := godotenv.Load(env_path)

	if err != nil {
		log.Fatal(err)
	}

	db_port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	utils.LogIfError(err)

	dbCfg := DatabaseConfig{
		os.Getenv("POSTGRES_HOST"),
		db_port,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	}

	pgsqlConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbCfg.host, dbCfg.port, dbCfg.user, dbCfg.password, dbCfg.dbname)
	return pgsqlConfig, err
}

func InitDB() (*gorm.DB, error) {

	pgsqlConfig, err := getDBConfig()
	utils.LogIfError(err)

	db, err := gorm.Open(postgres.Open(pgsqlConfig), &gorm.Config{})

	return db, err
}

func InitAutoMigrations() {
	db, err := InitDB()
	utils.LogIfError(err)

	err = db.AutoMigrate(&entity.Task{})
	utils.LogIfError(err)
}
