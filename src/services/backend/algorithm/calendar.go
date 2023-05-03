package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

func Calendar(dateStr string) {
	date := strings.Split(dateStr, "/")

	// validasi nilai negatif tanggal
	day, err := strconv.Atoi(date[0])
	if err != nil || day < 1 {
		fmt.Println("Tanggal tidak valid")
		return
	}

	// validasi bulan
	month, err := strconv.Atoi(date[1])
	if err != nil || month < 1 || month > 12 {
		fmt.Println("Bulan tidak valid")
		return
	}

	// validasi tahun
	year, err := strconv.Atoi(date[2])
	if err != nil {
		fmt.Println("Tahun tidak valid")
		return
	}

	// validasi tahun kabisat
	isLeapYear := (year%4 == 0 && year%100 != 0) || year%400 == 0

	// validasi tanggal sesuai dengan bulan
	maxDay := 31
	if month == 2 {
		if isLeapYear {
			maxDay = 29
		} else {
			maxDay = 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		maxDay = 30
	}

	if day > maxDay {
		fmt.Println("Tanggal tidak valid")
		return
	}

	// hitung hari dalam minggu
	daysOfWeek := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4} // tabel konstanta
	y := year
	if month < 3 {
		y--
	}
	dayOfWeekIndex := (y + y/4 - y/100 + y/400 + t[month-1] + day) % 7

	// cetak hari dalam minggu
	fmt.Println(daysOfWeek[dayOfWeekIndex])
}
