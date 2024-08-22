package main

import (
	"fmt"
	"html/template"
	"htmx-learning/config"
	"htmx-learning/filter"
	"htmx-learning/pkg/database"
	"htmx-learning/repository"
	"htmx-learning/services"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xuri/excelize/v2"
)

func main() {
	config := config.NewConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	f := excelize.NewFile()

	defer f.Close()

	// excelSrv := services.NewExportExcel(f)
	// err = excelSrv.ExportExcelDaily(datas)
	// if err != nil {
	// 	panic(err)
	// }

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	// film := map[string][]Film{
	// 	"Films": {
	// 		{Title: "Koe", Director: "Phongphat"},
	// 		{Title: "Minkwan", Director: "Rinlada"},
	// 	},
	// }

	// datas, err := repo.GetDataTempsByBayId(4)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, data := range datas {
	// 	log.Println(data.CurrentPhaseA)
	// }

	filterDaily := filter.SortData{
		Time: false,
	}

	e.GET("/", func(c echo.Context) error {
		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Render the template
		return tmpl.Execute(c.Response().Writer, nil)
	})

	e.GET("/monthly", func(c echo.Context) error {
		tmpl, err := template.ParseFiles("views/monthly.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Render the template
		return tmpl.Execute(c.Response().Writer, nil)
	})

	e.GET("/yearly", func(c echo.Context) error {
		tmpl, err := template.ParseFiles("views/yearly.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Render the template
		return tmpl.Execute(c.Response().Writer, nil)
	})

	// e.GET("/latest-year",func(c echo.Context) error {
	// 	var year int
	// 	date,err := repo.GetMaxDate()
	// 	if err != nil{
	// 		return c.s
	// 	}
	// 	return
	// })
	e.GET("/year-list", func(c echo.Context) error {
		years, err := repo.GetAllYears()
		if err != nil || years == nil {
			return c.String(200, `<li><a class="dropdown-item" href="#" style="color: var(--primary-color);">None</a></li>`)
		}
		htmlRes := ``
		for _, y := range years {
			htmlRes += fmt.Sprintf(`<li><a class="dropdown-item" href="#" style="color: var(--primary-color);">%d</a></li>`, y)
		}
		return c.String(200, htmlRes)
	})

	e.GET("/latest-year", func(c echo.Context) error {
		year := 0
		years, err := repo.GetAllYears()
		if err != nil || years == nil {
			return c.String(200, `None`)
		}
		if len(years) > 1 {
			year = years[1]
		} else {
			year = years[0]
		}
		return c.String(200, fmt.Sprintf(`%d`, year))
	})

	dailyBay := 1
	e.GET("daily-bay", func(c echo.Context) error {
		return c.String(200, fmt.Sprintf(`<span id="dropdownTitle">OUT%d</span>`, dailyBay))
	})
	e.GET("/daily-data", func(c echo.Context) error {

		if c.QueryParam("bay") == "" && c.QueryParam("order") == "" {
			dailyBay = 1
		}

		if c.QueryParam("bay") != "" {
			dailyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		if c.QueryParam("order") != "" {
			filterDaily.Time = !filterDaily.Time
		}

		datas, err := service.GetLatestData(dailyBay, filterDaily)

		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("2006/01/02 15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
											<th scope="row">%s</th>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("2006/01/02 15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})
	filterDayMonthly := filter.SortData{
		Time: false,
	}
	monthlyBay := 1
	e.GET("/day-time-peak", func(c echo.Context) error {
		log.Println("get day")
		if c.QueryParam("bay") == "" && c.QueryParam("order") == "" {
			monthlyBay = 1
		}
		if c.QueryParam("bay") != "" {
			monthlyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		if c.QueryParam("order") != "" {
			filterDayMonthly.Time = !filterDayMonthly.Time
		}
		datas, err := service.GetDataLatestMonthDayTime(monthlyBay, filterDayMonthly)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%s</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("2006/01/02 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
                        <th scope="row">%s</th>
                        <td>%s</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                    </tr>`, data.DataDatetime.Format("01/02/2006 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})

	filterNightMonthly := filter.SortData{
		Time: false,
	}
	e.GET("/night-time-peak", func(c echo.Context) error {
		log.Println("get night")
		if c.QueryParam("bay") != "" {
			monthlyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		log.Println("montly bay = ", monthlyBay)
		if c.QueryParam("order") != "" {
			filterNightMonthly.Time = !filterNightMonthly.Time
		}
		datas, err := service.GetDataLatestMonthNightTime(monthlyBay, filterNightMonthly)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%s</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("2006/01/02 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
                        <th scope="row">%s</th>
                        <td>%s</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                    </tr>`, data.DataDatetime.Format("01/02/2006 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})

	filterAllDayMonthly := filter.SortData{
		Time: false,
	}
	e.GET("/all-time-peak", func(c echo.Context) error {
		if c.QueryParam("bay") != "" {
			monthlyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		if c.QueryParam("order") != "" {
			filterAllDayMonthly.Time = !filterAllDayMonthly.Time
		}
		datas, err := service.GetDataLatestMonthAllTime(monthlyBay, filterAllDayMonthly)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%s</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("2006/01/02 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
                        <th scope="row">%s</th>
                        <td>%s</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                        <td>%.2f</td>
                    </tr>`, data.DataDatetime.Format("01/02/2006 15:04:05"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})

	filterPeekYearly := filter.SortData{
		Time: false,
	}
	yearlyBay := 1
	year := 0
	e.GET("/yearly-peak", func(c echo.Context) error {
		if c.QueryParam("bay") != "" {
			yearlyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		if c.QueryParam("order") != "" {
			filterPeekYearly.Time = !filterPeekYearly.Time
		}

		if c.QueryParam("year") != "" {
			year, _ = strconv.Atoi(c.QueryParam("year"))
		} else {
			years, err := repo.GetAllYears()
			if err != nil || years == nil {
				return c.String(200, ``)
			}
			if len(years) > 1 {
				year = years[1]
			} else {
				year = years[0]
			}
		}

		datas, err := service.GetDataLatestYearPeakTime(yearlyBay, year, filterPeekYearly)
		if err != nil {
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%s</td>
											<td>%s</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("January"), data.DataDatetime.Format("02"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
		                <th scope="row">%s</th>
						<td>%s</td>
						<td>%s</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
		            </tr>`, data.DataDatetime.Format("January"), data.DataDatetime.Format("02"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})
	filterLightYearly := filter.SortData{
		Time: false,
	}
	e.GET("/yearly-light", func(c echo.Context) error {
		if c.QueryParam("bay") != "" {
			yearlyBay, _ = strconv.Atoi(c.QueryParam("bay"))
		}
		if c.QueryParam("order") != "" {
			filterLightYearly.Time = !filterLightYearly.Time
		}
		if c.QueryParam("year") != "" {
			year, _ = strconv.Atoi(c.QueryParam("year"))
		} else {
			years, err := repo.GetAllYears()
			if err != nil || years == nil {
				return c.String(200, ``)
			}
			if len(years) > 1 {
				year = years[1]
			} else {
				year = years[0]
			}
		}

		datas, err := service.GetDataLatestYearLightTime(yearlyBay, year, filterLightYearly)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		responeHtml := ``
		for i, data := range datas {
			if (i+1)%2 == 0 {
				responeHtml += fmt.Sprintf(`<tr class="even-row">
											<th scope="row">%s</th>
											<td>%s</td>
											<td>%s</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
											<td>%.2f</td>
										</tr>`, data.DataDatetime.Format("January"), data.DataDatetime.Format("02"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			} else {
				responeHtml += fmt.Sprintf(`<tr class="odd-row">
                        <th scope="row">%s</th>
						<td>%s</td>
						<td>%s</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
						<td>%.2f</td>
                    </tr>`, data.DataDatetime.Format("January"), data.DataDatetime.Format("02"), data.DataDatetime.Format("15:04:05"), 0.00, 0.00, 0.00, data.CurrentPhaseA, data.CurrentPhaseB, data.CurrentPhaseC, data.ActivePower, data.ReactivePower, data.PowerFactor)
			}
		}

		return c.String(200, responeHtml)
	})
	e.GET("/export-pdf-daily", func(c echo.Context) error {

		DeleteFile()

		datas, err := service.GetLatestData(dailyBay, filterDaily)

		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		fileName := fmt.Sprintf("%d.pdf", time.Now().Unix())
		err = services.ExportPdfDaily(datas, fileName)
		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		//time.Sleep(1 * time.Second)
		return c.File(fileName)
	})

	e.GET("/export-xlsx-daily", func(c echo.Context) error {

		DeleteFile()
		datas, err := service.GetLatestData(dailyBay, filterDaily)

		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}
		fileName := fmt.Sprintf("%d.xlsx", time.Now().Unix())
		excelSrv := services.NewExportExcel(f)
		err = excelSrv.ExportExcelDaily(datas, fileName)

		if err != nil {
			log.Println("err:", err.Error())
			return c.String(200, ``)
		}

		// c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
		// c.Response().Header().Set("Content-Disposition", "inline; filename="+fileName)
		//time.Sleep(1 * time.Second)
		return c.File(fileName)
	})

	e.Logger.Printf("listening on port:", config.HTTP_PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))

}

func DeleteFile() {
	pdfFiles, err := filepath.Glob(filepath.Join("", "*.pdf"))
	if err != nil {
		log.Println(err)
	}

	// Iterate over the list of files and delete each one
	for _, file := range pdfFiles {
		err := os.Remove(file)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted: %s\n", file)
	}
	xlsxFiles, err := filepath.Glob(filepath.Join("", "*.xlsx"))
	if err != nil {
		log.Println(err)
	}
	for _, file := range xlsxFiles {
		err := os.Remove(file)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted: %s\n", file)
	}
}
