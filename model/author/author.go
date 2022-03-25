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
	Books		[]book.Book `gorm:"foreignKey:AuthorID;references:ID"`
}

func (Author) TableName() string {
	return "Author"
}

func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s", a.ID, a.Name, a.CreatedAt.Format("2006-01-02 15:04:05"))
}