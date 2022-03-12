package module

import "fmt"

func Intro() {
	fmt.Println("\nProgram coffee shop sedang dijalankan...")
}

func Salam(a string) string {
	return "Hai " + a + ", selamat datang"
}

func Recent(favorit string) (float32, float32) {
	var a, b float32

	if favorit == "kopi" {
		a = 30000
	} else if favorit == "air" {
		a = 10000
	}

	b = a + (a * 10 / 100)
	return a, b
}

func Member(nama string, status func(string) bool) {
	if status(nama) {
		fmt.Println(nama, "merupakan seorang member")
	} else {
		fmt.Println(nama, "bukan member")
	}
}

type DaftarJaga struct {
	Nama string
	Id   int8
}

func TampilJaga() {
	jaga1 := DaftarJaga{
		Nama: "Budi",
		Id:   1,
	}
	jaga2 := DaftarJaga{
		Nama: "Agus",
		Id:   2,
	}
	jaga3 := DaftarJaga{
		Nama: "Ricky",
		Id:   3,
	}

	fmt.Println(jaga1, jaga2, jaga3)
}

///struct method
func (namaJaga DaftarJaga) ShiftJaga(nama string) {
	fmt.Println("Halo", nama, ",", namaJaga.Nama, "akan membantu anda")
}

//interface struct
//menghitung pajak menggunakan struct interface
type hitungInterface interface {
	menghitungPPN() float32
}

type PPN struct {
	Harga float32
}

func Menghitung(a float32) {
	b := a + (a * 10 / 100)
	fmt.Println("Setelah pajak :", b)
}
