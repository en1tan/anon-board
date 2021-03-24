package services

import (
	"errors"

	"github.com/en1tan/anon-board/database"
	"github.com/en1tan/anon-board/models"
	"github.com/en1tan/anon-board/utils"
	"gorm.io/gorm"
)

type ThreadService interface {
	List(page int) (error, *[]models.Thread)
	GetById(id string) (error, *models.Thread)
	Create(t models.Thread) (error, *models.Thread)
	Report(id string) error
	DeleteWithPassword(id, password string) error
}

type threadService struct {
	db *gorm.DB
}

func NewThreadService(conn database.DatabaseConnection) ThreadService {
	return &threadService{db: conn.Get()}
}

func (ts *threadService) List(page int) (error, *[]models.Thread) {
	var t []models.Thread
	offset := 0
	if page > 0 {
		offset = page - 1
	}
	result := ts.db.Limit(10).Offset(offset).Find(&t).Order("bumped_on DESC")
	return result.Error, &t
}

func (ts *threadService) GetById(id string) (error, *models.Thread) {
	var t models.Thread
	result := ts.db.First(&t).Where("id = ?", id)
	return result.Error, &t
}

func (ts *threadService) Create(t models.Thread) (error, *models.Thread) {
	password, err := utils.HashPassword(t.DeletePassword)
	if err != nil {
		return err, nil
	}
	t.DeletePassword = password
	result := ts.db.Create(&t)
	return result.Error, &t
}

func (ts *threadService) Report(id string) error {
	return ts.db.Transaction(func(tx *gorm.DB) error {
		var t models.Thread
		if result := tx.Where("id = ?", id).First(&t); result.Error != nil {
			return result.Error
		}
		if result := tx.Model(&t).Where("id = ?", id).Update("reported", true); result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (ts *threadService) DeleteWithPassword(id, password string) error {
	return ts.db.Transaction(func(tx *gorm.DB) error {
		var t models.Thread
		if result := tx.Where("id = ?", id).First(&t); result.Error != nil {
			return result.Error
		}
		if !utils.CheckPassword(password, t.DeletePassword) {
			return errors.New("Incorrect Password")
		}
		if result := tx.Model(&t).Where("id = ?", id).Update("text", "[deleted]"); result.Error != nil {
			return result.Error
		}
		return nil
	})
}
