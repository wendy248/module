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
}

func TampilJaga() {
	jaga := DaftarJaga{
		Nama: "Budi",
	}

	fmt.Println(jaga)
}

///struct method
func (namaJaga DaftarJaga) ShiftJaga(nama string) {
	fmt.Println("Halo", nama, ",", namaJaga.Nama, "akan membantu anda")
}

//interface struct
func (a PPN) menghitungPPN() float64 {
	return a.harga + (a.harga * 10 / 100)
}
