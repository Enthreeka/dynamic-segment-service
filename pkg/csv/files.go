package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//type Record struct {
//	UserID    string
//	Segment   string
//	Operation string
//	Date      string
//}
//
////func csv() {
////	//rec := &record{
////	//	UserID:    "14124",
////	//	Segment:   "AVITO_DISCOUNT_1000",
////	//	Operation: "delete",
////	//	Date:      time.Now().Format("2006-01-02 15:04:05"),
////	//}
////	//
////	//WriteCSV(rec)
////
////	date, err := time.Parse("2006-01-02", "2023-08-29")
////
////	if err != nil {
////		log.Fatalf("%v", err)
////	}
////
////	ReadFromCSV("14124", date)
////}
//
////func NewCSVFile(log *logger.Logger) *Record {
////	return &Record{
////		log: log,
////	}
////}
//
//func (r *Record) createFile(fileName string) (*os.File, error) {
//	file, err := os.Create(fileName)
//	if err != nil {
//		return nil, err
//	}
//
//	return file, nil
//}
//
//func fileName(userID string, segment string) string {
//	var builder strings.Builder
//
//	builder.Write([]byte(userID))
//	builder.Write([]byte("_"))
//	builder.Write([]byte(segment))
//	builder.Write([]byte(".csv"))
//
//	return builder.String()
//}
//
//func (r *Record) writeNewFile(userID string, segment string, rec *Record) *os.File {
//
//	name := fileName(userID, segment)
//	file, err := r.createFile(name)
//	if err != nil {
//		log.Errorf("%v", err)
//	}
//
//	//	defer file.Close()
//
//	writer := csv.NewWriter(file)
//	writer.Comma = ','
//	writer.UseCRLF = true
//
//	writer.Write([]string{"user_id", "segment", "operation", "date_time"})
//
//	writer.Flush()
//
//	return file
//}
//
//func (r *Record) Write() {
//	file, err := os.OpenFile("usersLogs.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Errorf("%v", err)
//	}
//
//	defer file.Close()
//
//	writer := csv.NewWriter(file)
//	writer.Comma = ','
//	writer.UseCRLF = true
//
//	writer.Write([]string{r.UserID, r.Segment, r.Operation, r.Date})
//
//	writer.Flush()
//}
//
//func (r *Record) Read(userID string, segment string, date time.Time) {
//
//	file, err := os.Open("usersLogs.csv")
//	if err != nil {
//		log.Errorf("%v", err)
//	}
//
//	defer file.Close()
//
//	reader := csv.NewReader(file)
//	if err != nil {
//		log.Errorf("%v", err)
//	}
//
//	for {
//		row, err := reader.Read()
//
//		if err == io.EOF {
//			break
//		}
//
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//
//		rec := &Record{
//			UserID:    row[0],
//			Segment:   row[1],
//			Operation: row[2],
//			Date:      row[3],
//		}
//
//		if rec.UserID == userID && rec.Segment == segment {
//			r.searchDate(date, rec)
//		}
//	}
//}
//
//func (r *Record) searchDate(start time.Time, rec *Record) {
//	recordDate, err := time.Parse("2006-01-02 15:04:05", rec.Date)
//	if err != nil {
//		log.Errorf("%v", err)
//	}
//
//	file := r.writeNewFile(rec.UserID, rec.Segment, rec)
//
//	defer file.Close()
//
//	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
//		if recordDate.Year() == d.Year() && recordDate.Month() == d.Month() && recordDate.Day() == d.Day() {
//
//			writer := csv.NewWriter(file)
//			writer.Write([]string{rec.UserID, rec.Segment, rec.Operation, rec.Date})
//			writer.Flush()
//		}
//	}
//
//}

type Record struct {
	UserID    string
	Segment   string
	Operation string
	Date      string
}

func (r *Record) Write() {
	file, err := os.OpenFile("usersLogs.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening main file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ','
	writer.UseCRLF = true

	writer.Write([]string{r.UserID, r.Segment, r.Operation, r.Date})

	writer.Flush()
}

func (r *Record) Read(userID string, segment string, date time.Time) {
	inputFile, err := os.Open("usersLogs.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer inputFile.Close()

	outputFileName := fileName(userID, segment, date)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	inputReader := csv.NewReader(inputFile)
	outputWriter := csv.NewWriter(outputFile)

	header, err := inputReader.Read()
	if err != nil {
		log.Fatalf("Error reading header from input file: %v", err)
	}
	outputWriter.Write(header)
	outputWriter.Flush()

	for {
		row, err := inputReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading from input file: %v", err)
			break
		}

		rec := &Record{
			UserID:    row[0],
			Segment:   row[1],
			Operation: row[2],
			Date:      row[3],
		}

		recordDate, err := time.Parse("2006-01-02 15:04:05", rec.Date)
		if err != nil {
			log.Printf("Error parsing record date: %v", err)
			continue
		}

		if rec.UserID == userID && rec.Segment == segment && recordDate.Year() == date.Year() && recordDate.Month() == date.Month() && recordDate.Day() == date.Day() {
			outputWriter.Write(row)
			outputWriter.Flush()
		}
	}
}

func fileName(userID string, segment string, date time.Time) string {
	return fmt.Sprintf("%s_%s_%s.csv", userID, segment, date.Format("2006-01-02"))
}
