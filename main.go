package main

import(
	"log"
	"github.com/mehmetveyselsekendiz/homework-3-week-4/file_operation"
)

func main(){
	//	fmt.Println("Hi")

	//	_, err := CreateEmptyFile()
	//	if err != nil{
	//		log.Fatal(err)
	//	}

	//	file, err := os.OpenFile(filename,os.O_RDWR, 8755)
	//	if err != nil{
	//		log.Fatal(err)
	//	}
	//	WriteFile(file)

	//err := GetFileInfo()
	//if err != nil{
	//	log.Fatal(err)
	//}

	//err := ReadFileLines()
	//if err != nil{
	//	log.Fatal(err)
	//}

	err := file_operation.ReadFileWords()
	if err != nil{
		log.Fatal(err)
	}
}