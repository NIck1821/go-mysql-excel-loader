package xls_loader

import (
	"fmt"
	"sync"

	"github.com/Nick1821/go-mysql-excel-loader/internal/app/model"
	"github.com/xuri/excelize/v2"
)

func Loader_XLS(leads []model.Lead) error {
	file := excelize.NewFile()
	// Create a new sheet.
	index := file.NewSheet("Sheet1")

	file.SetCellValue("Sheet1", "A1", "firstname")
	file.SetCellValue("Sheet1", "B1", "lastname")
	file.SetCellValue("Sheet1", "C1", "test")
	file.SetCellValue("Sheet1", "D1", "phone")
	file.SetCellValue("Sheet1", "E1", "city")
	// Set value of a cell.
	columns := []string{"A", "B", "C", "D", "E"}

	wg := sync.WaitGroup{}
	for _, column := range columns {
		wg.Add(1)
		go func(column string) {
			var row string
			for i := 0; i < len(leads); i++ {
				switch column {
				case "A":
					row = leads[i].FirstName
				case "B":
					row = leads[i].LastName
				case "C":
					row = leads[i].Text
				case "D":
					row = leads[i].Phone
				case "E":
					row = leads[i].City
				}
				file.SetCellValue("Sheet1", column+fmt.Sprint(i+2), row)
			}
			wg.Done()
		}(column)
	}
	wg.Wait()
	// Set active sheet of the workbook.
	file.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := file.SaveAs("exmpl.xlsx"); err != nil {
		return err
	}
	return nil
}
