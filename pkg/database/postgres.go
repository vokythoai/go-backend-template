package database

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *PostgresDB
	once     sync.Once
)

type PostgresDB struct {
	db *gorm.DB
}

func GetInstance(host, port, user, password, dbname string) (*PostgresDB, error) {
	var err error
	once.Do(func() {
		instance, err = newPostgresDB(host, port, user, password, dbname)
	})
	return instance, err
}

func newPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(5)

	sqlDB.SetConnMaxLifetime(time.Hour)

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresDB) DB() *gorm.DB {
	return p.db
}
