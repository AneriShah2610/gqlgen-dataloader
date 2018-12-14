package gqlgen_authentication

import (
	"fmt"
	"log"
	"sync"

	"github.com/aneri/gqlgen-authentication/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	Db *gorm.DB
}

var Once sync.Once
var instance *DbConnection

func Connect() (*DbConnection, error) {
	connectionString := "postgresql://root@localhost:26257/training?sslmode=disable"
	db, err := gorm.Open("postgres", connectionString)
	Once.Do(func() {
		if err != nil {
			log.Fatal("Error while initializing database", err)
		}
		fmt.Println("Database successfully connected")
		instance = &DbConnection{
			Db: db,
		}
	})
	db.AutoMigrate(&models.Application{})
	return instance, err
}
