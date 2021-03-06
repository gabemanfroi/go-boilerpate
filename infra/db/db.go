package db

import (
	"fmt"
	"github.com/gabemanfroi/go-boilerplate/infra/core"
	"github.com/gabemanfroi/go-boilerplate/internal/utils"
	"github.com/golobby/container/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresSession struct{}

func CreateSqlInstance() *gorm.DB {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		core.AppConfig.DbHost,
		core.AppConfig.DbPort,
		core.AppConfig.DbUser,
		core.AppConfig.DbName,
		core.AppConfig.DbPassword)

	db, err := gorm.Open(postgres.Open(dbUrl))
	utils.Check(err, "Error creating database connection ")

	return db
}

func GetSqlInstance() *gorm.DB {
	var injector *gorm.DB
	utils.Check(container.Resolve(&injector), "Error while retrieving Database instance ")
	return injector
}

func Migrate() {
	log.Printf("Starting Database Migrations...")
	var db *gorm.DB
	err := container.Resolve(&db)

	if err != nil {
		log.Fatal("error while retrieving database instance" + err.Error())
	}

	// register your models following the example bellow:
	//utils.Check(db.AutoMigrate(&models.yourModel{}), "failed to migrate yourModel")

	log.Printf("Database Migrations Completed...")
}
