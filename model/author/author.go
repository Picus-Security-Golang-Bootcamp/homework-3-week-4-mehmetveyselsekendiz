package author

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/model/book"
)

type Author struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string
	Books		[]book.Book `gorm:"foreignKey:Author;references:ID"`
}

func (Author) TableName() string {
	return "Author"
}

func (c *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s, UpdatedAt : %s", c.ID, c.Name, c.CreatedAt.Format("2006-01-02 15:04:05"), c.UpdatedAt.Format("2006-01-02 15:04:05"))
}