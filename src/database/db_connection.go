package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Connection *gorm.DB
}

func NewDbConnection() *DbConnection {
	conn := &DbConnection{}
	if err := conn.getConnection(); err != nil {
		log.Panic("Failed to connect to the database: ", err)
	}
	return conn
}

func (db *DbConnection) getConnection() error {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	dbname := os.Getenv("DATABASE_NAME")
	password := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
	if conn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{}); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	} else {
		db.Connection = conn
		fmt.Println("Connected to the database successfully")
		return nil
	}
}
