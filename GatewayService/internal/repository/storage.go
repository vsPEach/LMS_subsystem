package repository

import (
	"github.com/google/uuid"
	"github.com/vsPEach/LMS_subsystem/DistributorService/internal/models"
	"github.com/vsPEach/LMS_subsystem/DistributorService/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
	dsn        string
}

func NewDatabase() *Database {
	return &Database{dsn: "host=localhost port=5432 user=postgres password=postgres dbname=auth sslmode=disable"}
}

func (db *Database) Create(user models.User) (err error) {
	user.Password, err = utils.HashPassword(user.Password)
	user.ID = uuid.New()
	if err != nil {
		return
	}
	res := db.connection.Create(&user)
	return res.Error
}

func (db *Database) Read(password string) error {
	var user models.User
	res := db.connection.First(&user)
	if res.Error != nil {
		return res.Error
	}
	if err := utils.ValidatePassword(user.Password, password); err != nil {
		return err
	}
	return nil
}

func (db *Database) Connect() error {
	c, err := gorm.Open(postgres.Open(db.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.connection = c
	err = db.connection.AutoMigrate(&models.User{})
	return err
}
