package database

import (
	"PI6/models/entity"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func CheckDatabase() error {
	db, err := GetConn()
	if err != nil {
		return err
	}

	return db.AutoMigrate(&entity.Chemical{}, &entity.PriceUnity{})
}

func CloseConn(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func GetConn() (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             10 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Error,     // Log level
			IgnoreRecordNotFoundError: false,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Enable color
		},
	)

	conn, err := gorm.Open(sqlserver.Open(getDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := conn.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetConnMaxLifetime(time.Minute * 2)
	sqlDb.SetMaxIdleConns(1000)
	sqlDb.SetMaxOpenConns(1000)
	sqlDb.SetConnMaxIdleTime(time.Minute * 2)
	return conn, nil
}

// GetDSN get DB's DSN based on env vars
func getDSN() string {

	schemaName := viper.GetString("DB_SCHEMA")
	if len(schemaName) == 0 {
		log.Fatal(errors.New("unauthorized, expected an schema for db"))
	}

	user := viper.GetString("DB_USER")
	if len(user) == 0 {
		log.Fatal(errors.New("unauthorized, expected an user name for db"))
	}

	password := viper.GetString("DB_PASSWORD")
	if len(password) == 0 {
		log.Fatal(errors.New("unauthorized, expected a password for db"))
	}

	address := viper.GetString("DB_ADDRESS")
	if len(address) == 0 {
		log.Fatal(errors.New("unauthorized, expected an address for db"))
	}

	port := viper.GetInt("DB_PORT")
	if len(address) == 0 {
		log.Fatal(errors.New("unauthorized, expected an port for db"))
	}

	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, address, port, schemaName)
}
