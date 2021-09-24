package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var menu int
var ID int = 0
var title, name, saver string
var hour, billing float64
var data = [][4]string{{"ID", "name", "hour", "billing (IDR)"}}

func max() string {
	f := math.MaxFloat64
	return strconv.FormatFloat(f, 'f', 1, 64)
}

func add() {
	fmt.Println("Masukan Nama : ")
	fmt.Scan(&name)
	ID++
	fmt.Println("Masukan Jam : ")
	fmt.Scan(&hour)
	saver = strconv.FormatFloat(hour, 'f', 1, 64)
	billing = hour * 60 * 1000
	if valid, err := validate(saver); valid {
		newdata := [4]string{(strconv.Itoa(ID)), name, (strconv.FormatFloat(hour, 'f', 1, 64)), (strconv.FormatFloat(billing, 'f', 1, 64))}
		data = append(data, newdata)
		newdata = [4]string{}
	} else {
		fmt.Println(err.Error())
	}
}
func showAll() {
	fmt.Println("Daftar nama yang tersedia: \n", len(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(data[i][j], "  \t")
		}
		fmt.Println("")
	}
}

func searchID(ID int) int {
	var index int
	search := strconv.Itoa(ID)
	for i := 0; i < len(data); i++ {
		if data[i][0] == search {
			index = i
			break
		}
	}
	return index
}

func deleteID(ID int) {
	var before, after [][4]string
	index := searchID(ID)
	before = data[:index]
	after = data[index+1:]
	data = [][4]string{}
	for i := 0; i < len(before); i++ {
		data = append(data, before[i])
	}
	for i := 0; i < len(after); i++ {
		data = append(data, after[i])
	}
}

func searchname() {
	fmt.Println("Search result(s) : ")
	for i := 0; i < len(data); i++ {
		matched, _ := regexp.MatchString(`^a|^A`, data[i][2])
		if matched == true {
			fmt.Println("ID : ", data[i][0])
			fmt.Println("Title : ", data[i][1])
			fmt.Println("name : ", data[i][2])
			fmt.Println("Vote : ", data[i][3])
		}
	}
}
func stringtoFloat(input string) float64 {
	var save float64
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		save = s
	}
	return save
}
func topthree() {
	if len(data) < 4 {
		fmt.Println("Maaf datanya terlalu pendek, tolong tambahkan lagi")
	} else {
		var max string = max()
		// third := [4]string{"99999", "99999", "99999", "99999"}
		// second := [4]string{"99999", "99999", "99999", "99999"}
		// first := [4]string{"99999", "99999", "99999", "99999"}
		third := [4]string{max, max, max, max}
		second := [4]string{max, max, max, max}
		first := [4]string{max, max, max, max}
		for i := 1; i < len(data); i++ {
			if stringtoFloat(data[i][2]) < stringtoFloat(first[2]) {
				third = second
				second = first
				first = data[i]
			} else if stringtoFloat(data[i][2]) < stringtoFloat(second[2]) {
				third = second
				second = data[i]
			} else if stringtoFloat(data[i][2]) < stringtoFloat(third[2]) {
				third = data[i]
			}
		}
		fmt.Print("Top 3 jam paling sedikit teratas : \n")
		fmt.Print(data[0][0], "\t", data[0][1], "\t", data[0][2], "\t", data[0][3], "\n")
		fmt.Print(first[0], "\t", first[1], "\t", first[2], "\t", first[3], "\n")
		fmt.Print(second[0], "\t", second[1], "\t", second[2], "\t", second[3], "\n")
		fmt.Print(third[0], "\t", third[1], "\t", third[2], "\t", third[3], "\n")
	}
}
func countAverage() float64 {
	var count float64 = 0
	var sum float64 = 0
	for i := 1; i < len(data); i++ {
		sum = sum + stringtoFloat(data[i][2])
		count++
	}
	average := sum / count
	return average
}

func onlyfour() {
	average := countAverage()
	fmt.Println("Daftar billing perjam < ", average)
	for i := 1; i < len(data); i++ {
		if stringtoFloat(data[i][2]) < average {
			fmt.Print(data[i][0], "\t", data[i][1], "\t", data[i][2], "\t", data[i][3], "\n")
		}
	}
}
func validate(input string) (bool, error) { //bug, don't know why doesn't work
	if input == "" {
		return false, errors.New("Tidak boleh kosong")
	}
	m := regexp.MustCompile("[0-9]")
	if m.MatchString(input) == false {
		return false, errors.New("Silahkan masukan nomor saja")
	}
	return true, nil
}

func main() {
	fmt.Println("========================")
	fmt.Println("Daftar Pilihan :")
	fmt.Println("1. Tambah billing baru")
	fmt.Println("2. Hapus billing")
	fmt.Println("3. Tampilkan  billing")
	fmt.Println("4. Rata-rata jam")
	fmt.Println("5. Menampilkan 3 buah data yang paling sedikit")
	fmt.Println("6. Menampilkan seluruh data customer yang menyewa komputer kurang dari rata – rata")
	fmt.Println("0. Keluar")
	fmt.Println("========================")
	fmt.Print("Pilih Menu angka diatas [0...6] : ")
	fmt.Scan(&menu)
	if menu == 1 {
		add()
		main()
	} else if menu == 2 {
		fmt.Println("ID apa yang ingin Anda hapus? : ")
		fmt.Scan(&saver)
		var num, err = strconv.Atoi(saver)
		if err == nil && num > 0 {
			deleteID(num)
		}
		showAll()
		main()
	} else if menu == 3 {
		showAll()
		main()
	} else if menu == 4 {
		fmt.Println("Rata-rata jam yang digunakan : ", countAverage())
		main()
	} else if menu == 5 {
		topthree()
		main()
	} else if menu == 6 {
		onlyfour()
		main()
	} else if menu == 7 {
		fmt.Println("Thanks")
		os.Exit(1)
	}

}