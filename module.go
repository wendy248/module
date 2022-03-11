package module

import "fmt"

type DaftarMenu struct {
	nama	string
	harga	int16
	stok	bool
}

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
		fmt.Println(nama, " merupakan seorang member")
	} else {
		fmt.Println(nama, " bukan member")
	}
}

func TampilMenu(){
	menu1 := DaftarMenu{
		nama: "kopi",
		harga: 30000,
		stok: true,
	}

	// menu2 := DaftarMenu{
	// 	nama : "air",
	// 	harga : 10000,
	// 	stok: true,
	// }

	fmt.Println(menu1)
	// fmt.Println(menu2)
}
