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

type DaftarMenu struct {
	Nama  string
	Harga int16
	Stok  bool
}

// func TampilMenu() {
// 	Menu1 := DaftarMenu{
// 		Nama:  "kopi",
// 		Harga: 30000,
// 		Stok:  true,
// 	}

// 	fmt.Println(Menu1)
// }
type Pegawai struct {
	LamaBekerja  int8
	Nama, Divisi string
	Aktif        bool
}

func TampilData() {
	pegawai := Pegawai{
		LamaBekerja: 5,
		Nama:        "Anwars",
		Divisi:      "Keuangan",
		Aktif:       true,
	}

	fmt.Println(pegawai)
}
