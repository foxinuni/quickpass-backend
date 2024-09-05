package main

import (
	"os"

	"github.com/foxinuni/quickpass-backend/internal/domain/services"
)

func main() {
	// open file
	reader, err := os.Open("test.xlsx")
	if err != nil {
		panic(err)
	}

	// Create service and test
	importService := services.NewExcelImportService()
	importService.ImportFromFile(reader)
}
