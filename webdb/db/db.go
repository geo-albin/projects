package db

import (
	"github.com/gofrs/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type User struct {
	//gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;`
	FirstName string
	LastName  string
	Age       uint8
}

func ConnectToDB() error {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if nil != err {
		return err
	}
	db = database
	db.AutoMigrate(&User{})
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	u.ID = uuid
	return nil
}

func InsertUser(u *User) error {
	result := db.Select("ID", "FirstName", "LastName", "Age").Create(u)
	return result.Error
}

func DeleteAllUsers() error {
	result := db.Exec("DELETE FROM users")
	return result.Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
