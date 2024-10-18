package main

import (
	"fmt"
	"homework/modul"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var listBuku []modul.Buku
	listBuku = append(listBuku, modul.Buku{"React.js", "Lumoshive", "isbn-1"})
	listBuku = append(listBuku, modul.Buku{"Golang", "Lumoshive", "isbn-2"})
	
	for{
		var input string
		fmt.Println("Menu")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Show All Book")
		fmt.Println("4. Exit")
		fmt.Print("Masukkan nomor menu: ")
		fmt.Scanln(&input)
		if input == "1" {
			ClearScreen()
			title, author, isbn := formBook()
			modul.InterfaceAdd(&modul.Buku{title, author, isbn}, &listBuku)
		}else if input == "2" {
			isPanic := false
			ClearScreen()
			indeks := formRemove(listBuku, &isPanic)
			if isPanic {
				continue
			}
			modul.InterfaceRemove(&modul.Buku{}, indeks, &listBuku)
		}else if input == "3" {
			ClearScreen()
			modul.InterfaceShow(&modul.Buku{}, &listBuku)
		}else if input == "4" {
			break
		}else{
			ClearScreen()
			fmt.Println("Hanya boleh memasukkan angka dari 1 sampai 4")
		}
	}
}
func formRemove(listBuku []modul.Buku, isPanic *bool) int {
	var isbn string
	indeks := -1
	fmt.Print("Masukkan nomor ISBN : ")
	fmt.Scanln(&isbn)
	defer func() {
		if r := recover(); r != nil {
			ClearScreen()
			*isPanic = true
			fmt.Println("Pulih dari panic:", r)
		}
	}()
	isValid := false
	for i, v := range listBuku {
		if v.ISBN == isbn {
			isValid = true
			indeks = i
		}
	}
	if !isValid {
		panic("Buku tidak ditemukan!")
	}
	ClearScreen()
	return indeks
}
func formBook() (string, string, string) {
	var title, author, isbn string
	fmt.Print("Masukkan judul buku: ")
	fmt.Scanln(&title)
	fmt.Print("Masukkan penulis buku: ")
	fmt.Scanln(&author)
	fmt.Print("Masukkan nomor ISBN: ")
	fmt.Scanln(&isbn)
	ClearScreen()
	return title, author, isbn
}
func init()  {
	ClearScreen()
}
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}