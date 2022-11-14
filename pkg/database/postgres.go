package database

import (
	"fmt"
	"gigmile/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgresDB struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db := postgresql()
	return db
}

func NewInstance(db *gorm.DB) *PostgresDB {
	return &PostgresDB{DB: db}
}

// Init sets up the mongodb instance
func postgresql() *gorm.DB {
	// Database Variables
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	DBTimeZone := os.Getenv("DB_TIMEZONE")
	DBMode := os.Getenv("DB_MODE")
	var dsn string
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", DBHost, DBUser, DBPass, DBName, DBPort, DBMode, DBTimeZone)
	} else {
		dsn = databaseUrl
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	//err = db.AutoMigrate(&model.Country{})
	//if err != nil {
	//	log.Fatalf("failed to automigrate: %v", err)
	//}

	log.Println("Connected to the database")

	return db
}

func (postgresDB *PostgresDB) CreateCountry(country *model.Country) (*model.Country, error) {
	err := postgresDB.DB.Create(country).Error
	return country, err
}

func (postgresDB *PostgresDB) FindCountryByID(ID string) (*model.Country, error) {
	var country *model.Country
	err := postgresDB.DB.Where("id = ?", ID).First(&country).Error
	return country, err
}

func (postgresDB *PostgresDB) UpdateCountry(id string, update map[string]interface{}) error {

	result :=
		postgresDB.DB.Model(model.Country{}).
			Where("id = ?", id).
			Updates(&update)
	return result.Error
}

func (postgresDB *PostgresDB) GetCountries() []model.Country {
	var countries []model.Country
	postgresDB.DB.Find(&countries)

	return countries
}
