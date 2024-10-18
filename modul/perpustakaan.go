package modul

import "fmt"

type Buku struct {
	Title  string
	Author string
	ISBN   string
}
type Perpustakaan interface {
	AddBook(listBuku *[]Buku)
	RemoveBook(indeks int, listBuku *[]Buku)
	ShowBooks(listBuku *[]Buku)
}

func (buku *Buku) AddBook(listBuku *[]Buku) {
	*listBuku = append(*listBuku, *buku)
	fmt.Println("Buku berhasil ditambahkan!")
}
func (buku Buku) RemoveBook(indeks int, listBuku *[]Buku) {
	*listBuku = append((*listBuku)[:indeks], (*listBuku)[indeks+1:]...)
	fmt.Println("Buku berhasil dihapus!")
}
func (buku Buku) ShowBooks(listBuku *[]Buku) {
	for _, v := range *listBuku {
		fmt.Printf("%+v\n", v)
	}
}
func InterfaceAdd(obj Perpustakaan, listBuku *[]Buku) {
	obj.AddBook(listBuku)
}
func InterfaceRemove(obj Perpustakaan, indeks int, listBuku *[]Buku) {
	obj.RemoveBook(indeks, listBuku)
}
func InterfaceShow(obj Perpustakaan, listBuku *[]Buku) {
	obj.ShowBooks(listBuku)
}