package config

import (
	"database/sql"
	"evently/models"
	"fmt"
	//"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDB() *sql.DB {
	// Define the connection string
	dsn := "root:Techie1@db@tcp(127.0.0.1:3306)/evently"

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil

	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing the database connection:", err)
			return
		}
	}(db)

	// Verify the connection is successful
	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging the database:", err)
		return nil
	}

	fmt.Println("Successfully connected to the database!")

	// Set the database connection pool settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

var DB **gorm.DB

func GetDbInstance() **gorm.DB {
	if DB == nil {
		InitGorm()
		return DB
	}
	
	return DB
}

func InitGorm() {
	if DB != nil {
		return
	}

	// Connect to the database
	dsn := "root:Techie1@db@tcp(127.0.0.1:3306)/evently?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate the schema
	runMigrations(&db)

	// initialize the global DB variable
	DB = &db
}

func runMigrations(db **gorm.DB) {
	// Migrate the schema
	err := (*db).AutoMigrate(
		&models.Event{},
	)

	if err != nil {
		fmt.Println("Error migrating the schema:", err)
		return
	} else {
		fmt.Println("Successfully migrated the schema!")
	}
}
