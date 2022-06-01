package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/todoApp/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var log = logrus.WithField("ctx", "database")

type Item struct {
	gorm.Model
	Description string `json:"description"`
	Completed   bool
}

type Database struct {
	client *gorm.DB
}

type Storage interface {
	AddItem(description string) (*Item, error)
	GetSingleItem(id uint32) (*Item, error)
	GetItems() ([]Item, error)
	UpdateItemDescription(id uint32, description string) (*Item, error)
	UpdateItemCompletion(id uint32, completed bool) (*Item, error)
	DeleteItem(id uint32) (*Item, error)

	HealthChecker()
	Close()
}

func dsn(conf *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Database.Username, conf.Database.Password, conf.Database.HostName, conf.Database.Port, conf.Database.Name)
}

func NewDBConnection(conf *config.Config) (*gorm.DB, error) {
	if conf.Application.MockStorage {
		return gorm.Open(
			sqlite.Open("file::memory:?cache=shared"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
		)
	}
	return gorm.Open(mysql.Open(dsn(conf)), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
}

func NewSQLStore(conn *gorm.DB) (Storage, error) {
	gormDatabase := &Database{
		client: conn,
	}

	err := conn.AutoMigrate(&Item{})
	if err != nil {
		return nil, err
	}

	return gormDatabase, nil
}

func (d *Database) AddItem(description string) (*Item, error) {
	todo := &Item{
		Description: description,
		Completed:   false,
	}

	if err := d.client.Create(todo).Error; err != nil {
		log.WithFields(logrus.Fields{
			"description": description,
		}).WithError(err).Error("failed to add todo item to database")
		return nil, err
	}

	log.Info("todo item successfully added into database")
	return todo, nil
}

func (d *Database) GetSingleItem(id uint32) (*Item, error) {
	var todo Item
	if err := d.client.Where("id = ?", id).First(&todo, id).Error; err != nil {
		log.WithFields(logrus.Fields{
			"id": id,
		}).WithError(err).Error("error getting todo item from database")
		return nil, err
	}

	log.Info("todo item successfully retrieved from database")
	return &todo, nil
}

func (d *Database) GetItems() ([]Item, error) {
	var todos []Item
	if err := d.client.Find(&todos).Error; err != nil {
		log.WithError(err).Error("error getting all todo items from database")
		return nil, err
	}

	log.Info("todo items successfully retrieved from")
	return todos, nil
}

func (d *Database) UpdateItemDescription(id uint32, description string) (*Item, error) {
	foundTodo, err := d.GetSingleItem(id)
	if err != nil {
		return nil, err
	}
	foundTodo.Description = description

	if err = d.client.Save(&foundTodo).Error; err != nil {
		log.WithFields(logrus.Fields{
			"id": id,
		}).WithError(err).Error("failed to update todo item")
	}

	log.WithFields(logrus.Fields{
		"id": foundTodo.ID,
	}).Info("todo item successfully updated")

	return foundTodo, nil
}

func (d *Database) UpdateItemCompletion(id uint32, completed bool) (*Item, error) {
	foundTodo, err := d.GetSingleItem(id)
	if err != nil {
		return nil, err
	}
	foundTodo.Completed = completed

	if err = d.client.Save(&foundTodo).Error; err != nil {
		log.WithFields(logrus.Fields{
			"id": id,
		}).WithError(err).Error("failed to update todo item")
	}

	log.WithFields(logrus.Fields{
		"id": foundTodo.ID,
	}).Info("todo item successfully updated")

	return foundTodo, nil
}

func (d *Database) DeleteItem(id uint32) (*Item, error) {
	var item Item
	if err := d.client.Where("id = ?", id).Delete(&item).Error; err != nil {
		log.WithError(err).Error("failed to delete todo item from database")
		return nil, err
	}

	log.WithFields(logrus.Fields{
		"id": item.ID,
	}).Info("todo item successfully deleted")

	return &item, nil
}

func (d *Database) HealthChecker() {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Close() {}
