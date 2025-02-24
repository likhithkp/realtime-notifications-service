package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	DB   *pgxpool.Pool
	once sync.Once
)

// ConnectDB initializes the PostgreSQL connection pool
func ConnectDB() {
	once.Do(func() {
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found")
		}

		dbURL := os.Getenv("POSTGRES_URL")
		if dbURL == "" {
			log.Fatal("POSTGRES_URL is not set in .env")
		}

		// PostgreSQL connection pool configuration
		config, err := pgxpool.ParseConfig(dbURL)
		if err != nil {
			log.Fatalf("Unable to parse database config: %v", err)
		}

		// Set connection timeout (optional)
		config.MaxConns = 10
		config.MinConns = 2
		config.HealthCheckPeriod = 1 * time.Minute

		// Connect to DB
		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Unable to connect to database: %v", err)
		}

		DB = pool
		fmt.Println("✅ Connected to PostgreSQL")
	})
}

// GetDB returns the DB connection pool
func GetDB() *pgxpool.Pool {
	if DB == nil {
		ConnectDB()
	}
	return DB
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("❌ Database connection closed")
	}
}
