package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type record struct {
	UserID    string
	Segment   string
	Operation string
	Date      string
}

//func csv() {
//	//rec := &record{
//	//	UserID:    "14124",
//	//	Segment:   "AVITO_DISCOUNT_1000",
//	//	Operation: "delete",
//	//	Date:      time.Now().Format("2006-01-02 15:04:05"),
//	//}
//	//
//	//WriteCSV(rec)
//
//	date, err := time.Parse("2006-01-02", "2023-08-29")
//
//	if err != nil {
//		log.Fatalf("%v", err)
//	}
//
//	ReadFromCSV("14124", date)
//}

func WriteCSV(recordToCSV *record) {
	file, err := os.OpenFile("usersLogs.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ','
	writer.UseCRLF = true

	//writer.Write([]string{"user_id", "segment", "operation", "date_time"})

	writer.Write([]string{recordToCSV.UserID, recordToCSV.Segment, recordToCSV.Operation, recordToCSV.Date})

	writer.Flush()
}

func ReadFromCSV(userID string, date time.Time) {
	file, err := os.Open("usersLogs.csv")
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		rec := &record{
			UserID:    row[0],
			Segment:   row[1],
			Operation: row[2],
			Date:      row[3],
		}

		if rec.UserID == userID {
			searchDate(date, rec)
		}
	}
}

func searchDate(start time.Time, rec *record) {
	recordDate, err := time.Parse("2006-01-02 15:04:05", rec.Date)
	if err != nil {
		log.Fatal(err)
	}

	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
		if recordDate.Year() == d.Year() && recordDate.Month() == d.Month() && recordDate.Day() == d.Day() {
			fmt.Println(rec)
		}
	}

}
