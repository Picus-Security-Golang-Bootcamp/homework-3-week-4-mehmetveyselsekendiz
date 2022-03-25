package book

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/model/author"
)

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string
	Page        int
	Stock       int
	
}

func (Book) TableName() string {
	return "Book"
}

func (c *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Code : %s, CountryCode : %s,CreatedAt : %s", c.ID, c.Name, c.Code, c.CountryCode, c.CreatedAt.Format("2006-01-02 15:04:05"))
}