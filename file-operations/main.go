package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var filename = "patikadev.txt"

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

err := ReadFileWords()
if err != nil{
	log.Fatal(err)
}

}

func CreateEmptyFile() (*os.File, error){
	myFile, err := os.Create(filename)
	if err != nil{
		return nil, err
	}
	log.Println("File created", myFile.Name())
	return myFile, nil
}

func WriteFile(file *os.File){
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString("patika dev")
	w.Flush()
}

func GetFileInfo() error {
	fileInfo, err := os.Stat(filename)
	if err != nil{
		return err
	}

	fmt.Println("File name", fileInfo.Name())
	return nil
}

func ReadFileLines() error {
	f, err := os.Open(filename)
	if err != nil{
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		fmt.Printf("line %s\n", scanner.Text())
	}

	if err := scanner.Err(); err!=nil{
		return err
	}
	return nil
}

func ReadFileWords() error {
	f, err := os.Open(filename)
	if err != nil{
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		fmt.Printf("line %s\n", scanner.Text())
	}

	if err := scanner.Err(); err!=nil{
		return err
	}
	return nil
}