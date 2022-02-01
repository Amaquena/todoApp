package util

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"github.com/todolist/util/models"
	"log"
	"time"
)

// TODO Remove contexts in the db call

type DBConnection struct {
	dbConfig *models.DBConfig
	db       *sql.DB
	l        *log.Logger
}

var dbConn = &DBConnection{}

func dsn(conf models.DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.Username, conf.Password, conf.Hostname, conf.Port)
}

func loadDBConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("database")
	err := viper.ReadInConfig()
	if err != nil {
		dbConn.l.Println("Error reading in config")
		return err
	}
	err = viper.Unmarshal(&dbConn.dbConfig)
	if err != nil {
		return err
	}
	return nil
}

func CreateDBConnection(l *log.Logger, path string) error {
	dbConn.l = l
	err := loadDBConfig(path)
	if err != nil {
		l.Println("Failed to load database config file")
		return err
	}

	dbConn.db, err = sql.Open("mysql", dsn(*dbConn.dbConfig))
	if err != nil {
		l.Println("Error opening DB connection")
		return err
	}

	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		l.Println("Error closing Database connection")
	//	}
	//}(dbConn.db)

	err = createDatabase()
	if err != nil {
		return err
	}

	err = selectDatabase(dbConn.dbConfig.DBName)
	if err != nil {
		return err
	}

	err = createItemsTable()
	if err != nil {
		return err
	}

	dbConn.db.SetMaxOpenConns(20)
	dbConn.db.SetMaxIdleConns(20)
	dbConn.db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err = dbConn.db.PingContext(ctx)
	if err != nil {
		l.Println("Errors pinging DB")
		return err
	}
	l.Printf("Connected to DB %s successfully\n", dbConn.dbConfig.DBName)
	return nil
}

func createDatabase() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := `CREATE DATABASE IF NOT EXISTS ` + dbConn.dbConfig.DBName
	_, err := dbConn.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func createItemsTable() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	query := `
		CREATE TABLE IF NOT EXISTS items (
	    	id INT AUTO_INCREMENT,
			item TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	_, err := dbConn.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return err
}

func selectDatabase(dbname string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := dbConn.db.ExecContext(ctx, "USE "+ dbname)
	if err != nil {
		return err
	}
	return nil
}

func GetDbConnection() *sql.DB {
	return dbConn.db
}