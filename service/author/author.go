package author

import (
	"gorm.io/gorm"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/model/author"
)

type AuthorService struct {
	db *gorm.DB
}

func NewAuthorService(db *gorm.DB) *AuthorService {
	return &AuthorService{db: db}
}

func (a *AuthorService) Migration() {
	a.db.AutoMigrate(&author.Author{})
}

func (a *AuthorService) InsertSampleData() {
	authors := []author.Author{
		{Name: "Author1"},
		{Name: "Author2"},
	}

	for _, author := range authors {
		a.db.Create(&author)
	}
}