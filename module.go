package module

import (
	"fmt"
	"os"
)

var member bool
var a, b, c float32

func Intro() {
	fmt.Println("\nProgram coffee shop sedang dijalankan...")
}

func Salam(a string) string {
	return "Hai " + a + ", selamat datang di coffee shop Kodehive"
}

func Recent(favorit string, jumlah float32) (float32, float32) {

	if favorit == "kopi" || favorit == "Kopi" {
		a = 30000
	} else if favorit == "air" || favorit == "Air" {
		a = 10000
	} else if favorit == "sirup" || favorit == "Sirup" {
		a = 15000
	} else {
		fmt.Println("Tidak ada menu tersebut")
		os.Exit(0)
	}

	if member {
		b = (a - (a * 10 / 100)) * jumlah
	} else {
		b = a * jumlah
	}

	return a * jumlah, b
}

func Member(nama string, status func(string) bool) {
	if status(nama) {
		fmt.Println(nama, "merupakan seorang member Kodehive")
		member = true
	} else {
		fmt.Println(nama, "bukan member Kodehive")
		member = false
	}
}

type DaftarJaga struct {
	Nama string
}

func TampilJaga() {
	jaga1 := DaftarJaga{
		Nama: "Budi",
	}
	jaga2 := DaftarJaga{
		Nama: "Agus",
	}
	jaga3 := DaftarJaga{
		Nama: "Ricky",
	}

	fmt.Println("\nDaftar nama barista yang aktif hari ini :")
	fmt.Println(jaga1.Nama, jaga2.Nama, jaga3.Nama)
}

///struct method
func (namaJaga DaftarJaga) ShiftJaga(nama string) {
	fmt.Println("\nHalo", nama, ", barista", namaJaga.Nama, "akan membantu anda")
}

//interface struct
//menghitung pajak menggunakan struct interface
type hitungInterface interface {
	menghitungPPN() float32
}

type PPN struct {
	Harga float32
}

func Menghitung(harga float32) {
	c = harga + (harga * 10 / 100)
	fmt.Println("Harga yang perlu dibayar (PPN):", c)
}

//anonymous struct
func DaftarMenu() {
	menu1 := struct {
		Nama  string
		Harga int16
	}{
		Nama:  "Kopi",
		Harga: 30000,
	}

	menu2 := struct {
		Nama  string
		Harga int16
	}{
		Nama:  "Air",
		Harga: 10000,
	}

	menu3 := struct {
		Nama  string
		Harga int16
	}{
		Nama:  "Sirup",
		Harga: 15000,
	}

	fmt.Println("\nDaftar Menu :")
	fmt.Println(menu1.Nama, "- Harga :", menu1.Harga)
	fmt.Println(menu2.Nama, "- Harga :", menu2.Harga)
	fmt.Println(menu3.Nama, "- Harga :", menu3.Harga)
}
