package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"repos/task_manager/src/utils"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const env_path = "/app/.env"

type DatabaseConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func InitDB() (*sql.DB, error) {

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

	db, err := sql.Open("postgres", pgsqlConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
