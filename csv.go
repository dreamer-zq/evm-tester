package tester

import (
	"os"

	"github.com/gocarina/gocsv"
)

// SaveToCSV saves the given data to a CSV file at the specified path.
//
// The function takes two parameters:
// - path: a string representing the directory path where the CSV file will be saved.
// - data: an interface{} representing the data that will be saved to the CSV file.
//
// The function returns an error if any error occurs during the file operations.
func SaveToCSV(path string, data interface{}) error {
	csvFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	if _, err := csvFile.Seek(0, 0); err != nil { // Go to the start of the file
		return err
	}

	if err := gocsv.MarshalFile(data, csvFile); err != nil {
		return err
	}
	return nil
}