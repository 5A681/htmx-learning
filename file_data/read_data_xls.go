package filedata

import (
	"htmx-learning/entity"
	"htmx-learning/repository"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/extrame/xls"
)

func ReadFileXls(filepath string, sheet int, wg *sync.WaitGroup, dataTempRepo repository.Repository) {
	defer wg.Done()
	xlsFile, err := xls.Open(filepath, "utf-8")
	if err != nil {
		log.Fatalf("Failed to open .xls file: %v", err)
	}
	_ = xlsFile
	ws := xlsFile.GetSheet(sheet)
	maxRow := int(ws.MaxRow)
	maxCol := ws.Row(0).LastCol()

	for r := 5; r < maxRow; r++ {

		if ws.Row(r).Col(0) != "" {
			tempData := entity.DataTmps{
				BayId: sheet - 2,
			}
			for c := 0; c < maxCol; c++ {

				if c == 0 {

					//fmt.Printf("%v \t", ReadDateTimeColumn(ws.Row(r).Col(c)))
					dateTime := ReadDateTimeColumn(ws.Row(r).Col(c))
					if dateTime != nil {
						tempData.DataDatetime = *dateTime
					}
				}
				if c == 2 || c == 4 || c == 6 || c == 8 || c == 10 || c == 12 {
					//fmt.Printf("%v \t", ws.Row(r).Col(c))
					if c > 12 {
						continue
					} else if c == 2 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.CurrentPhaseA = 0
						} else {
							tempData.CurrentPhaseA = float32(floatData)
						}
					} else if c == 4 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.CurrentPhaseB = 0
						} else {
							tempData.CurrentPhaseB = float32(floatData)
						}
					} else if c == 6 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.CurrentPhaseC = 0
						} else {
							tempData.CurrentPhaseC = float32(floatData)
						}
					} else if c == 8 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.ActivePower = 0
						} else {
							tempData.ActivePower = float32(floatData)
						}
					} else if c == 10 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.ReactivePower = 0
						} else {
							tempData.ReactivePower = float32(floatData)
						}
					} else if c == 12 {
						floatData, err := strconv.ParseFloat(ws.Row(r).Col(c), 64)
						if err != nil {
							tempData.PowerFactor = 0
						} else {
							tempData.PowerFactor = float32(floatData)
						}
					}
				} else {
					continue
				}

			}
			tempData.CreatedAt = time.Now()

			err = dataTempRepo.CreateDataTmep(tempData)
			if err != nil {
				log.Println("could not insrt temp data", err.Error())
			}
		}

		//fmt.Printf("\n\n")
	}
}

func ReadDateTimeColumn(datetime string) *time.Time {
	var realTime *time.Time
	t, err := strconv.ParseFloat(datetime, 64)
	if err != nil {
		convertTime, err := time.Parse(time.RFC3339, datetime)
		if err != nil {
			return nil
		} else {
			if convertTime.Second() != 0 {
				convertTime = convertTime.Add(time.Second * (time.Duration(60 - convertTime.Second())))
			}
			realTime = &convertTime
		}

	} else {
		//realTime =
		convertTime := ExcelDateToTime(t)
		if convertTime.Second() != 0 {
			convertTime = convertTime.Add(time.Second * (time.Duration(60 - convertTime.Second())))
		}
		realTime = &convertTime
	}

	//}
	return realTime
}

// func ReadCurrentA() float32 {

// }

func ExcelDateToTime(serialDate float64) time.Time {
	// Excel uses 30-Dec-1899 as the epoch (0 date)
	excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

	// Convert the serial number to a duration
	days := int(serialDate)
	seconds := (serialDate - float64(days)) * 86400 // seconds in a day

	// Add the days and seconds to the epoch
	finalTime := excelEpoch.AddDate(0, 0, days).Add(time.Duration(seconds) * time.Second)

	return finalTime
}
