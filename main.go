package main

import(
	"log"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/file_operation"
	postgres "github.com/mehmetveyselsekendiz/homework-3-week-4/common/db"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/service/author"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/service/book"
	"github.com/joho/godotenv"
)

func main(){

	//err := ReadFileLines()
	//if err != nil{
	//	log.Fatal(err)
	//}

	err := file_operation.ReadFileWords()
	if err != nil{
		log.Fatal(err)
	}

	//Set environment variables
	godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	// Repositories
	bookService := book.NewBookService(db)
	bookService.Migration()
	bookService.InsertSampleData()

	authorService := author.NewAuthorService(db)
	authorService.Migration()
	authorService.InsertSampleData()
}