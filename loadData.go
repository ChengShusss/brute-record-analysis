package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Date       string
	Time       string
	RemoteAddr string
	Client     string
	User       string
	Password   string
}

var (
	db *gorm.DB

	CntInvalidLine int = 0
	CntValidLine   int = 0
)

func main() {

	dbFile := flag.String("db", "test.db", "sqlite db file name")
	flag.Parse()
	err := openDb(*dbFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = db.AutoMigrate(&Record{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: loadData [-db dbFilePath] <file>")
		return
	}

	fileName := flag.Args()[0]
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		processLine(line)
		// process one line for record
	}

	fmt.Printf("Total line: %d, Invalid line: %d, Valid line: %d", CntInvalidLine+CntValidLine, CntInvalidLine, CntValidLine)
}

func openDb(fileName string) (err error) {
	db, err = gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	return
}

func processLine(line string) {
	fields := strings.Split(line, " ")
	if len(fields) != 6 {
		CntInvalidLine += 1
		return
	}
	record := &Record{
		Date:       fields[0],
		Time:       fields[1],
		RemoteAddr: fields[2],
		Client:     fields[3],
		User:       fields[4],
		Password:   fields[5],
	}

	db.Create(&record)
	CntValidLine += 1
}
