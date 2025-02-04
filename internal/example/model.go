package example

import (
    "go-api/database"
		"time"
)

type Example struct {
	ID        int       `json:"id" gorm:"column:id"`
	Example   string    `json:"example" gorm:"column:example"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Example) TableName() string {
	return "example"
}

func GetExample() ([]Example, error) {
	var example []Example
	if err := database.DB.Find(&example).Error; err != nil {
			return nil, err
	}
	return example, nil
}

func (example *Example) SaveExample() (uint, error) {
	if err := database.DB.Create(&example).Error; err != nil {
		return 0, err
	}
	return uint(example.ID), nil
}

func (example *Example) UpdateExampleByID(id string) error {
	return database.DB.Model(&Example{}).Where("id = ?", id).Updates(example).Error
}

func DeleteExampleByID(id string) error {
	return database.DB.Where("id = ?", id).Delete(&Example{}).Error
}