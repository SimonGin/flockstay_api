package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	// user=postgres.mjlfjnezllktgtpqpsgr password=[YOUR-PASSWORD] host=aws-0-ap-northeast-2.pooler.supabase.com port=6543 dbname=postgres
	dsn := os.Getenv("DB_CONNECT_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
