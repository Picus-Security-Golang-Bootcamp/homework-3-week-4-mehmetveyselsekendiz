package book

import (
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string
	Page        int
	Stock       int
	Price       float32   `sql:"type:decimal(10,2);"`
	StockCode   string
	ISBN        string
	AuthorID    int
}

func (Book) TableName() string {
	return "Book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Page : %d, Stock : %d, Price : %g, StockCode : %s, ISNB : %s, AuthorID : %d, CreatedAt : %s",
	 b.ID, b.Name, b.Page, b.Stock, b.Price, b.StockCode, b.ISBN, b.AuthorID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}