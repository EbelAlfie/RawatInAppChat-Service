package repository

import (
	"chat_service/domain"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		db: openConnection(),
	}
}

func (repo userRepository) Login() {

}

func openConnection() *sql.DB {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	port := os.Getenv("PORT")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/user", user, pass, port)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
