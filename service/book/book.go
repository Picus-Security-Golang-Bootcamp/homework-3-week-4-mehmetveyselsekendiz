package book

import (
	"gorm.io/gorm"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/model/book"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		db: db,
	}
}

func (b *BookService) Migration() {
	b.db.AutoMigrate(&book.Book{})
}

func (b *BookService) InsertSampleData() {
	books := []book.Book{
		{Name: "Book1", Page: 11, Stock: 111, Price: 11.11, StockCode: "B1", ISBN: "ISBN-B1", AuthorID: 1},
		{Name: "Book2", Page: 22, Stock: 222, Price: 22.22, StockCode: "B2", ISBN: "ISBN-B2", AuthorID: 2},
	}

	for _, book := range books {
		b.db.Create(&book)
	}
}