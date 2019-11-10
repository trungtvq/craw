package craw

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var err error
var page int64
var maxpage int64
var xlsx *excelize.File
var currentRow int64

func get(page int64) string {
	client := &http.Client{}
	var addr string
	if page == 1 {
		addr = "https://admin.tuyendungsieuviet.com/employer"
	} else {
		addr = "https://admin.tuyendungsieuviet.com/employer?page=" + strconv.FormatInt(page, 10)
	}
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	req.Header.Add("authority", "admin.tuyendungsieuviet.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("cookie", "PHPSESSID=llp7uv79kf32mospqngjqic8b4")

	resp, err := client.Do(req)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}

func GetEmail(strs string) {
	l := len(strs)
	for i := 0; i < l-14; i++ {
		if strs[i:i+14] == "Tài Khoản: " {
			email := ""
			i = i + 14
			for strs[i] != 32 {
				email = email + string(strs[i])
				i = i + 1
			}

			xlsx.SetCellValue("Sheet1", "B"+strconv.FormatInt(currentRow, 10), email)

			GetPhone(strs[i:])
			return
		}
		if strs[i:i+8] == "</tbody>" {
			IncreasePage()
			if maxpage > page {
				GetName(get(page))
				return
			}
		}
	}
}
func GetPhone(strs string) {

	l := len(strs)

	for i := 0; i < l; i++ {
		if strs[i:i+20] == "Điện thoại: <b>" {
			phone := ""
			i = i + 20
			for strs[i] != 60 {
				phone = phone + string(strs[i])
				i = i + 1
			}

			xlsx.SetCellValue("Sheet1", "C"+strconv.FormatInt(currentRow, 10), phone)

			GetAddress(strs[i:])
			return
		}
		if strs[i:i+8] == "</tbody>" {
			IncreasePage()
			if maxpage > page {
				GetName(get(page))
				return
			}
		}
	}
}
func IncreasePage() {
	page = page + 1
	fmt.Println("Processing in page: " + strconv.FormatInt(page, 10))
}
func GetAddress(strs string) {

	l := len(strs)
	for i := 0; i < l; i++ {
		if strs[i:i+14] == "Địa chỉ: " {
			address := ""
			i = i + 14
			j := i
			for strs[j] != 60 {
				j = j + 1
			}
			address = strs[i:j]
			i = j
			xlsx.SetCellValue("Sheet1", "D"+strconv.FormatInt(currentRow, 10), address)
			GetName(strs[i:])
			return
		}
		if strs[i:i+8] == "</tbody>" {
			IncreasePage()
			if maxpage > page {
				GetName(get(page))
				return
			}
		}
	}
}
func GetName(strs string) {

	l := len(strs)
	for i := 0; i < l-31; i++ {
		if strs[i:i+31] == `                            <b>` {
			name := ""
			i = i + 31
			j := i
			for strs[j] != 60 {
				j = j + 1
			}
			name = strs[i:j]
			i = j
			xlsx.SetCellValue("Sheet1", "A"+strconv.FormatInt(currentRow, 10), name)
			currentRow = currentRow + 1
			GetEmail(strs[i:])
			return
		}
		if strs[i:i+8] == "</tbody>" {
			IncreasePage()
			if maxpage > page {
				GetName(get(page))
				return
			}
		}
	}
	if maxpage > page {
		GetName(get(page))
		return
	}
}
func run() {
	fmt.Println("Input number of page need craw:")
	fmt.Println("---------------------")
	fmt.Scanf("%d", &maxpage)
	page = 1
	xlsx = excelize.NewFile()
	currentRow = 2
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet1")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A1", "Name")
	xlsx.SetCellValue("Sheet1", "B1", "Email")
	xlsx.SetCellValue("Sheet1", "C1", "Phone")
	xlsx.SetCellValue("Sheet1", "D1", "Address")
	GetName(get(page))
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

