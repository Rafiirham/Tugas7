package main

import (
	"fmt"
	"sort"
)

var increment int = 1
var ID_SINGER = make(OrderList, len(VOTE_SINGER))
var TITLE_SINGER = map[int]string{}
var NAME_SINGER = map[int]string{}
var VOTE_SINGER = map[int]int{}

func main() {
	// Risky Kurniawan - ARS University
MenuUtama:
	var menu, submenu string

	fmt.Println("-------APLIKASI VOTE PENYANYI-------")
	fmt.Println("Silahkan pilih menu dibawah ini : ")
	fmt.Println("1. Tampil Data Voting")
	fmt.Println("2. Input Data Vote")
	fmt.Println("3. Hapus Data Vote")
	fmt.Println("4. TOP 3 Musik Terfavorit")
	fmt.Println("5. Penyanyi awalan huruf 'A'")
	fmt.Println("0. Keluar Aplikasi")
	fmt.Println("-------------------------------------")
	fmt.Print("Pilihan anda adalah : ")
	fmt.Scanf("%s\n", &menu)
	fmt.Println()

Pil:
	if menu == "0" {
		fmt.Println("----------------------------")
		fmt.Println("NUR EVATUL NISA")
		fmt.Println("----------------------------")
	} else if menu == "1" {
		SelectData()
	} else if menu == "2" {
		InputData()
	} else if menu == "3" {
		DeleteData()
	} else if menu == "4" {
		Top3Data()
	} else if menu == "5" {
		ASinger()
	} else {
		fmt.Println("-------Menu Tidak Tersedia-------")
		fmt.Println("---Silahkan pilih menu kembali---")
	}

PilSub:

	if menu != "0" {
		fmt.Println("9. Kembali ke menu")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilihan anda adalah : ")
		fmt.Scanf("%s\n", &submenu)

		if submenu == "9" {
			goto MenuUtama
		} else if submenu == "0" {
			menu = "0"
			goto Pil
		} else {
			fmt.Println("-------Menu Tidak Tersedia-------")
			fmt.Println("---Silahkan pilih menu kembali---")
			goto PilSub
		}
	}
}

type Order struct {
	Key   int
	Value int
}
type OrderList []Order

func (p OrderList) Len() int           { return len(p) }
func (p OrderList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p OrderList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func Ordering() {
	ID_SINGER = make(OrderList, len(VOTE_SINGER))

	i := 0
	for k, v := range VOTE_SINGER {
		ID_SINGER[i] = Order{k, v}
		i++
	}

	sort.Sort(ID_SINGER)
}

func SelectData() {
	fmt.Println("-------TAMPIL DATA VOTING PENYANYI-------")
	var sum_vote int = 0
	for _, ID := range ID_SINGER {
		sum_vote += VOTE_SINGER[ID.Key]
		fmt.Printf("ID:%d -> [%s - %s] VOTE(%d) \n", ID.Key, TITLE_SINGER[ID.Key], NAME_SINGER[ID.Key], VOTE_SINGER[ID.Key])
	}
	if len(TITLE_SINGER) == 0 {
		fmt.Println("          NO RECORD DATA")
	}
	fmt.Println("Jumlah Record =", len(TITLE_SINGER))
	fmt.Println("Jumlah Vote   =", sum_vote)
	fmt.Println("------------------------------------------")
}

func InputData() {
	var title, name string = "", ""
	var notif, err string = "", ""
	var vote int
	fmt.Println("-------INPUT DATA VOTING PENYANYI-------")
	fmt.Print("Masukan Judul Lagu    : ")
	fmt.Scanf("%s\n", &title)
	fmt.Print("Masukan Nama Penyanyi : ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Masukan Jumlah Vote   : ")
	fmt.Scanf("%d\n", &vote)

	_, err = validation(title, "title")
	notif += err
	_, err = validation(name, "name")
	notif += err

	if notif == "" {
		TITLE_SINGER[increment] = title
		NAME_SINGER[increment] = name
		VOTE_SINGER[increment] = vote
		increment++
		Ordering()
		fmt.Println("            INPUT BERHASIL")
	} else {
		fmt.Println("NOTIF: ", notif)
		fmt.Println("-------INPUT GAGAL-------")
	}

	fmt.Println("-----------------------------------")
}

func validation(value string, name string) (string, string) {
	if value == "" {
		return value, fmt.Sprintf("required %s;", name)
	} else {
		return "", ""
	}
}

func DeleteData() {
	var id int
	fmt.Println("-------HAPUS DATA VOTING PENYANYI-------")
	fmt.Print("Masukan ID Data : ")
	fmt.Scanf("%d\n", &id)
	fmt.Println("----------------------------------------")
	if TITLE_SINGER[id] != "" {
		delete(TITLE_SINGER, id)
		delete(NAME_SINGER, id)
		delete(VOTE_SINGER, id)
		Ordering()
		fmt.Println("        DATA BERHASIL DIHAPUS")
	} else {
		fmt.Println("        ID Tidak Ditemukan")
	}
	fmt.Println("----------------------------------------")
}

func ASinger() {
	fmt.Println("-------DATA VOTING PENYANYI 'A'-------")
	var exist bool = false
	for _, ID := range ID_SINGER {
		if string(NAME_SINGER[ID.Key][0]) == "A" || string(NAME_SINGER[ID.Key][0]) == "a" {
			fmt.Printf("ID:%d -> [%s - %s] VOTE(%d) \n", ID.Key, TITLE_SINGER[ID.Key], NAME_SINGER[ID.Key], VOTE_SINGER[ID.Key])
			exist = true
		}
	}
	if exist == false {
		fmt.Println("-------NO RECORD DATA-------")
	}
	fmt.Println("----------------------------")
}
func Top3Data() {
	fmt.Println("-------TOP 3 DATA VOTING PENYANYI-------")
	var counter int = 0
	for _, ID := range ID_SINGER {
		counter++
		fmt.Printf("ID:%d -> [%s - %s] VOTE(%d) \n", ID.Key, TITLE_SINGER[ID.Key], NAME_SINGER[ID.Key], VOTE_SINGER[ID.Key])
		if counter == 3 {
			break
		}
	}
	if len(TITLE_SINGER) == 0 {
		fmt.Println("          NO RECORD DATA")
	}
	fmt.Println("------------------------------------------")
}
